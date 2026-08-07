---
source: "https://github.com/masondelan/selvedge"
hn_url: "https://news.ycombinator.com/item?id=49205059"
title: "Show HN: Selvedge – an append-only log of what your AI agent already rejected"
article_title: "GitHub - masondelan/selvedge: Long-term memory for AI-coded codebases. A git blame for AI agents — but for the why. MCP server that captures the agent's reasoning live, in context, as each change is made. Local SQLite, zero deps. · GitHub"
author: "masondelan"
captured_at: "2026-08-07T02:07:21Z"
capture_tool: "hn-digest"
hn_id: 49205059
score: 1
comments: 0
posted_at: "2026-08-07T01:54:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Selvedge – an append-only log of what your AI agent already rejected

- HN: [49205059](https://news.ycombinator.com/item?id=49205059)
- Source: [github.com](https://github.com/masondelan/selvedge)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T01:54:13Z

## Translation

タイトル: Show HN: Selvedge – AI エージェントがすでに拒否した内容の追加専用ログ
記事のタイトル: GitHub - masondelan/selvedge: AI でコーディングされたコードベースの長期記憶。 AI エージェントには大きな責任がありますが、その理由は次のとおりです。 MCP サーバーは、変更が行われるたびに、エージェントの推論を状況に応じてライブでキャプチャします。ローカル SQLite、ゼロデプス。 · GitHub
説明: AI でコード化されたコードベースの長期記憶。 AI エージェントには大きな責任がありますが、その理由は次のとおりです。 MCP サーバーは、変更が行われるたびに、エージェントの推論を状況に応じてライブでキャプチャします。ローカル SQLite、ゼロデプス。 - メイソンデラン/セルビッジ

記事本文:
GitHub - masondelan/selvedge: AI でコード化されたコードベースの長期記憶。 AI エージェントには大きな責任がありますが、その理由は次のとおりです。 MCP サーバーは、変更が行われるたびに、エージェントの推論を状況に応じてライブでキャプチャします。ローカル SQLite、ゼロデプス。 · GitHub
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
メイソンデラン
/
耳
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスのメイン ブランチ タグで表示します

s ファイルコードに移動 その他のアクションメニューを開く フォルダーとファイル
181 コミット 181 コミット .claude-plugin .claude-plugin .github .github .selvedge .selvedge bin bin コマンド コマンド docs docs features/ src/ selvedge features/ src/ selvedge フック フック infra/ telemetry-worker infra/ telemetry-worker npm npm スクリプト スクリプト selvedge selvedge skill/ selvedgeスキル/耳のテスト テスト .dockerignore .dockerignore .gitignore .gitignore .mcp.json .mcp.json .mcpbignore .mcpbignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml Glama.json Glama.json manifest.json manifest.json pyproject.toml pyproject.toml server.json server.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI でコーディングされたコードベースの長期記憶 — すでに存在していたものを含む
試してみて拒否されました。
行の帰属により、誰が何かを書いたかがわかります。セルビッチはエージェントに伝えます
次に書かないこと: このコードベースがすでに試したアプローチ、
元に戻った理由と。それはAIエージェントのせいであり、むしろその理由です。
どのモデルがどのラインに触れたかよりも、エージェントによってライブでキャプチャされ、
変化が起こるので、下流側でそれを推測する必要はありません。
Selvedge はローカル MCP サーバーです。 AIコーディングエージェント（クロードコード、カーソル、
Copilot) で構造化変更イベントをログに記録するときにこれを呼び出します。
推論。データは、.selvedge/ の隣の SQLite ファイルに保存されます。
あなたのコード。
デフォルトではローカルファースト、選択によりチームサーバー、常にゼロLLM。
6 か月前、AI エージェントは user_tier_v2 という列を追加しました。あなたはそうではありません
理由がわかります。 gitblam は、生成された claude-code からのコミットを指します。
「スキーマを更新してください」というメッセージが表示されます。変更を行ったセッションが長い
消えた — そして、pr というプロンプトも消えた

それを示唆した。
Selvedge を使用する場合は、代わりにこれを実行します。
$ 耳のせい user_tier_v2
ユーザー層_v2
変更 2025-10-14 09:31:02
エージェント クロード コード
コミット3e7a991
理由 ユーザーは従来の無料利用枠に祖父フラグを追加するように要求しました
価格設定移行中のユーザー。元の階層を保存します
そのため、請求履歴に触れることなく割引を埋め戻すことができます。
その推論はその瞬間にエージェントによって捕らえられ、
変更を生じた同じコンテキストからの耳。から推測されない
その後、2 番目の LLM によって差分が取得されます。手書きのコミットメッセージではありません。
Selvedgeには2つの聴衆がいます。同じツール、同じ pip install、同じ SQLite
.selvedge/ の下にあるファイル。痛みのスケールが違う。
AI でコーディングされたコードベースを長期的に実行するチーム。
プロジェクトがあなた (または他の誰か) が触れるほど大きい場合
6か月後、12か月後、3年後にまた書きますが、そのほとんどは書かれたものです
各 PR が出荷された日にコンテキストが蒸発したエージェントによるものです。責める
何が変わったかを教えてくれます。エージェントの後でも、セルビッジがその理由を教えてくれます
セッション、プロンプト テンプレート、それを要求した開発者、およびモデル
バージョンはすべてなくなってしまいました。これは元のユースケースです: 本番環境
コードベース、スキーマの決定、移行、依存関係の変更など、
売上高を維持する監査証跡。
日常のプロジェクトで Claude Code を使用するソロ開発者。
サイド プロジェクト、週末のビルド、使い続ける小さな社内ツール。
エンタープライズ ガバナンスは必要ありません。必要なのは、なぜそうするのか (または
あなたのエージェント）は、あなたが昨日、先週、最後のスプリントでやったことを行いました。走る
耳を一度初期化します。 CLAUDE.md に 4 行を追加します。それ以来、
耳のせいは筋肉の記憶です - いつでも過去の自分と対話する方法です
あなたの前世は LLM でした。
自分の AI で構築したプロジェクトに戻って、「何だろう」と思ったことがあるなら、

だった
これまた?」、セルビッジが欠けています。
人間が書いたコードは、コミット メッセージ、PR の説明、メッセージなど、あらゆる場所に意図を漏洩します。
インライン コメント、その前の Slack スレッド。 AI が書いたコードはそうではありません。
エージェントはそれぞれの決定を下した理由を完全に明確に理解していますが、
コンテキストはプロンプト内に存在し、会話が終了すると消滅します。
6 か月後、あなたのチームは証跡のないスキーマ決定をデバッグしています。
gitblame は、いつ何が変更されたかを示します。理由はわかりません。
セルビッジは、変化が起こったときに、エージェント自身によってその理由を生き生きと捉えます。
作った。 diff は git の仕事です。その理由はセルビッチにあります。
記憶はエージェントに届き、店舗はダイヤルを受け取ります。 ２つのテーマ、
構成の半分が残りを読み取る必要があるため、一緒に出荷されます。
からの設定。
配達。 Selvedge は、元に戻されたエンティティの再編集をすでにブロックしています。何があったのか
拒否するものが何もない場合、配達がありませんでした。 2 つの新しいフック:
SessionStart は、セッションの開始時にコンパクトなダイジェストを挿入します - 決定が必要です
再訪するには、試行されて元に戻されたエンティティ、最近の変更セットを参照してください。
PreCompact は、コンテキストの圧縮によってこのセッションが破壊される直前に起動されます。
推論して、編集したがログに記録していない監視エンティティに名前を付けます。
どちらも、何も言うことがなく、サイズ制限があり、読み取り専用で、
テンプレート化された。どちらも何もブロックできません。PreCompact は意図的に
API が提供するフックを拒否します。これは測定された故障モードに対する答えです: 2
2026 年の論文には、プル モデルの記憶ツールが完全に使用されなくなったと記録されています (ゼロ
事前にシードされたストアに対する 114 ターンにわたる自発的なメモリ操作)
決定的注入は毎回実行されます。
selvedge export --format マークダウンはストアをレビュー可能としてレンダリングします
ダイジェストを使用してその隣にコミットするため、キャプチャされたインテントがプル リクエストに表示されます
こんにちはの代わりに

バイナリ内で問題が発生します。決定論的 — 新しいものを使わずに再生する
events はゼロ行の差分です。
構成。 .selvedge/config.toml はファーストクラスになり、正規の
セルビッジドクターが設定ごとに印刷する優先チェーン。それは以下をもたらします:
selvedge prune --include-events — 削除できる最初のパス
推論を捉えたものであるため、確認と
SELVEDGE_DESTRUCTIVE=1 。どちらか一方だけでは十分ではありません。なぜなら --yes は cron 内にあるからです
エントリはプロンプトを無効にし、シェル プロファイルは環境変数を無効にします。イベント
保持のデフォルトは「なし」です。
大幅に切り詰められるイベント サイズの境界 ( diff_bytes 、reasoning_bytes )
— テキスト内のマーカー、書き込み時の警告、耳の統計のカウント。
log_change でのシークレットシェイプの警告、経由で拡張可能
redaction_patterns に加えて、すでに保存されているものをスキャンするドクター行。
警告し、決して拒否しないでください。
また、5 件のレビュー問題が終了しました。強制フックの許可パスは次のとおりです。
40% 高速化 (ゲート呼び出しあたり 33.6 ミリ秒 → 20.1 ミリ秒) および SELVEDGE_HOOK_DISABLE=1
最後に、スキップすることが文書化されているインポートの前に短絡します。
log_change は revisit_after /constraint/stale_when を破棄しなくなりました
名前の変更と置き換えについて。 CLI の --json と MCP ツールが返されるようになりました。
同一の構造。また、Docker イメージにはメンテナーのイメージが同梱されなくなりました。
独自のデータベース。テスト826→984。
壊れたインストールを修正し、完全なコード品質のパスを取得します。 mcp2.0.0
(2026-07-28 リリース) mcp.server.fastmcp が削除され、Selvedge が宣言されました
mcp>=1.0.0 (上限なし) - その日付以降の pip インストールセルビッジ
2.0.0 をプルし、selvedge-server がインポート時に失敗しました。このリリースでは、
依存性。サーバーが起動しなくなった場合は、アップグレードが必要です。
コードベース上に 9 つの独立したパスを配置するレビューとともに出荷されます。
そして、それに基づいて行動する前に、すべての発見を反証しようとしました。セブン

ティーン
不具合が修正されていることを確認しました。実際に気づいたであろうもの:
強制フックはブロックすべきではないものをブロックしなくなりました。読書
追跡されたファイル — cat 、 git diff 、 pytest 、 ruff check — がブロックされ、
エラー メッセージで実行するよう指示された修復は、同じエラー メッセージによってブロックされました。
そのため、CLI から出る方法はありませんでした。同じものを供給するさらに 2 つのパス
false ブロック: コメントアウトされた SQL 行は実際の削除としてカウントされます。
「revert」という単語を含むだけのコミットメッセージは、すべてのファイルにマークを付けます
元に戻されたようにタッチされました。
ルックアップが大規模に高速化されました。主要なエンティティの読み取りはすべての行をスキャンしていました -
100k イベントで 7.4 ミリ秒→ 0.35 ミリ秒と測定され、フックに時間がかかっていました
大型店舗では数秒。
耳のセットアップでは CLAUDE.md の一部を削除できなくなりました。
中断されたバックアップによって最後の正常なバックアップが破壊されることはなくなり、アップグレードが可能になります
2 つの Selvedge プロセスが実行中にエラーが発生してクラッシュすることはなくなりました。
データベースが破損しているようです。
テストは 739 → 826 になりました。スキーマの変更やツール表面の変更はなかったので、これは次のとおりです。
0.3.9.x 上の誰でもドロップインできます。
AI エージェントは作業中に Selvedge に電話をかけます。セルビッジがその理由を捉えています
耐久性がありクエリ可能なストアに格納し、それを送り返します。
エージェントトレースレコード
クロスツールリーダーにリンクする可観測性メタデータとして
Sentry/Datadog スタック トレース、および SOC 2 のコンプライアンス アーティファクトとして
EU AI 法の監査。
Selvedge は git に代わるものではありません (行レベルの内容/いつ)、PR レビュー
ツール (レビュー時の品質)、エージェントの可観測性 (LLM コール トレース)、
または汎用コードホスト AI 機能。それはそれらの間に位置します -
第一級市民としての来歴層、他のすべてのもの
参考文献。
「AI エージェントに対する git の責任」というカテゴリが急速に成長しています。ここが場所です
耳の部分はフィットしますが、意図的にフィットしない部分もあります。
「拒否されたパス」が重要な理由

— コピーできないもの。高価なもの
失敗とは、列が存在する理由を忘れることではありません。自信を持ってエージェントです
正当な理由でチームが既に廃止したものを再実装する、6
それを知っていた全員がコンテキスト ウィンドウを離れてから数か月後。どれも
表面上のライン属性ツールはパスを拒否しましたが、それは問題ではありません。
機能のギャップはリリースで埋めることができます - ライン指向のストアには概念がありません
試行→元に戻す→再試行のサイクルにわたって持続したエンティティの。参照
docs/demos/prior-attempts.md 。
決定論が重要な理由セルビッチの推論はエージェント自身の意図であり、
変更を行った同じコンテキスト ウィンドウから書き込まれます。ありません
モデルはストレージまたは取得パスのどこにでも存在するため、同じクエリで
現在も 2 年後も、モデルのバージョンを超えて同じ答えが得られます。推論するツール
事後推論では、元の LLM を一度も見たことのない 2 番目の LLM が実行されています。
プロンプト: 生成されるのは言い換えであり、再実行すると生成される可能性があります
同じ変更に対して異なるカテゴリ。 Hacker Newsのコメント投稿者はこう述べています
競合するアプローチについて、「拒否されたため、grep はコミットを見つけられません」
「oauth-library」…決定的な強制がない限り」
( 0x457 )。
決定論だけではもはや区切り文字ではありません — OpenLore は決定論ネイティブです
も、そして

[切り捨てられた]

## Original Extract

Long-term memory for AI-coded codebases. A git blame for AI agents — but for the why. MCP server that captures the agent's reasoning live, in context, as each change is made. Local SQLite, zero deps. - masondelan/selvedge

GitHub - masondelan/selvedge: Long-term memory for AI-coded codebases. A git blame for AI agents — but for the why. MCP server that captures the agent's reasoning live, in context, as each change is made. Local SQLite, zero deps. · GitHub
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
masondelan
/
selvedge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
181 Commits 181 Commits .claude-plugin .claude-plugin .github .github .selvedge .selvedge bin bin commands commands docs docs features/ src/ selvedge features/ src/ selvedge hooks hooks infra/ telemetry-worker infra/ telemetry-worker npm npm scripts scripts selvedge selvedge skills/ selvedge skills/ selvedge tests tests .dockerignore .dockerignore .gitignore .gitignore .mcp.json .mcp.json .mcpbignore .mcpbignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml glama.json glama.json manifest.json manifest.json pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Long-term memory for AI-coded codebases — including what was already
tried and rejected.
Line attribution tells you who wrote something. Selvedge tells your agent
what not to write next: the approaches this codebase already tried,
reverted, and why. It's a git blame for AI agents, for the why rather
than which model touched which line — captured live, by the agent, as the
change happens, so nothing downstream has to guess at it.
Selvedge is a local MCP server. AI coding agents (Claude Code, Cursor,
Copilot) call it as they work to log structured change events with
reasoning. Your data stays in a SQLite file under .selvedge/ next to
your code.
Local-first by default, team-server by choice, zero-LLM always.
Six months ago, your AI agent added a column called user_tier_v2 . You don't
know why. git blame points to a commit from claude-code with a generated
message that says "Update schema." The session that made the change is long
gone — and so is the prompt that produced it.
With Selvedge, you run this instead:
$ selvedge blame user_tier_v2
user_tier_v2
Changed 2025-10-14 09:31:02
Agent claude-code
Commit 3e7a991
Reasoning User asked to add a grandfathering flag for legacy free-tier
users during the pricing migration. Stores the original tier
so we can backfill discounts without touching billing history.
That reasoning was captured by the agent in the moment — written into
Selvedge from the same context that produced the change. Not inferred from
the diff afterward by a second LLM. Not a hand-typed commit message.
Selvedge has two audiences. Same tool, same pip install , same SQLite
file under .selvedge/ . Different scale of pain.
Teams running long-term, AI-coded codebases.
When the project is big enough that you (or someone else) will touch it
again in six months, twelve months, three years — but most of it was written
by an agent whose context evaporated the day each PR shipped. git blame
tells you what changed. Selvedge tells you why — even after the agent
session, the prompt template, the developer who asked for it, and the model
version are all long gone. This is the original use case: production
codebases, schema decisions, migrations, dependency changes that need an
audit trail that survives turnover.
Solo developers using Claude Code on everyday projects.
Side projects, weekend builds, the small internal tool you keep poking at.
You don't need enterprise governance — you just need to remember why you (or
your agent) did the thing you did yesterday, last week, last sprint. Run
selvedge init once. Add four lines to your CLAUDE.md . From then on,
selvedge blame is muscle memory — a way to talk to your past self when
your past self was an LLM.
If you've ever come back to your own AI-built project and thought "what was
this for again?", Selvedge is the missing piece.
Human-written code leaks intent everywhere — commit messages, PR descriptions,
inline comments, the Slack thread that preceded it. AI-written code doesn't.
The agent has perfect clarity about why it made each decision, but that
context lives in the prompt and evaporates when the conversation ends.
Six months later, your team is debugging a schema decision with no trail.
git blame tells you what changed and when . It can't tell you why .
Selvedge captures the why — live, by the agent itself, as the change is
made. The diff is git's job. The why is Selvedge's.
The memory comes to the agent, and the store gets its dials. Two themes,
shipped together because the config half is what the rest needed to read
settings from.
Delivery. Selvedge already blocked re-edits of reverted entities. What was
missing was delivery when there is nothing to veto. Two new hooks:
SessionStart injects a compact digest as a session begins — decisions due
for a revisit, entities that were tried and reverted, recent changesets.
PreCompact fires just before context compaction destroys this session's
reasoning and names the watched entities you edited but never logged.
Both are quiet when they have nothing to say, size-capped, read-only, and
templated. Neither can block anything — PreCompact deliberately declines the
veto the hook API offers it. This is the answer to a measured failure mode: two
2026 papers recorded pull-model memory tools going unused entirely (zero
voluntary memory operations across 114 turns against a pre-seeded store) while
deterministic injection landed every time.
selvedge export --format markdown renders the store as a reviewable
digest to commit next to it, so captured intent shows up in a pull request
instead of hiding inside a binary. Deterministic — regenerating with no new
events is a zero-line diff.
Config. .selvedge/config.toml is now first-class, with a canonical
precedence chain that selvedge doctor prints per setting. It brings:
selvedge prune --include-events — the first path that can delete
captured reasoning, so it needs both a confirmation and
SELVEDGE_DESTRUCTIVE=1 . Neither alone is enough, because --yes in a cron
entry defeats a prompt and a shell profile defeats an env var. Events
retention defaults to never.
Event-size bounds ( diff_bytes , reasoning_bytes ) that truncate loudly
— a marker in the text, a warning at write time, a count in selvedge stats .
Secret-shape warnings at log_change , extendable via
redaction_patterns , plus a doctor row that scans what's already stored.
Warn, never reject.
Also: five review issues closed. The enforcement hook's allow path is
40% faster (33.6 ms → 20.1 ms per gated call) and SELVEDGE_HOOK_DISABLE=1
finally short-circuits before the imports it was documented to skip;
log_change no longer discards revisit_after / constraint / stale_when
on renames and supersedes; the CLI's --json and the MCP tools now return
identical structures; and the Docker image no longer ships the maintainer's
own database. Tests 826 → 984.
Fixes a broken install, and lands a full code-quality pass. mcp 2.0.0
(released 2026-07-28) removed mcp.server.fastmcp , and Selvedge declared
mcp>=1.0.0 with no upper bound — so any pip install selvedge after that date
pulled 2.0.0 and selvedge-server failed at import. This release pins the
dependency. If your server stopped starting, this is why — upgrade.
It ships alongside a review that put nine independent passes over the codebase
and then tried to disprove every finding before acting on it. Seventeen
confirmed defects fixed. The ones you would actually have noticed:
The enforcement hook stopped blocking things it shouldn't. Reading a
tracked file — cat , git diff , pytest , ruff check — was blocked, and
the remediation the error message told you to run was blocked by the same
gate, so there was no way out from the CLI. Two more paths fed the same
false blocks: a commented-out line of SQL counted as a real deletion, and any
commit message merely containing the word "revert" marked every file it
touched as reverted.
Lookups got fast at scale. The main entity read was scanning every row —
measured 7.4 ms → 0.35 ms at 100k events, and the hook had been taking
seconds on large stores.
selvedge setup can no longer delete parts of your CLAUDE.md , an
interrupted backup can no longer destroy your last good one, and upgrading
while two Selvedge processes are running no longer crashes with an error that
looked like database corruption.
Tests went 739 → 826. No schema change and no tool-surface change, so this is
drop-in for anyone on 0.3.9.x .
AI agents call Selvedge as they work. Selvedge captures the why
into a durable, queryable store and emits it back out — as
Agent Trace records for
cross-tool readers, as observability metadata that links into
Sentry/Datadog stack traces, and as compliance artifacts for SOC 2
and EU AI Act audits.
Selvedge does not replace git (line-level what/when), PR review
tools (review-time quality), agent observability (LLM call traces),
or general-purpose code-host AI features. It sits between them — the
provenance-as-first-class-citizen layer that everything else
references.
There's a fast-growing "git blame for AI agents" category. Here's where
Selvedge fits — and where it deliberately doesn't.
Why "rejected paths" matter — the one that isn't copyable. The expensive
failure isn't forgetting why a column exists. It's an agent confidently
re-implementing something the team already killed for a good reason, six
months after everyone who knew that left the context window. None of the
line-attribution tools above surface rejected paths at all, and it isn't a
feature gap they can close in a release — a line-oriented store has no notion
of an entity that persisted across a try → revert → retry cycle. See
docs/demos/prior-attempts.md .
Why determinism matters. Selvedge's reasoning is the agent's own intent,
written from the same context window that produced the change. There is no
model anywhere in the storage or retrieval path, so the same query returns the
same answer today and in two years, across model versions. Tools that infer
reasoning post-hoc are running a second LLM that never saw the original
prompt: what it produces is paraphrase, and re-running it can produce
different categories for the same change. As a Hacker News commenter put it
about a competing approach, "grep won't find your commit because you rejected
'oauth-library'… unless there is deterministic enforcement"
( 0x457 ).
Determinism alone is no longer a separator — OpenLore is deterministic-native
too, and

[truncated]
