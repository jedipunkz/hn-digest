---
source: "https://github.com/wdhwg001/csift"
hn_url: "https://news.ycombinator.com/item?id=49356290"
title: "Show HN: Csift – the missing tool to sift your Claude Code sessions"
article_title: "GitHub - wdhwg001/csift: The missing tool to sift your Claude Code sessions: regex search across all record types, recover files, extract images, inspect subagent topologies, match between plan files and sessions. · GitHub"
image: "https://opengraph.githubassets.com/29d4a4b01cf000b45202ab1da4267312d08285dcb338ef733832c56fe3f11727/wdhwg001/csift"
author: "wdhwg001"
captured_at: "2026-08-19T03:38:34Z"
capture_tool: "hn-digest"
hn_id: 49356290
score: 1
comments: 0
posted_at: "2026-08-19T03:24:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Csift – the missing tool to sift your Claude Code sessions

- HN: [49356290](https://news.ycombinator.com/item?id=49356290)
- Source: [github.com](https://github.com/wdhwg001/csift)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T03:24:31Z

## Translation

タイトル: HN を表示: Csift – クロード コード セッションを選別するための不足しているツール
記事のタイトル: GitHub - wdhwg001/csift: クロード コード セッションを選別するための欠けているツール: すべてのレコード タイプにわたる正規表現検索、ファイルの回復、イメージの抽出、サブエージェント トポロジの検査、プラン ファイルとセッション間の一致。 · GitHub
説明: クロード コード セッションを選別するための不足しているツール: すべてのレコード タイプにわたる正規表現検索、ファイルの回復、イメージの抽出、サブエージェント トポロジの検査、プラン ファイルとセッション間の照合。 - wdhwg001/csift
HN テキスト: クロードによる独自のセッション ファイルの処理に関するすべての不満に対処するために csift を構築しました。クロード コードのセッションはすべてを含むプレーンテキストの JSONL ですが、構造がどこにも文書化されることはありません。毎回、クロードは推測するためのアドホック ツールを作成する必要がありますが、SKILL.md だけではすべてをカバーするのに十分ではありません。これがその夜の出来事で、これが本当のツールになったのです。クロードはプロジェクト ファイルを削除し、フリーズし、真夜中にごめんなさいと言いました。私は Ralph をループさせず、独自の JSONL から回復して続行するように依頼しました。簡単だと思っていましたが、書き換えの途中で再び停止してしまいました。これは、サブエージェントに分散されており、5 時間の制限が切れたためです。 --dangerously-skip-permissions を使用したのはこれが初めてではありませんでした。そして、ファイルをバイト単位で元に戻すことができるツールを作成しました。 csift は、削除されたファイルと計画を回復し、ターンを再構築し、計画とセッションを照合し、25 のレコード ラベルに対する型指定検索を実行し、貼り付けられたイメージを抽出し、サブエージェント トポロジ ツリーをマッピングします。これは純粋にローカルであり、セッションに対して読み取り専用であり、テレメトリはありません。そして、埋め込みをせず、デーモンを必要としない純粋な正規表現を使用することにしました。これは mmap、SIMD スキャン、レーヨンを使用して Rust で実行されるため、高速です。手書きで時間を無駄にしたくない

Claude Code がこのプロジェクトを作成し、私がそれをレビューしました。突然変異テストに合格し、94.9% のカバー率を達成したので、喜んで共有します。 README.md を除いて人間が判読できるドキュメントを省略することを意図していたので、人間のエンジニアが時間をかけて手動で改良する必要はありません。クロードに仕事を任せてください。カーゴインストール cshift
npx スキル追加 wdhwg001/csift
楽しんで、ご意見をお聞かせください。

記事本文:
GitHub - wdhwg001/csift: クロード コード セッションを選別するための不足しているツール: すべてのレコード タイプにわたる正規表現検索、ファイルの回復、イメージの抽出、サブエージェント トポロジの検査、プラン ファイルとセッション間の一致。 · GitHub
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
wdhwg001
/
シフト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
197 コミット 197 コミット フォルダーとファイル
.cargo-husky/ フック .cargo-husky/ フック アセット アセット s

rc src テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md SKILL.md SKILL.md SPEC.md SPEC.md Clippy.toml Clippy.toml Rustfmt.toml Rustfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
「c-シフト」（シーシフト）と発音します。
Claude Code セッションのトランスクリプトに欠けているツール。
.jsonl ログから直接クロード コード セッションを検索、回復、監査します。
───────────────────────
❯ そのセッションでレート制限について何を決定するのか、またコードはどこにあるのか?
───────────────────────
⏺ Bash(csift 検索 "レート制限" @13d9645a -t エージェント --since 1d)
1 つの交換、1 つのセッション、古い順に一致します
13d9645a·t42 2026-06-20 22:14:07.811 AEST(UTC+10)
▸ Agent.message L8821 スライディング ウィンドウ リミッター (10 req/min/IP) を追加しました。今は429号線
Retry-After を返し、問題のある IP (gateway/src/rate_limit.rs:88) をログに記録します。
一致した 1 つの交換 · 1 セッション · ラベル=エージェント
1 つの正規表現、完全なラウンドトリップ、トークン効率の高い出力。埋め込みもデータベースもデーモンもありません。
Claude Code のセッションは JSONL のプレーンテキストですが、なぜ grep ではなく csift を使うのでしょうか?
はい。これには、すべてのプロンプト、思考、ツール呼び出し、ファイル編集、貼り付けられたイメージ、および生成されたサブエージェントが含まれます。しかし、次のいずれかがあなたに起こるかもしれません。
「私のプロジェクトファイルを削除し、その後謝罪しました。」
--dangerously-skip-permissions を実行すると、

構築したファイルを消去し、ごめんなさいと言って、立ち止まりました。自身のセッションからそれらを掘り戻すように要求されたとき、それは手探りで (編集はサブエージェントにまたがり、すべての復元は部分的に戻ってきました)、その後、パニックに陥り、そのほうが「簡単である」と判断したため、静かにメモリからそれらを書き直し始めました。
「圧縮後にデザイン画像が失われ、再送信するように求められました。」
イメージは会話の中にありました。 1 回の圧縮の後、エージェントはそれらが存在しなかったかのように動作します。つまり、圧縮では代わりに損失のあるテキストの断片が渡され、それを信じました。
「jazzy-twilight-sparkle.md はどのセッションに属しますか?」
計画ファイルには重要なアーキテクチャが含まれます。その発明された名前には、どのセッションがそれを作成したかについては何も示されておらず、バインディングは EnterPlanMode ツール呼び出しの中に存在しており、干し草の山の中の針です。
「3 つのセッションが開いているときにマシンがクラッシュしました。どれがどれでしたか?」
フロントエンド、バックエンド、デバッグ、およびクロード --resume を使用すると、顔のないラインナップが得られます。
長期実行に数回の圧縮を行っても、エージェントはまだ自分のタスクを認識しています。あなたが実際に言ったことや読んだことはすべて失われ、コーデックスやパイのように最後の数ターンを機械的に維持するものはありません。
「私はあなたに言いました。質問ダイアログで。」
ユーザーの役割ごとに独自のセッションを検索しましたが、何も見つかりませんでした。そして、あなたがそれを言ったことはないと結論付けました。 AskUserQuestion でそれを言いました。
「あなたはオーケストレーターです。なぜどのセッションが停止しているのかがわからないのですか?」
なぜなら、未回答の AskUserQuestion はディスクにまったく書き込まれず、「したがって、何も見ることができない」からであると説明されています。さらに悪いことに、「はい、考え方から解釈できます。」
それぞれの結末は同じです。エージェントは「あなたの言うとおりです」と言いました。 、その後、生の JSON を使用した 1 回限りのスクリプトでつまずいて、微妙に間違っています。
csift にはあるべきツールがありません。
⏪ ファイルと削除されたプランを復元します。
csift回復

すべての編集/読み取りを分析および集計し、いつでも、差分パッチまたは最終ファイル内のファイルの正確なバイトを復元し、クロード コードのファイル変更マーカーを尊重します。境界が変更されたために確実に回復できない場合は、ギャップがマークされて生き残ったものをサルベージします。
csift イメージのリスト、重複除去、抽出を行い、セッションで使用するのと同じ [イメージ #N] ハンドルでアドレス指定できます。 ( #N は一意の ID ではありませんでした。csift は対応します。)
csift plan は、セッションが作成した計画を検索します。
csift plan --reverse jazzy-twilight-sparkle.md は、ファイルを所有するセッションに名前を付けます。
両方向。
📇 見分けられるセッション。
csift list は、最初と最後のメッセージによってすべてのセッションを識別します。
これは、claude --resume では決して得られなかった完成です。
概要はタスクの状態とコンテキストを保持します。会話は別の軸であり、 csift は --budget 内で切り取られたやり取りを逐語的に再構築します。それをフックに接続すると、すべての圧縮に最近のダイアログが添付されて届きます。
バックグラウンドタスクの完了通知は「ユーザー」ロールです。
サブエージェントの戻り値: これも "user" です。
AskUserQuestion の答え: ツールの結果。
csift はすでにこれらの罠のすべてに足を踏み入れているため、 csift search -t user.answer (25 個の {role}.{class}.{sub} ラベルの 1 つ) を実行すると、単純な grep が決して言わなかったと断言する内容が正確に見つかります。
🛎 保留中の質問は記録に残ります。
csift は、AskUserQuestion / ExitPlanMode / MCP の引き出しをサイドカー ファイルに記録するフックを同梱しており、すべての csift サーフェスは未解決のものを透過的にマージします。オーケストレーターは最終的に、どのセッションが人間の待機中にスタックしているのか、また何を担当しているのかを確認できるようになります。
ヒットすると、uuid /parentUuid グラフから再構築された交換全体が返されます。つまり、一致したツール呼び出しとその結果、ユーザーのターンとエージェントの応答が返されます。これは、grep や Claude のアドホック スクリプトが決して依存できないコンテキストです。

素直に提供してください。
生成されたすべてのエージェントの種類、ライフサイクル、および親→子ツリーに加えて、保留中の許可承認で凍結されたレーンの検出。
🤖 人間とLLM向けに設計されています。
出力は簡潔で、再フィード可能で、さらに構造的です。スキルをインストールするだけで、クロードがパワーを獲得します。
純粋な正規表現。埋め込み、インデックス、データベース、デーモン、ネットワーク、テレメトリ、隠れた検出はありません。すでにディスク上にあるファイルを読み取り、セッション履歴を変更することはありません。
⚡ Rust + mmap + SIMD 改行スキャン + バイト プレフィルター + rayon。
200 MB のトランスクリプトと数 GB のコーパスを約 1 秒で処理でき、フック内から気づかずに呼び出すことができるほどの速さです。
カーゴインストール cshift
Rust 1.89以降が必要です。またはソースから:
git クローン https://github.com/wdhwg001/csift.git
cd cシフト
カーゴインストール --path 。 # 最適化されたバイナリをビルドし、`csift` を PATH に置きます
# …または単に `cargo build --release` → ./target/release/csift
csift はデフォルトで ~/.claude を読み取ります。 --claude-home <DIR> または Claude Code 独自の $CLAUDE_CONFIG_DIR を使用して他の場所に指定します。 csift <command> --help は完全なマニュアルです。
csift の主要ユーザーはエージェント自体です。出力は簡潔で解析可能で、すべてのレコードには再フィード可能なハンドルが付いています。このスキルは、クロード コードにいつ、どのように到達するかを教えます。
npx スキル追加 wdhwg001/csift
クイックスタート
csift <コマンド> [ターゲット] [フラグ] 。ターゲットは、位置 @<uuid> (セッション)、@<agent-hex> (サブエージェント)、プロジェクト パス、または です。 (このCWD);すべてのプロジェクトをスキャンするには省略します。
完全なフラグ セットと例については、 csift <command> --help を実行してください。
リスト
早い「これはどのセッションですか？」インデックス: 最初/最後のユーザー + 最後のエージェント、セッションごと
検索
トランスクリプト上の正規表現 → ヒットごとの完全なラウンドトリップ (-t / -T ラベル フィルター、-l マッチング セッション、--raw 逐語的行)
ショー
正確なレコードを取得します

名前: --line N|A..B / --uuid U / --turn N|A..B|-k (検索ヘッダーからの ·tN ターン インデックス)、完全なバイトまたは --raw バイトでレンダリングされます。
統計
セッションごとの 1 スキャン集計: モデル別のトークン、ツール呼び出し、ターン、スパン、圧縮
エージェント
セッションのサブエージェント: 種類、ライフサイクル、ステータス、および親→子のトポロジ
おいおい
$CLAUDE_CODE_SESSION_ID から呼び出しセッションを識別し、誤検知に対して安全です
ファイル
セッションがいつ変更したファイル/ディレクトリ、およびツール ストリームの外部で行われた編集
回復する
読み取り/書き込み/編集ストリームからファイル (または削除されたプラン) を再構築します: バイト正確または正直なギャップ
計画
セッションにバインドされた Plan-Mode プラン ファイルを見つけます (逆に、どのセッションがプランを所有しているか)
逐語的に
予算内で、圧縮の概要を切り取って逐語的に復元します（ライブテールピークは show --turn です）。
画像
リスト + トランスクリプトに貼り付けられた画像の抽出 (ハンドル/ロケーターのアドレス指定、フォーマットのトランスコード)
要約は選択です。 csift は会話を続けます。
クロード コードを圧縮すると、高密度の要約が再生成されます。数十の履歴圧縮の代わりにキロバイトが必要になります。良い選択ですね。主要な調査結果 (file:line で何が変わったのか) と、多くの場合、常駐の指示をそのまま保持します。しかし、それは選択であり、毎回再抽象化され、最適化される軸は会話ではなくタスクの継続です。
csift 逐語的は、もう 1 つの軸を保持します。つまり、逐語的な User ↔ Agent のやり取り、実際に言ったこと、および終了時にエージェントが実際に報告した内容を保持します。その間にある何百ものツール呼び出しはカウントにまとめられます。最近のウィンドウの場合、サマリーが 1 ～ 2 行にとどまる完全忠実度のダイアログが数十 KB に達します。あなたの指示だけでなく、エージェントの「HEAD に対して検証済み、前提条件が成り立つ、これが証拠です」というレポートバックがタスクフィンに送信されます。

ディンスの選択は単にキャリーしません。
概要に代わるものではありません。圧縮後のコンテキストを、予算制限付きで最新の会話を完全に再現して拡張します。最も古いターンには到達せず、圧縮されていないプラン ファイルが引き続きプランを所有します。これを SessionStart(compact) フック ( SKILL.md を参照) に接続すると、すべての圧縮に最近の逐語的な会話が添付されて届きます。
csift は、必要がない場合はトランスクリプト全体をロードしません。各 .jsonl を mmap し、SIMD ( memchr ) で改行をスキャンし、安価な byte/regex プレフィルターを実行し、候補行のみ serde_json を解析します。 list は先頭と末尾のみを読み取ります。レーヨンはファイル全体に広がります。難しいのはスピードではありません。これは、Claude Code のログのセマンティクスです。role:"user" レコードは通常、ツールの結果であり、人間の手によるものではありません。保留中の質問がディスクに書き込まれることはありません。圧縮には特定の形状があります。サブエージェントはディスク上にフラットに配置され、そのトポロジは生成ツールの呼び出しから再構築されます。すべては経験に基づいて SPEC.md に文書化されています。
人間の感覚では、何もありません。このリポジトリは完全に Claude Code によって作成および保守されており、3 つの参照ファイルは保守エージェントが作業するコーパスです。
SKILL.md : csift を操作するためにエージェントがロードする使用法リファレンス。
SPEC.md : 設計意図、レコードモデルのセマンティクス、および深い機能の背後にある測定。
AGENTS.md : リポジトリの操作マニュアル。アーキテクチャ、jsonl ドメインの知識、品質のゲート。
これらは高密度で相互参照されており、ユーザーの注意を引くためではなく、モデルの注意を引くために書かれています。

[切り捨てられた]

## Original Extract

The missing tool to sift your Claude Code sessions: regex search across all record types, recover files, extract images, inspect subagent topologies, match between plan files and sessions. - wdhwg001/csift

I built csift to address all my complaints about Claude's handling of its own session files. Claude Code sessions are plain-text JSONL with everything, but they never document the structure anywhere. Every time, Claude has to craft an ad hoc tool to guess, and SKILL.md isn't enough to cover everything. This is what happened in the night that made it a real tool: Claude deleted project files, froze, and said sorry at midnight. I didn't Ralph loop it, and I asked it to recover from its own JSONL and continue. I thought it would be easy, until it stopped again in the middle of rewriting them, since they're scattered across subagents, with its 5hr limit burnt out. This wasn't the first time after I used --dangerously-skip-permissions. And I created the tool that could put the file back together byte-exact. Now csift recovers deleted files and plans, reconstructs turns, matches plans to sessions, does typed search over 25 record labels, extracts pasted images, and maps the subagent topology tree. It's purely local, read-only to sessions, and has no telemetry. And I decided to use pure regex, no embeddings, and it requires no daemon. It's done in Rust with mmap, a SIMD scan, and rayon, so it's fast. I'd rather not waste my time hand-writing a tool to pick up Claude's slack, so Claude Code wrote this project, and I reviewed it. It passed mutation tests and achieved 94.9% coverage, so I'm happy to share it. I intended it to lack human-readable documentation except for the README.md, so human engineers shouldn't spend time polishing it by hand. Let your Claude do the work. cargo install csift
npx skills add wdhwg001/csift
Enjoy and let me know what you think.

GitHub - wdhwg001/csift: The missing tool to sift your Claude Code sessions: regex search across all record types, recover files, extract images, inspect subagent topologies, match between plan files and sessions. · GitHub
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
wdhwg001
/
csift
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
197 Commits 197 Commits Folders and files
.cargo-husky/ hooks .cargo-husky/ hooks assets assets src src tests tests .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SKILL.md SKILL.md SPEC.md SPEC.md clippy.toml clippy.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
Pronounced "c-sift" ( see-sift )
The missing tool for Claude Code session transcripts.
search , recover , and audit any Claude Code session straight from the .jsonl logs.
──────────────────────────────────────────────────────────────────────────────
❯ what that session decide about rate limiting and wheres the code?
──────────────────────────────────────────────────────────────────────────────
⏺ Bash(csift search "rate limit" @13d9645a -t agent --since 1d)
matches 1 exchange · 1 session · oldest first
13d9645a·t42 2026-06-20 22:14:07.811 AEST(UTC+10)
▸ agent.message L8821 Added a sliding-window limiter (10 req/min/IP); the 429 path now
returns Retry-After and logs the offending IP — gateway/src/rate_limit.rs:88.
matched 1 exchange · 1 session · label=agent
One regex, the complete round-trip , token-efficient output. No embeddings, no database, no daemon.
Claude Code's sessions are plain text in JSONL, so why csift, not grep?
Yes. It has every prompt, thought, tool call, file edit, pasted image, and subagent it spawned. But one of these may happen to you:
"It deleted my project files, then apologized."
Running --dangerously-skip-permissions , it wiped the files it had built, said sorry, and stopped dead. Asked to dig them back out of its own session, it fumbled around (the edits spanned subagents, every restore came back partial) and then quietly started rewriting them from memory as it panicked and decided that was "easier".
"It lost my design images after a compaction, then asked me to re-send them."
The images were right there in the conversation. One compaction later the agent acts like they never existed: the compaction handed it lossy text fragments instead, and it believed them.
"Which session does jazzy-twilight-sparkle.md belong to?"
The plan file carries the architecture that matters. Its invented name says nothing about which session wrote it, and the binding lives in an EnterPlanMode tool call, a needle in the haystack.
"The machine crashed with three sessions open. Which was which?"
Frontend, backend, debugging, and claude --resume gives you a lineup with no faces.
A few compactions into a long run, the agent still knows its task . Everything you actually said and read is gone, and nothing keeps the last few turns mechanically the way codex or pi do.
"I DID tell you. In the question dialog."
It searched its own session by user roles, found nothing, and concluded you never said it. You said it in AskUserQuestion .
"You're the orchestrator. Why can't you see which session is stuck?"
Because, it explains, an unanswered AskUserQuestion is never written to disk at all, "so there's nothing I can watch." Or worse: "Yes, I can interpret it from the thinking."
Each of these ends the same way: the agent says You are absolutely right. , then stumbles through a one-off script over raw JSON and gets it subtly wrong.
csift is the missing tool it should have had.
⏪ Recover files & deleted plans.
csift recover analyzes and aggregates every edit/read and restores the file's exact bytes, at any point in time, in diff patches or final files, and honors Claude Code's file-modified markers. When it can't reliably recover due to modified boundaries, it salvages what survived with gaps marked.
csift image lists, dedups, and extracts them, addressable by the same [Image #N] handle the session uses. ( #N was never a unique id. csift copes.)
csift plan finds the plan a session wrote.
csift plan --reverse jazzy-twilight-sparkle.md names the session that owns a file.
Both directions.
📇 Sessions you can tell apart.
csift list identifies every session by its first and last messages.
It is the completion claude --resume never had.
A summary keeps task state and context. The conversation is a different axis, and csift verbatim reconstructs the clipped back-and-forth within a --budget . Wire it into a hook and every compaction arrives with the recent dialogue attached.
A background task's completion notice is a "user" role.
A subagent's return: also "user" .
Your AskUserQuestion answer: a tool result .
csift stepped in every one of these traps already, so csift search -t user.answer (one of 25 {role}.{class}.{sub} labels) finds exactly what a naive grep swears was never said.
🛎 Pending questions, on the record.
csift ships a hook that records AskUserQuestion / ExitPlanMode / MCP elicitations to a sidecar file, and every csift surface merges the unresolved ones in transparently. An orchestrator can finally see which session is stuck waiting on a human, and on what.
A hit returns the whole exchange, rebuilt from the uuid / parentUuid graph: the matched tool call with its result, the user turn with the agent's reply. This is the context that grep or Claude's ad-hoc scripts can never reliably provide.
Kind, lifecycle, and the parent→child tree of every spawned agent, plus detection of lanes frozen on a pending permission approval.
🤖 Designed for humans and LLMs.
Output is terse, re-feedable, and even structural. Simply install the skill and your Claude will gain the power.
Pure regex. No embeddings, no index, no database, no daemon, no network, no telemetry, no hidden detections. It reads files already on your disk and never mutates your session histories.
⚡ Rust + mmap + SIMD newline scan + byte prefilters + rayon.
200 MB transcripts and multi-GB corpora in about a second , quick enough to call from inside a hook without noticing.
cargo install csift
Requires Rust 1.89+ . Or from source:
git clone https://github.com/wdhwg001/csift.git
cd csift
cargo install --path . # builds the optimized binary and puts `csift` on your PATH
# …or just `cargo build --release` → ./target/release/csift
csift reads ~/.claude by default; point it elsewhere with --claude-home <DIR> or Claude Code's own $CLAUDE_CONFIG_DIR . csift <command> --help is the full manual.
csift 's primary user is the agent itself : output is terse, parseable, and every record carries a re-feedable handle. The skill teaches Claude Code when and how to reach for it:
npx skills add wdhwg001/csift
Quickstart
csift <command> [TARGET] [flags] . A target is a positional @<uuid> (a session), an @<agent-hex> (a subagent), a project path, or . (this cwd); omit it to scan every project.
Run csift <command> --help for the full flag set and examples.
list
fast "which session is this?" index: first/last user + last agent, per session
search
regex over transcripts → the complete round-trip per hit ( -t / -T label filters, -l matching sessions, --raw verbatim lines)
show
fetch the exact record(s) you name: --line N|A..B / --uuid U / --turn N|A..B|-k (the ·tN turn index from search's headers) of one transcript, rendered full or --raw bytes
stats
one-scan aggregates per session: tokens by model, tool calls, turns, span, compactions
agents
a session's subagents: kind, lifecycle, status, and the parent→child topology
whoami
identify the calling session from $CLAUDE_CODE_SESSION_ID , false-positive-safe
files
which files/dirs a session changed, when, plus edits made outside the tool stream
recover
reconstruct a file (or a deleted plan) from the Read/Write/Edit stream: byte-exact or honest gaps
plan
locate the Plan-Mode plan file bound to a session (and reverse: which session owns a plan)
verbatim
restore the verbatim turns a compaction summary clipped, within a budget (the live-tail peek is show --turn )
image
list + extract images pasted into a transcript (handle/locator addressing, format transcode)
The summary is a selection. csift keeps the conversation.
When Claude Code compacts, it regenerates a dense summary: kilobytes standing in for dozens of compactions of history. It's a good selection. It keeps the key findings (what changed, with file:line) and often your standing directives verbatim. But it is a selection, re-abstracted every time, and the axis it optimizes is task continuation , not the conversation.
csift verbatim keeps the other axis: the verbatim User↔Agent exchange , what you actually said and what the agent actually reported back when it finished, with the hundreds of tool calls in between collapsed to a count. For the recent window that's tens of KB of full-fidelity dialogue where the summary kept a line or two. Not just your directive, but the agent's "validated against HEAD, the premise holds, here's the evidence" report-back, which a task-findings selection simply doesn't carry.
It doesn't replace the summary. It extends the post-compaction context with the recent conversation at full fidelity, budget-bounded and newest-first: it won't reach the oldest turns, and your never-compacted Plan file still owns the plan. Wire it into a SessionStart(compact) hook (see SKILL.md ) and every compaction arrives with the recent verbatim conversation attached.
csift never loads a whole transcript when it doesn't have to: it mmaps each .jsonl , scans newlines with SIMD ( memchr ), runs a cheap byte/regex prefilter , and only serde_json -parses candidate lines; list reads just the head and tail; rayon fans out across files. The hard part isn't speed. It's the semantics of Claude Code's log: a role:"user" record is usually a tool result, not a human turn; a pending question is never written to disk; compaction has a specific shape; subagents sit flat on disk and their topology is reconstructed from the spawning tool call. All of it is documented, with the empirical grounding, in SPEC.md .
There is none, in the human sense. This repository is written and maintained entirely by Claude Code, and the three reference files are the corpus the maintaining agent works from:
SKILL.md : the usage reference an agent loads to operate csift.
SPEC.md : design intent, record-model semantics, and the measurements behind the deep features.
AGENTS.md : the repo operating manual. Architecture, jsonl domain knowledge, the quality gate.
They are dense, cross-referenced, and written for a model's attention, not your

[truncated]
