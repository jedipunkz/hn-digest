---
source: "https://github.com/DangerousYams/muxer"
hn_url: "https://news.ycombinator.com/item?id=48764800"
title: "Show HN: Muxer – open-source model multiplexer for Claude Code"
article_title: "GitHub - DangerousYams/muxer: Model multiplexer plugin for Claude Code: premium model orchestrates, cheaper models execute, with quality-first routing rules and a real cost report · GitHub"
author: "DangerousYams"
captured_at: "2026-07-02T17:51:52Z"
capture_tool: "hn-digest"
hn_id: 48764800
score: 2
comments: 0
posted_at: "2026-07-02T17:36:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Muxer – open-source model multiplexer for Claude Code

- HN: [48764800](https://news.ycombinator.com/item?id=48764800)
- Source: [github.com](https://github.com/DangerousYams/muxer)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T17:36:48Z

## Translation

タイトル: Show HN: Muxer – クロード コード用のオープンソース モデル マルチプレクサー
記事のタイトル: GitHub - DangerousYams/muxer: Claude Code 用モデル マルチプレクサー プラグイン: 品質優先のルーティング ルールと実際のコスト レポートを使用して、プレミアム モデルをオーケストレーションし、安価なモデルを実行 · GitHub
説明: Claude Code のモデル マルチプレクサー プラグイン: 品質優先のルーティング ルールと実際のコスト レポートを使用して、プレミアム モデルを調整し、安価なモデルを実行します - DangerousYams/muxer
HN本文：トップモデルが担当し続ける。品質を犠牲にしないルーティング ルールを使用して、入力をより安価な層に渡します。各タスクの後に保存された $ を出力します。

記事本文:
GitHub - DangerousYams/muxer: クロード コード用モデル マルチプレクサー プラグイン: 品質最優先のルーティング ルールと実際のコスト レポートを使用して、プレミアム モデルをオーケストレーションし、安価なモデルを実行します · GitHub
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
危険な山芋
/
マルチプレクサー
公共
通知
あなたはしなければなりません

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .claude-plugin .claude-plugin エージェント エージェント フック フック スクリプト スクリプト スキル/マルチプレクサ スキル/マルチプレクサ .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude Code プラグインは、安価なモデルが実際の作業のほとんどを実行しながら、高価なモデルがセッションを実行できるようにします。
Fable が自分の作業を計画し、結果を判断したいモデルであるため、これを構築しましたが、Max プランでは、上限のない追加の使用量クレジットとして請求されます。一方、コーディング セッションで行われる作業のほとんど (ファイルの grep 、ボイラープレートの作成、ログの読み取り) には、カタログで最も高価なモデルは必要ありません。 muxer は Fable をディレクターズチェアに留めておき、入力作業を Opus、Sonnet、Haiku にプッシュするか、CLI を介して OpenAI Codex や Google Gemini にプッシュします。
Claude Code は、オーケストレーターがサブタスクを生成する瞬間にサブタスクのモデルを選択します。 muxer は 3 つの方向からその決定に依存します。
Agents/*.md 内のエージェントはそれぞれ、フロントマターに model: 行を持っているため、メイン セッションが何で実行されているかに関係なく、スカウトは常に Haiku で実行され、ビルダーは常に Opus で実行されます。ここで実際の多重化が行われますが、これは提案ではなく確実な保証です。
SessionStart フック ( scripts/session-policy.sh ) は、すべてのセッションに短いルーティング ポリシーを挿入するため、オーケストレーターはすべてを自分で行うのではなく委任する必要があることがわかります。ポリシーはセッション モデルを読み取り、適応します。 Fable では、「プレミアム価格なので、メインループを無駄のないものにしてください」と書かれています。安価なモデルでは、 muxer:oracle を介して Fable までのエスカレーション パスが追加されます。
3 番目の部分は PreToolUse ガード ( scripts/guard-

モデル.sh )。 Claude Code の組み込みサブエージェント (汎用、Explore、Plan) は、オーバーライドなしで生成されたときにメイン セッションのモデルを継承します。これは、Fable セッションでは、ファイル探索がプレミアム レートで静かに請求されることを意味します。モデルが明示的に選択されていない限り、ガードはそれらのスポーンをキャッチし、opus に固定します。
注入されたポリシーは優先順位について率直です。タスクを最高の品質で実行できる最も安価なモデルを選択し、迷った場合は下位階層ではなく上位階層に進みます。バーの下に戻る安っぽい作業は、再実行には節約されたトークンよりもコストがかかるため、全体のポイントを無効にします。
そのため、重要なタスクが丸ごと引き渡されることはありません。オーケストレーターはそれらを受け入れ基準を備えた範囲指定された概要に分解し、移植や移行などの反復作業の場合、1 つのイグザンプラが強力な層で構築され、何かがファンアウトする前に検証されます。完了した作業はレビュー担当者を通過し、ある段階で 2 回レビューに失敗したタスクは、新しい概要を使用して 1 つ上の段階でやり直されます。味が重要な作業 (UI、CSS、ゲームの雰囲気) は、コストのヒントに関係なく、Opus を下回ることはありません。また、検証者が判断するビルダーよりも安価なモデルになることはありません。両方のルールが存在するのは、故障モードが現実であるためです。安価なモデルは視覚的に何かを構築しており、同様に安価なレビュー担当者はその問題点を理解できません。
代表たちも最後までやり遂げる。ビルダーは、範囲を下回っているブリーフを提案するのではなく拒否し、ライターは、実際の判断が必要であることが判明したものはすべて拒否します。レビュー担当者は各失敗を仕様の問題または機能の問題としてラベル付けするため、オーケストレーターは概要を修正するかモデルを変更するかを判断できます。
エージェント
モデル
役割
マルチプレクサー:スカウト
俳句
読み取り専用の偵察: 探索、検索、要約。
マルチプレクサー:ライター
ソネット
ドキュメント、コピー、定型文、機械的な低リスク編集。
マルチプレクサー:ビルダー
オーパス

実装とデバッグの主力。
マルチプレクサ:レビューア
オーパス
完了した作業の敵対的検証。
マルチプレクサー:オラクル
寓話
エスカレーション層: ビジュアル/クリエイティブの方向性、最も難しい決定。
マルチプレクサ:コーデックス
外部
OpenAI Codex CLI ( codex exec ) にディスパッチします。人間トークンはゼロ。
マルチプレクサ:ジェミニ
外部
Google Gemini CLI ( gemini -p ) にディスパッチします。人間トークンはゼロ。
お金の行き先
100 万トークンあたりの API リスト価格 (イン/アウト): Fable $10/$50、Opus $5/$25、Sonnet $3/$15、Haiku $1/$5。 Max プランでは、Fable は追加の使用量クレジットとして請求し、残りはプランのクォータから引き出されるため、Fable から移動されたすべてのトークンは無制限の支出をクォータの支出に変換します。 Codex または Gemini に移動されたトークンは、Anthropic 法案から完全に外されます。
サブエージェントは独自のコンテキスト ウィンドウで実行され、独自のモデルのレートで請求されます。オーケストレーターは、メイン ループを通過するもの、つまり計画、読み取ったレポート、ユーザーへの返信に対してのみプレミアム レートを支払います。だからこそ、仕事を委任するだけでは仕事の半分にすぎません。オーケストレーターも生の素材の読み取りを停止する必要があり、それがスカウトの目的であり、寓話のループに何かが触れる前に、ログの 50,000 トークンを 300 ワードのレポートに変換します。レビュアーの背後にあるのと同じ理由です。レビュアーは差分を読み取るので、メインループがプレミアムレートで差分を再読み取りすることはありません。
モード A はプラグインが調整されているものです。 Fable ( /model fable ) でセッションを実行し、オーケストレートさせます。 Fable の計画、そのビジュアル テイスト、ジョブを実際に完了するまで実行する習慣は維持されますが、プレミアム請求は無駄のないメイン ループにのみ適用されます。
モード B は最大限の節約を目的としています。 Opus ( /model opus ) でセッションを実行し、クリエイティブなディレクションやハードなデザインの呼び出しなど、本当に必要な場合にのみ muxer:oracle を通じて Fable に問い合わせます。メインループのトークンのコストは約半分で、プランの割り当て量を維持します。

事前にお見積り、事後に領収書を発行いたします。
見積もりはオーケストレーター自体から得られます。このポリシーでは、どの層が何を実行するか、定価での大まかな金額の範囲、およびそのうちの Fable クレジットがどれだけであるかをカバーする 1 行で、大規模な多重化タスクを開くように指示しています。それは推測であり、1 つとしてラベル付けされています。作業が実行される前に、フックはタスクのサイズを知ることはできません。
レシートを計測します。 Stop フックは、セッション トランスクリプトとすべてのサブエージェント トランスクリプトを解析し、クロード コードがストリーミング中に繰り返し書き込む行の重複を除去し、キャッシュを意識した API リスト レートで各ターンの価格を計算します (キャッシュ読み取りは入力レートの 0.1 倍、5 分間のキャッシュ書き込みは 1.25 倍、1 時間の書き込みは 2 倍)。次に、Fable の超過使用量クレジットと最大プラン クォータに分割された実際の支出を、セッションのメイン モデルですべてのトークンが実行された反事実と比較する 1 行を出力します。前回のレポート以降に蓄積された節約額が MUXER_REPORT_MIN_USD (デフォルトは 0.25 ドル) を超えるまでは沈黙を保つため、小さなターンがスパムになることはありません。
muxer: このストレッチのコストは ~$1.25 (~$0.00 の Fable クレジット、~$1.25 の最大クォータ) に対し、~$3.95 の非多重化されたすべての Fable です。約 2.70 ドル (68%) 節約されました。
知っておく価値のある 2 つの盲点。 muxer:codex または muxer:gemini を通じて行われた作業は外部で請求され、トランスクリプトには決して表示されないため、ここでは表示されません。また、単一モデルの実行ではサブエージェントがコンテキストを再読み取りする費用は発生しないため、多重化されていない反事実は近似値です。すべての数字は、請求書ではなく、定価の大まかな数字として読んでください。
git クローン https://github.com/DangerousYams/muxer.git
クロード プラグイン マーケットプレイスに ./muxer を追加
クロード プラグインのインストール muxer@muxer-local
後で更新するには、 claude plugin update muxer@muxer-local をプルして実行します。 Claude Code はプラグインをキャッシュされたコピーとしてインストールするため、更新するまでリポジトリを編集しても何も行われず、セッションがすでに存在します。

実行中は、再起動されるまで古いコピーが保持されます。
外部デリゲートは、CLI をインストールしてサインインする必要があります。
npm install -g @openai/codex # その後、`codex` を 1 回実行してサインインします (ChatGPT アカウント)
npm install -g @google/gemini-cli # その後、`gemini` を 1 回実行してサインインします
チューニング
MUXER_GUARD=off は、PreToolUse のデフォルト モデル ガードをオフにします。
MUXER_BUILTIN_AGENT_MODEL=<alias> は、ガードが注入するものを変更します (デフォルト opus )。
MUXER_REPORT=off は、タスク後のコスト行を沈黙させます。 MUXER_REPORT=何かを消費したターンごとに常にそれを出力します。
MUXER_REPORT_MIN_USD=<n> は、コスト ラインの節約しきい値を設定します (デフォルトは 0.25 )。
/mux は、セッション内のルーティング テーブル、コスト モデル、およびセッションの経済性を示します。
ルーティングは、オーケストレーターがどのエージェントを生成するかを選択するため、オーケストレーターにとっての助言となります。ハード保証は、エージェントごとのモデル、つまりピンと PreToolUse ガードです。
プロンプト キャッシュの再読み取りを含むオーケストレーター独自のコンテキストは、常にセッション モデルのレートで課金されます。セッションが作業の指示ではなく、主に読み取りとチャットであ​​る場合は、最初から安価なモデルで実行します。
Anthropic は、サブエージェント モデルが Max プランに対してどのように請求されるかを明示的に文書化していません。また、上記のモデルごとの料金は API の定価です。導入後の使用状況を観察し、ご自身の計画で動作を確認してください。
Claude Code のモデル マルチプレクサー プラグイン: 品質最優先のルーティング ルールと実際のコスト レポートを使用して、プレミアム モデルを調整し、安価なモデルを実行します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Model multiplexer plugin for Claude Code: premium model orchestrates, cheaper models execute, with quality-first routing rules and a real cost report - DangerousYams/muxer

Keeps the top model in charge. Hands the typing to cheaper tiers, with routing rules that don't compromise on quality. Prints $ saved after each task.

GitHub - DangerousYams/muxer: Model multiplexer plugin for Claude Code: premium model orchestrates, cheaper models execute, with quality-first routing rules and a real cost report · GitHub
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
DangerousYams
/
muxer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude-plugin .claude-plugin agents agents hooks hooks scripts scripts skills/ mux skills/ mux .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
A Claude Code plugin that lets an expensive model run the session while cheaper models do most of the actual work.
I built this because Fable is the model I want planning my work and judging the results, but on a Max plan it bills as extra usage credits with no ceiling. Meanwhile most of what happens in a coding session (grepping through files, writing boilerplate, reading logs) doesn't need the priciest model in the catalog. muxer keeps Fable in the director's chair and pushes the typing down to Opus, Sonnet, and Haiku, or out to OpenAI Codex and Google Gemini through their CLIs.
Claude Code picks a model for a subtask at the moment the orchestrator spawns it. muxer leans on that decision from three directions.
The agents in agents/*.md each carry a model: line in their frontmatter, so the scout always runs on Haiku and the builder always runs on Opus no matter what the main session runs on. This is where the actual multiplexing happens, and it's a hard guarantee rather than a suggestion.
A SessionStart hook ( scripts/session-policy.sh ) injects a short routing policy into every session so the orchestrator knows to delegate instead of doing everything itself. The policy reads the session model and adapts. On Fable it says: you're premium-priced, keep the main loop lean. On cheaper models it adds an escalation path up to Fable through muxer:oracle .
The third piece is a PreToolUse guard ( scripts/guard-model.sh ). Claude Code's built-in subagents (general-purpose, Explore, Plan) inherit the main session's model when spawned without an override, which on a Fable session means your file exploration quietly bills at premium rates. The guard catches those spawns and pins them to opus unless a model was chosen explicitly.
The injected policy is blunt about priorities: pick the cheapest model that can do the task to full quality, and when in doubt go up a tier, not down. Cheap work that comes back below the bar defeats the whole point, since redoing it costs more than the tokens saved.
So big meaty tasks never get handed off whole. The orchestrator decomposes them into scoped briefs with acceptance criteria, and for repetitive work like a port or migration, one exemplar gets built at a strong tier and verified before anything fans out. Completed work goes through the reviewer, and a task that fails review twice at one tier gets redone a tier up with a fresh brief. Taste-critical work (UI, CSS, game feel) never goes below Opus regardless of cost hints, and the verifier is never a cheaper model than the builder it's judging. Both rules exist because the failure mode is real: a cheap model builds something visually off, and an equally cheap reviewer can't see what's wrong with it.
The delegates hold up their end too. The builder refuses briefs that are under-scoped rather than winging them, and the writer bounces anything that turns out to need real judgment. The reviewer labels each failure as either a spec problem or a capability problem, so the orchestrator knows whether to fix the brief or change the model.
Agent
Model
Role
muxer:scout
Haiku
Read-only recon: explore, find, summarize.
muxer:writer
Sonnet
Docs, copy, boilerplate, mechanical low-risk edits.
muxer:builder
Opus
Implementation and debugging workhorse.
muxer:reviewer
Opus
Adversarial verification of completed work.
muxer:oracle
Fable
Escalation tier: visual/creative direction, hardest decisions.
muxer:codex
external
Dispatches to OpenAI Codex CLI ( codex exec ). Zero Anthropic tokens.
muxer:gemini
external
Dispatches to Google Gemini CLI ( gemini -p ). Zero Anthropic tokens.
Where the money goes
API list prices per 1M tokens (in/out): Fable $10/$50, Opus $5/$25, Sonnet $3/$15, Haiku $1/$5. On a Max plan, Fable bills as extra usage credits while the rest draw from plan quota, so every token moved off Fable turns open-ended spend into quota spend. Tokens moved to Codex or Gemini leave the Anthropic bill entirely.
Subagents run in their own context windows and bill at their own model's rate. The orchestrator pays its premium rate only for what passes through the main loop: its plans, the reports it reads back, its replies to you. Which is why delegating the work is only half the job. The orchestrator also has to stop reading raw material, and that's what the scout is for, turning 50k tokens of logs into a 300-word report before anything touches the Fable loop. Same reasoning behind the reviewer: it reads the diffs so the main loop doesn't re-read them at premium rates.
Mode A is what the plugin is tuned for. Run the session on Fable ( /model fable ) and let it orchestrate. You keep Fable's planning, its visual taste, and its habit of running a job until it's actually done, while premium billing applies only to the lean main loop.
Mode B is for maximum savings. Run the session on Opus ( /model opus ) and consult Fable only through muxer:oracle when something genuinely needs it, like creative direction or a hard design call. Main-loop tokens cost roughly half and stay on plan quota.
You get an estimate before and a receipt after.
The estimate comes from the orchestrator itself. The policy tells it to open any sizable multiplexed task with a single line covering which tiers will do what, a rough dollar range at list prices, and how much of that is Fable credits. It's a guess, labeled as one. No hook can know a task's size before the work runs.
The receipt is measured. A Stop hook parses the session transcript plus every subagent transcript, dedupes the rows Claude Code writes repeatedly while streaming, and prices each turn at API list rates, cache-aware (cache reads at 0.1x the input rate, 5-minute cache writes at 1.25x, 1-hour writes at 2x). It then prints one line comparing actual spend, split into Fable extra-usage credits and Max-plan quota, against the counterfactual where every token ran on the session's main model. It stays quiet until the savings accumulated since its last report cross MUXER_REPORT_MIN_USD (default $0.25), so small turns never spam you.
muxer: this stretch cost ~$1.25 (~$0.00 Fable credits, ~$1.25 Max-quota) vs ~$3.95 un-muxed all-Fable. Saved ~$2.70 (68%).
Two blind spots worth knowing about. Work done through muxer:codex or muxer:gemini bills externally and never appears in the transcripts, so it's invisible here. And the un-muxed counterfactual is approximate, because a single-model run wouldn't have paid for subagents re-reading context. Read every figure as a rough number at list prices, not an invoice.
git clone https://github.com/DangerousYams/muxer.git
claude plugin marketplace add ./muxer
claude plugin install muxer@muxer-local
To update later, pull and run claude plugin update muxer@muxer-local . Claude Code installs plugins as cached copies, so edits to the repo do nothing until you update, and sessions already running keep the old copy until restarted.
The external delegates need their CLIs installed and signed in:
npm install -g @openai/codex # then run `codex` once to sign in (ChatGPT account)
npm install -g @google/gemini-cli # then run `gemini` once to sign in
Tuning
MUXER_GUARD=off turns off the PreToolUse default-model guard.
MUXER_BUILTIN_AGENT_MODEL=<alias> changes what the guard injects (default opus ).
MUXER_REPORT=off silences the after-task cost line. MUXER_REPORT=always prints it after every turn that spent anything.
MUXER_REPORT_MIN_USD=<n> sets the savings threshold for the cost line (default 0.25 ).
/mux shows the routing table, cost model, and session economics from inside a session.
Routing is advisory for the orchestrator, since it chooses which agent to spawn. The hard guarantees are the per-agent model: pins and the PreToolUse guard.
The orchestrator's own context, including prompt-cache re-reads, always bills at the session model's rate. If a session will be mostly reading and chatting rather than directing work, run it on a cheaper model to begin with.
Anthropic doesn't explicitly document how subagent models bill against Max plans, and the per-model rates above are API list prices. Watch /usage after adopting this and confirm the behavior on your own plan.
Model multiplexer plugin for Claude Code: premium model orchestrates, cheaper models execute, with quality-first routing rules and a real cost report
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
