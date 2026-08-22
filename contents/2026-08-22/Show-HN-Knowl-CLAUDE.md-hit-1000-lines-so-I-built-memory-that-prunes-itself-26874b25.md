---
source: "https://github.com/dat999zx/knowl"
hn_url: "https://news.ycombinator.com/item?id=49399942"
title: "Show HN: Knowl – CLAUDE.md hit 1000 lines, so I built memory that prunes itself"
article_title: "GitHub - dat999zx/knowl: Knowledge Operating System for AI agents · GitHub"
image: "https://repository-images.githubusercontent.com/1279171235/7b01d37f-9e44-4b8b-a533-4b10e62e01c7"
author: "dat999zx"
captured_at: "2026-08-22T15:11:57Z"
capture_tool: "hn-digest"
hn_id: 49399942
score: 1
comments: 1
posted_at: "2026-08-22T14:13:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Knowl – CLAUDE.md hit 1000 lines, so I built memory that prunes itself

- HN: [49399942](https://news.ycombinator.com/item?id=49399942)
- Source: [github.com](https://github.com/dat999zx/knowl)
- Score: 1
- Comments: 1
- Posted: 2026-08-22T14:13:22Z

## Translation

タイトル: Show HN: Knowl – CLAUDE.md が 1000 行に達したので、自動的にプルーニングするメモリを構築しました
記事のタイトル: GitHub - dat999zx/knowl: AI エージェント用のナレッジ オペレーティング システム · GitHub
説明: AI エージェント用のナレッジ オペレーティング システム。 GitHub でアカウントを作成して、dat999zx/knowl の開発に貢献してください。
HN text: AI エージェントの使用は素晴らしいですが、長期的には、AI エージェントのコンテキストに関してますます問題を抱え始めました。彼らは私たちが取り組んでいることを忘れ続け、新しいチャットセッションを開くとすべてのコンテキストが消去され、私のCLAUDE.mdには1000行ほどありました。そこで、市販されているエージェントメモリを取り付けてみることにしました。最初はうまくいきましたが、それらのメモリは追加されるだけで、古いものは修正されないことに気づきました。初日、MoR として Lemon Squeezy を使用するように指示しましたが、2 日目には Polar に変更するように指示しました。しかし、もう一度尋ねると、両方の答えが返され続けるか、どちらを使用しているのか判断できません。それが私がKnowlを作成した理由です。 Knowlは知識の鮮度で問題を解決します。知識を「アトム」と呼ばれる小さなデータビットに分割します。アトムには、ファクト、決定、目標、制約、アーキテクチャ、状態、スキルのタイプがあるため、カテゴリごとにアトムを取得できます。書き込み時に競合が発生すると (新しいアトムが古いアトムと競合する)、Knowl は古いアトムをリタイアし (置き換え済みのフラグを立てて)、メインの取得から削除しますが、完全な履歴は引き続き保持します。トランスクリプト検索、マルチワークスペース共有、変更検出の影響など、他にも多くの優れた機能があり、チーム同期用の Knowl Cloud もあります。 MemoryAgentBench - FactConsolidation シングルホップ @262K コンテキストで Knowl のベンチマークを実行したところ、驚くべき結果が得られました。 - Knowl: 0.90 <- これを実行しました - Agentmemory: 0.79 <- そしてこれ - Gpt-4o (フル コンテキスト): 0.60 - Mem0: 0.18 - Zep: 0.07 マルチホップでは 0.07 というスコアでした(セイ

史上最高のリングは 0.14)
気温0.7度で走りました。完全なベンチマークは私のリポジトリにあります。これは完全にオープンソースであり、MCP (Claude Code、Codex、Cursor、Antigravity...) を通じてほぼすべてのプロバイダーに接続します。また、完全にローカルです (Knowl Cloud を使用しない限り)。フィードバックをいただければ幸いです。乾杯！

記事本文:
GitHub - dat999zx/knowl: AI エージェント用のナレッジ オペレーティング システム · GitHub
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
dat999zx
/
知っている
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
1,019 コミット 1,019 コミット フォルダーとファイル
.claude-plugin .claude-plugin .github .github ベンチマーク ベンチマーク ドキュメント ドキュメント統合/ クライン統合/ クライン スクリプト スクリプト src src テスト テスト .gitattributes .gitattributes .gitignore .git

無視 AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md KNOWL.md KNOWL.md ライセンス ライセンス README.md README.md drizzle.config.ts drizzle.config.ts eslint.config.mjs eslint.config.mjs Glama.json Glama.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts vitest.mutation.config.ts vitest.mutation.config.ts すべてのファイルを表示リポジトリ ファイルのナビゲーション
CLAUDE.md は成長するだけです。 Knowl は、事実が変わると事実を廃止します。
クイックスタート ·
なぜスーパーセッションなのか・
保存されるもの ·
特徴・
エージェントのセットアップ ·
ビューア ·
要件 ·
完全なリファレンス→
エージェントはすべてのセッションを空白から開始するため、 CLAUDE.md を保持します。それは成長するだけです。 6ヶ月間
去年の春に移行したデータベースに名前を付けたままですが、エージェントは両方の答えを得ることができました。
Knowl は、クロード コード、カーソル、およびコーデックスの永続的なメモリです。
MCP または CLI。ファクトが置き換えられると、古いファクトは廃止されます
新しいものと競合するのではなく。 API キーは必要ありません。 Knowl が新しい事実を確信していない場合
古いものを置き換えると、両方がアクティブのままになり、そうするための knowl supersede コマンドが渡されます。
これをオフにすると、取得率が 98% から 47% に低下します。エンドツーエンド、90対73。
測った様子↓
40 秒、1 つの決定、3 人のエージェント:
Node.js 22 以降が必要です。 macOS、Linux、Windows。
npm install -g @dat999zx/knowl
プロジェクトの cd
初期化を知る
他のパッケージマネージャー
公開されたパッケージはどの場合でも同じです。これらのそれぞれがそれをインストールし、知識を与えます
PATH 上にあります。
pnpm add -g @dat999zx/knowl
糸グローバル追加 @dat999zx/know
bun add -g @dat999zx/knowl
または、インストールせずに実行します。
npx @dat999zx/knowl init
Knowl はこれらすべてで Node.js 上で実行されます。Bun はそれをインストールします。

ノードがそれを実行します。ネイティブをバンドルします
アドオン (SQLite、ツリーシッター、埋め込みランタイム) なので、Bun または Deno で CLI を実行します。
ランタイムを直接サポートすることはできません。
knowl init は .knowl/ を作成し、プロジェクト ガイダンス ファイルをインストールし、 .gitignore を更新します。
検出したエージェントに Knowl を登録します。また、ローカル埋め込みモデル (~53 MB) もウォームします。
バックグラウンドで — init はいずれの方法でも成功しますが、init がなくてもキーワード検索は可能です。
それが全体のセットアップです。手動でメモリを記録するのではなく、エージェントがメモリを読み書きします。
動作します。
クロード・コード
MCP・ライフサイクル・ゲート
コーデックス
MCP・ライフサイクル・ゲート
副操縦士
MCP・ライフサイクル・ゲート
カーソル
MCP・ライフサイクル・ゲート
オープンハンズ
MCP・ライフサイクル・ゲート
反重力
MCP・ライフサイクル・ゲート
ウィンドサーフィン
MCP・ライフサイクル・ゲート
クライン
MCP・ライフサイクル・プラグイン
ゼッド
MCP・キャプチャ・ACP
ジェットブレインズ
MCP・キャプチャ・ACP
オープンコード
MCP・マニュアルループ
クロードデスクトップ
MCP・マニュアルループ
knowl init は、見つかったすべてのホストの MCP サーバーを登録します。その後、新しいセッションを開始します。
エージェントはそのガイダンスを受け取り、独自にメモリにクエリを実行し、書き込みます。
ゲートとは、Knowl が別のセッションが保持しているコードを無効にする編集を拒否できることを意味します。
Neovim と Kiro は、 knowl acp を通じて Zed や JetBrains と同じように機能します。クラインにはそれが必要です
同梱のプラグインを指す行。他の MCP クライアントはまったく統合せずに動作します。
→ すべてのホストと各ホストができること · エージェントによる使用方法 · MCP ツールとリソース
アイデア: 自動的に引退する記憶
ほとんどのメモリ システムは追加専用です。 「SQLite に移行しました」を保存すると「PostgreSQL を使用します」のままになります
アクティブで取得可能なため、エージェントは両方を取得し、ランクによって選択します。 Knowl は同じ件名の書き込みを扱います
訂正: 先行者は superseded とマークされ、通常の取得から外れ、

滞在します
既知のタイムラインを通じてクエリ可能。
この 1 つの動作が精度の違いのほとんどです。で
MemoryAgentBench 競合解決コーパス —
455 の事実、どの事実が最新であるかに関する 100 の質問、トップ 5 の検索、LLM リーダーなし:
同じコーパス、同じランカー、同じクエリ パス。唯一の変数は、古い事実がまだ残っているかどうかです
アクティブな。これは、Knowl 独自のハーネスにおける検索レベルの測定です。
現在のファクトが最初に返され、ループ内にはモデルはありません。
ベンチマーク独自のハーネスでエンドツーエンドで検証済み
自分で得点した数字は他の人が得点した数字よりも価値が低いため、同じ主張が行われました。
MemoryAgentBench のハーネス内で再実行し、独自のコードによってスコア付けされ、LLM が何を読み取るか
返された Knowl は、タスクが提供する最大のコンテキストでの、より困難で完全なエンドツーエンドのセットアップです。
18,332 の事実、100 の質問、部分文字列の完全一致。すべての行はリーダーとして gpt-4o-mini を使用します。
Knowl も含まれています。論文にはすべての RAG エージェントとメモリ エージェントについて記載されているため、これらは同様です。
ここでは、Knowl とエージェントのメモリが測定されました。他のすべての図は MemoryAgentBench からのものです
紙、arXiv 2507.05257v4、表 3. エージェントメモリはありません
その論文で評価された - その公表された数値は LongMemEval-S 検索再現率であり、別の
タスク - したがって、同じ構成の同じハーネスを通じて実行され、両方のアダプターが 1 つのアダプターを共有します
リーダー コード パスを使用するため、どちらも論文独自の RAG ハンドラーから逸脱することはできません。方法や仕組み、
再現手順: FINDINGS.md 。
それ以外の場合は、この論文で評価したすべての商用メモリ システムと最高得点者を示します。
各ベースラインファミリーから。論文の表はバージョン間で変更されています - BM25 は v1 で 56 を読み取り、
v4 では 48 と読み取られます。つまり、表だけでなくバージョンが引用されています。
Knowl's 90は2026年8月8日に測定され、独立しています

2026 年 8 月 19 日の 89.0 で、
チェックインされたアダプター。 Agentmemory の 79 は 1 回の実行です。ここにあるすべての数値は 1 回の実行に相当します
温度: 0.7 、同じ 6k セルの 2 つの実行間でアブレーション ギャップが 4 ポイント移動したため、
小数点までではなく、要点まで読んでください。
同じハーネスでスーパーセッションをオフにすると、Knowl が 73 に低下し、そのギャップは
コーパスサイズの 40 倍の変化:
2 つのセクションは異なるものを測定するため、互いに比較できません: 98% は検索です
リーダーなしの場合は 6K でトップ 1、リーダーありの場合は 262K でエンドツーエンドの精度は 90 です。 2番目だけは
上記の公開されているシステムと同等です。のベンチマークを参照してください。
プロトコル、チェックイン結果、およびタスクでカバーされない内容 (マルチホップを含む)
ノールは 14 ポイントのリトリーブ上限に対して 7 ポイントを獲得しました。
置き換えは削除ではなく修正です。アイテム、そのアサーション、およびその履歴はすべて存続します。
モックアップではありません - 公開された CLI に対する同じシーケンスであり、
デモ.テープ:
チーム全体でメモリを共有する: knowl.cloud
上記はすべてローカルであり、アカウントは必要ありません。 knowl.cloud はオプションです
1 台のマシンでは不十分な場合のホスト層:
共有ワークスペース。 1 回のチェックアウトで書き込まれたナレッジがチームメイトのエージェントに届き、それぞれの
リポジトリは公開したものをまだ所有しています。
ブラウザエージェント。 claude.ai と chatgpt.com はローカル プロセスを実行できないため、
1 つのワークスペースをスコープとするトークンを持つリモート MCP エンドポイント。
ローカルのみは、Knowl を実行するための第一級の方法のままです。上記のものを使用するためにここで必要なものは何もありません。
すべての原子には、次の 7 つのカテゴリのうちの 1 つが正確にあります。
コンテンツに加えて、各アトムはステータス ( active 、 deprecated 、拒否された、 archived 、
置き換えられた )、鮮度フラグ、信頼性、タグ、ソースコミット、影響を受けるパス、およびオプション
を指す証拠

ファイル、コミット、テスト、コマンド、URL、またはインデックス付きコード シンボル。ファイルと
コードが移動すると、記号の証拠は自然に古くなります。これが、アトムが自分がそうかもしれないことを認める方法です。
もう存在しないリポジトリのバージョンをアサートするのではなく、古いものにします。
Knowl が意図的に保存しないのは、あなたの会話です。ライフサイクル キャプチャ レコードの制限付き
イベントと概要 - プロンプト、トランスクリプト、標準出力、または環境変数は使用しません。生の転写物
検索はオプトインのデフォルトでオフのインデックスとして存在します
ホストがすでに書き込んだファイルを上書きします。
knowlserve は stdio MCP 経由でストアを公開します。 knowl init がそれを登録します。ワークフローは、
エージェントに従うように求めるインストールされたガイダンスは短いです。
リポジトリ ファイルを読み取る前に、主題の名前を表す単語でメモリをクエリします。
アクティブなヒットを直接使用します。ミス、競合、または古い結果があった場合にのみファイルを検査します。
永続的な所見、明示された目標、繰り返し発生する診断を保存し、修正します。
記憶を複製するのではなく、矛盾した記憶。
実際には次のようになります。新しいセッション、コンテキストなし、何も貼り付けられていません。
なぜ Postgres ではなく SQLite を選んだのですか?
エージェント → knowl_query "sqlite postgres データベースの選択"
← 決定 · SQLite を使用 · アクティブ · フレッシュ
「ストレージ リポジトリをローカルに保ち、操作を簡単にします。」
代替手段: PostgreSQL、MongoDB
タグ: データベース、ローカルファースト
SQLite は、ストア リポジトリをローカルに保ち、操作を簡単にします。
Postgres と MongoDB はどちらも検討されましたが、却下されました。
基礎。
エージェントは 1 つのファイルを開く前に応答し、ユーザーが拒否したオプションを認識していました。
拒否された代替案はコードベースに痕跡を残さないため、コードではそれを判断できません。
詳細と各ギャップが存在する理由については、 docs/hosts.md を参照してください。
フックが利用可能な場合、フックはセッションのライフサイクルを所有します: ブートストラップ コンテキスト、キャプチャ、チェックポイント、

そしてファイナライズはエージェントに尋ねることなく行われます。そうでない場合は、タスクの実行を知ってください。
タスクの開始、タスクのチェックポイント、タスクの終了は同じ領域を手動でカバーします。
knowl init は、検出したすべてのホストの MCP 登録を書き込みます。手動で配線するには、
エントリはどこでも同じです:
{
"mcpサーバー": {
"knowl" : { "command" : " knowl " , "args" : [ "serve " ] }
}
}
Windows ではコマンドとして knowl.cmd を使用します。 Codex は、 mcp_servers の下の同じエントリを読み取ります。
→ MCP ツールとリソース · ライフサイクル リファレンス
Knowl は 1 つの仕事をします。それは、リポジトリで作業するエージェントに対してリポジトリのエンジニアリングの真実を正確に保つことです。
ユーザーの好みやチャット履歴ではなく、意思決定、制約、アーキテクチャです。
コードベース、そしてそれらのうちどれが現在でも真実であるか。
そこから 3 つの選択肢が導き出されます。
フリーテキストではなく、入力されたものです。決定には理由があり、拒否した選択肢も含まれます。あ
制約は保持し続けなければならないルールです。状態アトムは期限切れになることが予想されます。
検索はそれらの違いに基づいてランク付けできます。ノート ファイル内の段落にランクを付けることはできません。
追加専用ではなく管理されています。ステータス、鮮度、来歴、競合アイデンティティ、およびスーパーセッション
ストアに、何かが真実ではなくなったと伝えてもらいましょう。それが両者の違いです
記憶と増え続けるメモの山。
リポジトリ-

[切り捨てられた]

## Original Extract

Knowledge Operating System for AI agents. Contribute to dat999zx/knowl development by creating an account on GitHub.

Using AI agents is amazing, but in long-term, I started to have more and more problems about the context of them. They keep forgetting what we are working on, opening a new chat session clears all their context, my CLAUDE.md had like 1000 lines. So I thought to install those agent memory that are on the market. Started out great, but I noticed that those memory only append and does not fix what is stale. First day I told it to use Lemon Squeezy as our MoR but second day I tell it to change to Polar. But when I ask it again, they keeps returning both answers or cannot decide which we are using. That's why I created Knowl. Knowl solves the problem with freshness of the knowledge. It split the knowledge into small data bits called "atom". Atoms can be of the following types: fact, decision, goal, constraint, architecture, state, skill so we can retrive atoms in categories. When conflict happens at write-time (a new atom conflicts with an old one), Knowl retires the old one (flag it with superseded) and remove it out of main retrieval but still keeps full history. There are many more cool features like transcript search, multi-workspace sharing, change detection impact... and Knowl Cloud for team-sync too. We benchmarked Knowl on MemoryAgentBench - FactConsolidation single-hop @262K context and got suprising result: - Knowl: 0.90 <- I ran this - agentmemory: 0.79 <- and this - Gpt-4o (full context): 0.60 - Mem0: 0.18 - Zep: 0.07 In multi-hop we scored 0.07 (ceiling of all time is 0.14)
I ran at temperature 0.7. You can find full benchmark in my repo. This is fully open-source, connect to almost every providers through MCP (Claude Code, Codex, Cursor, Antigravity...). And it is fully local (unless you use Knowl Cloud). I'd love to get some feedback. Cheers!

GitHub - dat999zx/knowl: Knowledge Operating System for AI agents · GitHub
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
dat999zx
/
knowl
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,019 Commits 1,019 Commits Folders and files
.claude-plugin .claude-plugin .github .github benchmarks benchmarks docs docs integrations/ cline integrations/ cline scripts scripts src src tests tests .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md KNOWL.md KNOWL.md LICENSE LICENSE README.md README.md drizzle.config.ts drizzle.config.ts eslint.config.mjs eslint.config.mjs glama.json glama.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts vitest.mutation.config.ts vitest.mutation.config.ts View all files Repository files navigation
Your CLAUDE.md only grows. Knowl retires facts when they change.
Quick start ·
Why supersession ·
What gets stored ·
Features ·
Agent setup ·
Viewer ·
Requirements ·
Full reference →
Your agent starts every session blank, so you keep a CLAUDE.md . It only grows. Six months in it
still names the database you migrated off last spring, and now the agent gets both answers.
Knowl is persistent memory for Claude Code, Cursor and Codex, over
MCP or the CLI. When a fact is replaced, the old one is retired
instead of competing with the new one. No API key needed. When Knowl isn't sure the new fact
replaces the old, it leaves both active and hands you the knowl supersede command to say so.
Turn that off and retrieval drops from 98% to 47%. End to end, 90 to 73.
How it was measured ↓
Forty seconds, one decision, three agents:
Requires Node.js 22 or later. macOS, Linux and Windows.
npm install -g @dat999zx/knowl
cd your-project
knowl init
Other package managers
The published package is the same one in every case; each of these installs it and puts knowl
on your PATH .
pnpm add -g @dat999zx/knowl
yarn global add @dat999zx/knowl
bun add -g @dat999zx/knowl
Or run it without installing:
npx @dat999zx/knowl init
Knowl runs on Node.js in all of these — Bun installs it, Node executes it. It bundles native
addons (SQLite, tree-sitter, the embedding runtime), so running the CLI under the Bun or Deno
runtime directly is not supported.
knowl init creates .knowl/ , installs the project guidance files, updates .gitignore , and
registers Knowl with whichever agents it detects. It also warms a local embedding model (~53 MB)
in the background — init succeeds either way, and without it you still get keyword search.
That is the whole setup. You do not record memory by hand: your agent reads and writes it as it
works.
Claude Code
MCP · lifecycle · gate
Codex
MCP · lifecycle · gate
Copilot
MCP · lifecycle · gate
Cursor
MCP · lifecycle · gate
OpenHands
MCP · lifecycle · gate
Antigravity
MCP · lifecycle · gate
Windsurf
MCP · lifecycle · gate
Cline
MCP · lifecycle · plugin
Zed
MCP · capture · ACP
JetBrains
MCP · capture · ACP
OpenCode
MCP · manual loop
Claude Desktop
MCP · manual loop
knowl init registers the MCP server for every host it finds. Start a new session afterwards so
the agent picks up its guidance, and it will query and write memory on its own.
gate means Knowl can refuse an edit that invalidates code another session is holding.
Neovim and Kiro work the same way as Zed and JetBrains, through knowl acp . Cline needs one
line pointing it at the shipped plugin. Any other MCP client works with no integration at all.
→ Every host, and what each one can do · How agents use it · MCP tools and resources
The idea: memory that retires itself
Most memory systems are append-only. Storing "we moved to SQLite" leaves "we use PostgreSQL"
active and retrievable, so the agent gets both and picks by rank. Knowl treats a same-subject write
as a correction: the predecessor is marked superseded , drops out of normal retrieval, and stays
queryable through knowl timeline .
That single behavior is most of the accuracy difference. On the
MemoryAgentBench Conflict Resolution corpus —
455 facts, 100 questions about which fact is current, top-5 retrieval, no LLM reader:
Same corpus, same ranker, same query path. The only variable is whether the outdated fact is still
active. This is a retrieval-level measurement in Knowl's own harness: it asks whether the
current fact comes back first, with no model in the loop.
Verified end-to-end, in the benchmark's own harness
Because a number you score yourself is worth less than one somebody else scores, the same claim was
re-run inside MemoryAgentBench's harness, scored by its own code , with an LLM reading what
Knowl returned — the harder, fully end-to-end setup, at the largest context the task offers:
18,332 facts, 100 questions, substring exact match. Every row uses gpt-4o-mini as the reader ,
Knowl's included — the paper states it for all RAG and memory agents, so these are like-for-like.
Knowl and agentmemory were measured here; every other figure is from the MemoryAgentBench
paper , arXiv 2507.05257v4 , Table 3. agentmemory is not
evaluated in that paper — its published numbers are LongMemEval-S retrieval recall, a different
task — so it was run through the same harness with the same config, and both adapters share one
reader code path so neither can drift from the paper's own RAG handler. Method, mechanism and
reproduction steps: FINDINGS.md .
Otherwise shown are every commercial memory system the paper evaluates, plus the highest scorer
from each baseline family. The paper's table has changed between versions — BM25 read 56 in v1 and
reads 48 in v4 — so the version is cited, not just the table.
Knowl's 90 was measured 2026-08-08 and independently reproduced at 89.0 on 2026-08-19 with the
checked-in adapter; agentmemory's 79 is a single run. Every figure here is one run at
temperature: 0.7 , and the ablation gap moved 4 points between two runs of the same 6k cell, so
read them to the point rather than the decimal.
Switching supersession off in that same harness drops Knowl to 73 , and the gap holds across a
40× change in corpus size:
The two sections measure different things and are not comparable to each other: 98% is retrieval
top-1 at 6K with no reader, 90 is end-to-end accuracy at 262K with one. Only the second is
comparable to the published systems above. See benchmarks for the
protocol, the checked-in results, and what the task does not cover — including multi-hop, where
Knowl scores 7 against a 14-point retrieval ceiling.
Supersession is a correction, not a delete: the item, its assertions, and its history all survive.
Not a mock-up — the same sequence against the published CLI, recorded from
demo.tape :
Sharing memory across a team: knowl.cloud
Everything above is local and needs no account. knowl.cloud is the optional
hosted layer for when one machine is not enough:
Shared workspaces. Knowledge written in one checkout reaches teammates' agents, with each
repository still owning what it publishes.
Browser agents. claude.ai and chatgpt.com cannot run a local process, so they connect over a
remote MCP endpoint with a token scoped to one workspace.
Local-only remains a first-class way to run Knowl. Nothing here is required to use anything above.
Every atom has exactly one of seven categories:
Alongside the content, each atom keeps a status ( active , deprecated , rejected , archived ,
superseded ), a freshness flag, confidence, tags, source commit, affected paths, and optional
evidence pointing at files, commits, tests, commands, URLs, or indexed code symbols. File and
symbol evidence go stale on their own when the code moves, which is how an atom admits it may be
out of date instead of asserting a version of the repository that no longer exists.
What Knowl deliberately does not store is your conversations. Lifecycle capture records bounded
events and summaries — never prompts, transcripts, stdout, or environment variables. Raw transcript
search exists as an opt-in, off-by-default index
over files the host already wrote.
knowl serve exposes the store over stdio MCP; knowl init registers it for you. The workflow the
installed guidance asks agents to follow is short:
Query memory with the words that name the subject before reading repository files.
Use an active hit directly; inspect files only on a miss, conflict, or stale result.
Store durable findings, stated goals, and recurring diagnoses as you go, and correct
contradicted memory rather than duplicating it.
In practice that looks like this — a new session, no context, nothing pasted in:
You why did we pick SQLite over Postgres?
Agent → knowl_query "sqlite postgres database choice"
← decision · Use SQLite · active · fresh
"Keeps storage repository-local and simple to operate."
alternatives: PostgreSQL, MongoDB
tags: database, local-first
SQLite keeps the store repository-local and simple to operate.
Postgres and MongoDB were both considered and rejected on that
basis.
The agent answered before opening a single file, and it knew the options you rejected —
which the code cannot tell it, because rejected alternatives leave no trace in a codebase.
Full detail, and why each gap exists, in docs/hosts.md .
Where hooks are available, they own the session lifecycle: bootstrap context, capture, checkpoints,
and finalization happen without the agent being asked. Where they are not, knowl task run ,
task start , task checkpoint , and task finish cover the same ground manually.
knowl init writes the MCP registration for every host it detects. To wire one by hand, the
entry is the same everywhere:
{
"mcpServers" : {
"knowl" : { "command" : " knowl " , "args" : [ " serve " ] }
}
}
Use knowl.cmd as the command on Windows. Codex reads the same entry under mcp_servers .
→ MCP tools and resources · Lifecycle reference
Knowl does one job: keep a repository's engineering truth accurate for the agents working on it.
Not user preferences, not chat history — the decisions, constraints, and architecture of a
codebase, and which of them are still true today.
Three choices follow from that:
Typed, not free text. A decision carries reasoning and the alternatives you rejected. A
constraint is a rule that must keep holding. A state atom is expected to go out of date.
Retrieval can rank on those differences; it cannot rank on paragraphs in a notes file.
Governed, not append-only. Status, freshness, provenance, conflict identity, and supersession
let the store tell you that something stopped being true. That is the whole difference between
memory and an ever-growing pile of notes.
Repository-

[truncated]
