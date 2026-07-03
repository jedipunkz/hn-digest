---
source: "https://github.com/YuanpingSong/ultracodex"
hn_url: "https://news.ycombinator.com/item?id=48776386"
title: "Show HN: Ultracodex – Run Claude Ultracode Dynamic Workflows with Codex Agents"
article_title: "GitHub - YuanpingSong/ultracodex: Run Claude Code workflow scripts, unmodified, on the OpenAI Codex CLI — fable plans, codex executes, fable verifies. Parallel agent fleets, builder–verifier loops, token budgets, full-screen TUI. · GitHub"
author: "yuansong"
captured_at: "2026-07-03T15:49:36Z"
capture_tool: "hn-digest"
hn_id: 48776386
score: 1
comments: 0
posted_at: "2026-07-03T15:45:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ultracodex – Run Claude Ultracode Dynamic Workflows with Codex Agents

- HN: [48776386](https://news.ycombinator.com/item?id=48776386)
- Source: [github.com](https://github.com/YuanpingSong/ultracodex)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T15:45:01Z

## Translation

タイトル: 表示 HN: Ultracodex – Codex エージェントを使用したクロード Ultracode 動的ワークフローの実行
記事タイトル: GitHub - YuanpingSong/ultracodex: OpenAI Codex CLI でクロード コード ワークフロー スクリプトを変更せずに実行します — 寓話の計画、コーデックスの実行、寓話の検証。並列エージェント フリート、ビルダーと検証者のループ、トークン バジェット、全画面 TUI。 · GitHub
説明: OpenAI Codex CLI でクロード コードのワークフロー スクリプトを変更せずに実行します。寓話の計画、コーデックスの実行、寓話の検証を行います。並列エージェント フリート、ビルダーと検証者のループ、トークン バジェット、全画面 TUI。 - YuanpingSong/ultracodex
HN テキスト: Claude Fable は素晴らしかったですが、特にウルトラコード モード (Claude Code のワークフロー機能。モデルがサブエージェントを調整する小さな JavaScript プログラムを作成する) を使用し、エージェントを解放した場合、プランの使用量があまりにも早くなくなります。 Fable には、日常的な実装作業でトークンを消費するのではなく、計画や検証などの価値の高いタスクに集中してもらいたいと考えています。そこで私は、同じ Ultracode ワークフローを完了し、Claude との間でシームレスに引き継ぎを行うコーデックス エージェントを生成するエンジン、ultracodex を構築しました。クロードが作成したものと同じワークフロー スクリプトを実行します。このツールは、Claude と Codex の両方のサブスクリプションを所有している人、または Fable を楽しんでクォータを持続させたい人のために設計されました。このツールはまだ初期段階にありますが、その可能性に非常に興奮しています。また、これは私たちにとって素晴らしい「ループ エンジニアリング」スターター パックでもあると思います。1) Claude はループの作成が得意で、2) Ultracodex は Codex エージェントを使用してループを実行します。制限が十分にあるため、使用量についてあまり心配する必要はありません。したがって、「ループ エンジニアリング」はもはや神秘的なものではなく、あらゆるタスクに投入できるツールです。私のリポジトリにも「ループ」の例がいくつかあります。プロジェクトは完全にオープンソースです

レース。 Apache-2.0、npm i -g Ultracodex、README のデモ ビデオ。ぜひチェックしてみてください！

記事本文:
GitHub - YuanpingSong/ultracodex: OpenAI Codex CLI でクロード コード ワークフロー スクリプトを変更せずに実行します — 寓話の計画、コーデックスの実行、寓話の検証。並列エージェント フリート、ビルダーと検証者のループ、トークン バジェット、全画面 TUI。 · GitHub
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
{{ メッセージ

}}
元平歌
/
ウルトラコーデックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .github/ workflows .github/ workflows .ultracodex .ultracodex アセット アセット ドキュメント ドキュメントの例 例 フィクスチャ フィクスチャ スクリプト スクリプト src src テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI Codex CLI で、Claude Code ワークフロー スクリプトを変更せずに実行します。
Claude Code のワークフロー ツールには優れたオーケストレーション形式があります: plain-JS
エージェントを並行して扇動するスクリプト、ステージを介したパイプライン作業、
トークンの予算を強制し、結果を敵対的に検証します。しかし、それらを実行する
上流では、クロード コード自体の下で、クロード クォータを次のことに費やします。
実行レベルの作品。
Ultracodex は、これらの同じスクリプトと互換性のあるランタイムです (バイト同一)。
トランスパイルなし — 各agent()呼び出しをルーティングします。
代わりにコーデックス (または設定されたコーデックス)
バックエンド、ラベルごと）。存在するパターン:
寓話の計画、コーデックスの実行、寓話の検証
(寓話 = クロードのフロンティア モデル。お好みのプランナーを代用してください)
最も有能なモデルが作品を作成し、審査します。コーディングが安くて速い
エージェントがその大部分を行います。 1 つの台本、1 つのジャーナル、1 つの予算。
クロード・コード (右) は、「人生の意味についてのエッセイを書いてください — 俳優と批評家のループ、3 ラウンド。ultracodex で実行してください。」と尋ねられます。ワークフローを作成して開始し、TUI (左) が Codex エージェントによるループの実行を監視し、結果の JSON が Claude に返されます。 1.3倍のスピードアップ。 (HDビデオ)
割り当て量

アービトラージ。実行を Codex (または別のバックエンド) にオフロードします。
クロードを裁きに委ねる。実装に費やされるクロード トークンはゼロです。
ベンダー間の検証。批判:* を別のモデルにルーティングします
仕事をした人よりも家族。同一ベンダーのセルフレビューは次の傾向にあります。
ゴム印。ジャッジを別のモデル ファミリにすることも 1 つの構成です
ここのライン。
耐久性があり、検査可能な走行。デーモンはありません。実行は分離されたプロセスであり、
プレーンファイルのディレクトリ。追加専用ジャーナルは、
TUI、CLI、マシンの消費者にとっても同様に真実です。を閉じます
ターミナル。実行は関係ありません。
スクリプトは ES モジュールです。つまり、純粋なリテラルのメタ エクスポート、次にプレーンな JS です。
8 つの注入されたグローバルに対する非同期ボディ。インポートも TypeScript もありません。
エクスポート const メタ = {
名前: 'レビューファイル' 、
description : 'レビュー担当者をファンアウトし、調査結果を検証し、報告します' ,
フェーズ: [ { title : 'レビュー' } 、 { title : '検証' } ] 、
}
const FILES = args ?。ファイル?? [ 'src/auth.ts' 、 'src/api.ts' 】
フェーズ ( 'レビュー' )
const 結果 = ( await 並列 ( FILES .map ( f => ( ) =>
Agent ( `${ f } でバグを確認してください。スキーマ経由で返します。` , {
ラベル: `レビュー: ${ f } `、
スキーマ : { タイプ : 'オブジェクト' 、プロパティ : {バグ : { タイプ : '配列' 、項目 : { タイプ : '文字列' } } }、必須 : [ 'バグ' ] } 、
})
) ) ) 。 filter ( Boolean ) // 失敗したエージェントは null であり、スローされません
フェーズ ( '検証' )
const verify = await Pipeline ( // バリアなし - 各アイテムは独立してフローします
所見。 flatMap ( f => f . bugs ) 、
( bug ) => エージェント ( `反論してみます: ${ bug } ` , { label : 'verify' } ) 、
( verdict , bug ) => ( { bug , verdict } ) , // verdict が null の可能性があります (検証失敗) — チェックしてください
)
return {検証済み:検証済み。フィルター ( v => v && v . 判定 ) }
グローバル
それは何をするのか
エージェント(プロンプト、オプト?)
1 つのエージェントを実行します。最終テキスト、スキーマ va を解決します。

蓋付きオブジェクト、または失敗した場合は null (拒否されません - スローされる予算/上限を除く)
パラレル(サンク)
同時サンクに対する障壁。スローされたサンクは null になります
パイプライン(アイテム、...ステージ)
アイテムごとのステージチェーン、アイテム間のバリアなし。ステージ get (prev、item、index) ;最初のステージの prev はアイテム自体です。スローするステージは、その項目を null にドロップします。 null を解決するステージ (失敗したエージェントなど) は次に進みます。後のステージでは null チェックが必要です。
フェーズ(タイトル)
後続のエージェントの進行状況のグループ化
ログ(メッセージ)
TUI / --watch 出力のナレーター行
引数
実行の --args 入力、逐語的
予算
{ total, Spent(), Remaining() } — 出力トークンの上限。それを超えると、さらなるagent()呼び出しがスローされます
ワークフロー(名前、引数?)
保存されたワークフローをインラインで実行します (1 つのネストレベル)
Agent() オプション: ラベル (表示 + ルーティング)、フェーズ、スキーマ (JSON スキーマ)、
モデル/作業 (アドバイザリー層、構成にマップ)、分離: 'worktree'
(新鮮なギットを
[切り捨てられた]
スクリプトはプレーンな JavaScript であるため、ループはネイティブであり、ビルダー→ベリファイアーとなります。
ループは、このツールが簡単に作成できるように構築されたパターンです。ランタイムが与えるのは、
ガードレールをループします: ハード出力トークン バジェット (さらに Agent()
コールは使用済みになるとスローされます）、エージェントの有効期間上限は 1000、ライブです
実行ごとに一時停止/スキップ/停止のコントロールを実行します。
let Try = null 、フィードバック = 'なし — 最初の試行'
while ( Budget . Remaining ( ) > 20_000 ) { // レール 1: 予算ガバナー
試行 = エージェントを待ちます ( `ビルド X。レビューアーのフィードバック: ${ フィードバック } ` , { ラベル : 'ビルド' } )
if (attempt === null ) continue // レール 2: ビルダーは失敗する可能性があります
const verdict = await エージェント ( `これに反論してみてください: ${ 試行 } ` , {
label : 'verify:attempt' , // 構成内の別のベンダーにルーティングします
スキーマ : { タイプ : 'オブジェクト' 、プロパティ : { パス : { タイプ : 'ブール' } 、問題 : { タイプ : '配列' 、項目

: { type : 'string' } } }、必須 : [ 'pass' , 'issues' ] } 、
})
if ( verdict ?. pass ) Break // レール 2 をもう一度: 裁判官も失敗する可能性があります
フィードバック = 評決 ?評決。問題 。 join ( '; ' ) : '検証者が利用できません'
}
リターン { 試行 }
3 つの軸、1 つの形式:Parallel() は幅、pipeline() はフロー、
ループは深さです。予算を基準に、良好であることが検証されるまで繰り返します。
ガバナーと走行のライブ コントロール (一時停止/スキップ/停止) をブレーキとして使用します。の
フルパターン — 最大ラウンド、フィードバックスレッド、クロスベンダージャッジ — は
Examples/03-builder-verifier.js 。ルート
設定の "verify:*" = "claude" とビルダーの判断は異なります
モデル家族。ループは変わりません。
完全な規範的定義 - 文法、セマンティクス、ループ パターン、
適合性 — docs/agent_script_spec.md にあります。
同じファイルが Claude Code のワークフロー ツールと Ultracodex で実行されます。
Ultracodex validate --strict は、スクリプトがポータブル サブセットに残っているかどうかをチェックします。
前提条件: ノード ≥ 20、 pnpm 、および
Codex CLI がインストールされ、認証されている
(コーデックスログイン)。 codex-cli 0.142.4 に対して開発およびテスト済み
(ultracodex ドクターがあなたのバージョンを印刷します)。オプション: クロード コード
クロードに誘導されたエージェント。
npm install -g Ultracodex # または: pnpm add -g Ultracodex
Ultracodex Doctor # ノード、コーデックス、認証、構成をチェックします
代わりにソースから:
git clone https://github.com/YuanpingSong/ultracodex && cd Ultracodex
pnpm インストール && pnpm ビルド
pnpm link --global # → PATH の `ultracodex` (または、node dist/cli.js を使用し続ける)
サンプルはパッケージに同梱されています。最初にサンプルを実行します (チェックアウトから、または
$(npm root -g)/ultracodex/ ) 経由でインストールされたパッケージから:
Ultracodex run example/01-hello.js --watch # 1 つのエージェント、ストリーミング イベント
Ultracodex run example/02-fanout-critique.js # ファンアウト + スキーマ、TUI を開きます
ウルトラコーデックス実行例/03-builder-

verifier.js --watch --budget 200k # ループ
Ultracodex ls # 実行ごとに、pid-liveness をチェック
Ultracodex show < runId > --json # 機械可読な結果
実際の作業の場合は、プロジェクトのルートから実行します。エージェントは cwd で動作します。
run は切り離されたランナーを起動します。TUI を終了しても run が強制終了されることはありません。
--json はブロックして結果を出力します (これは駆動 LLM が呼び出すものです)。
--watch はイベントを 1 行ずつストリームします。 --detach は runId を出力するだけで、
出る。 --budget は、オプションの k / m を備えた整数として出力トークンを受け取ります
サフィックス ( --budget 500k )。
Claude Code からの Ultracodex の駆動
それが上記のデモで示されているもので、セットアップには 1 つのコマンドが必要です。
Ultracodex sync-skills # クロード コードの `ultracodex` スキルをインストールします
このスキルはクロードに契約全体を教えます - スクリプトを正確に作成します
ワークフロー ツールの場合は、ultracodex run <file> --json [--budget 500k] を使用して実行し、結果をそのまま伝えます。その後、プロンプトはただ
タスク:
人生の意味についてのエッセイを書きます — 俳優と批評家のループ、3 ラウンド。
ウルトラコーデックスで実行します。
スキルがない場合は、明示的にプロンプトを表示します。内容に応じて 2 つのバリエーションがあります。
セッションのコンテキスト内:
オプション A — スキーマネイティブ (セッションにはワークフロー ツール、例: Ultracode が含まれます)
モード;クロードは、そのツールのスキーマとして形式をすでに知っています)。
ワークフロー ツールの場合とまったく同じようにワークフローを作成します。
スクリプトをバイトごとに作成します。ただし、ツールを呼び出す代わりに、ツールを
ファイルを作成して、ultracodex を実行します。 run <file> --json --budget 300k 。中継する
結果はそのままの JSON です。実行が失敗した場合は、失敗を報告します。
オプション B — 仕様から教える (プレーンセッション、headless claude -p 、
他のエージェントは完全に）: オプション A として、ただし形式ドキュメントを指します —
「docs/agent_script_spec.md のエージェント スクリプト形式を使用する」 - と状態
タスク。 (このレポはこのようにして

それ自体が検証されました: 新鮮な首なしクロード
仕様に基づいて段階的ビルドのワークフローを作成し、これを実行しました。
Codex エージェントを使用してプロジェクトを再構築するための CLI。)
スクリプトを .ultracodex/workflows/<name>.js にドロップすると、実行可能になります。
名前 (ultracodex run <name>) で表示され、TUI ランチャー (ベア) で表示されます。
ウルトラコーデックス）。次に:
ウルトラコーデックスの同期スキル
ワークフローごとにクロード コード スキルを生成するため、クロードは
ワークフローを尋ねられることなく名前で保存 - の完全自動層
同じ統合です。
.ultracodex/config.toml (プロジェクト) または ~/.ultracodex/config.toml (グローバル):
[ ルート ] # 最初の一致が勝ちます: ラベル、次にフェーズ
"critique:*" = " claude " # 判断はクロードにあります
"verify:*" = " claude " # クロスベンダー判定が必要な場合は、ループ検証器も使用します
"*" = " codex " # 実行は Codex に移動します
[バックエンド。コーデックス]
サンドボックス = " ワークスペース書き込み "
デフォルト_モデル = " gpt-5.5 "
デフォルト_エフォート = " xhigh "
service_tier = " standard " # ~/.codex から高速モードを継承しない
model_map = { opus = " gpt-5.5 " 、 Sonnet = " gpt-5.4 " 、 haiku = " gpt-5.4-mini " }
[走る]
同時実行数 = 6 # デフォルト: min(16, cores-2)
ルーティングは設定内に存在し、スクリプト内に存在することはありません。それがスクリプトを維持するものです。
ランタイムとバックエンド間で移植可能です。
Ultracodex TUI ホーム: 保存されたワークフロー + 最近の実行
ウルトラ

[切り捨てられた]

## Original Extract

Run Claude Code workflow scripts, unmodified, on the OpenAI Codex CLI — fable plans, codex executes, fable verifies. Parallel agent fleets, builder–verifier loops, token budgets, full-screen TUI. - YuanpingSong/ultracodex

Claude Fable has been incredible, however the plan usage runs out too fast, especially if you use ultracode mode (Claude Code's workflow feature, where the model writes small JavaScript programs that orchestrate subagents) and let the agents go brrr. I want Fable to focus on high-value tasks such as planning or verification, and not burn tokens on mundane implementation work. So I built ultracodex, an engine that spawns codex agents to complete the same ultracode workflows and hand off seamlessly to and from Claude. It runs the same workflow scripts that Claude writes. I designed this tool for anyone owning both Claude and Codex subscriptions or anyone who enjoys Fable and wants to make the quota last. This tool is in early days but I am very excited about the potential. I think it’s also a great “loop engineering” starter pack for us because 1) Claude is great at writing loops and 2) ultracodex runs the loops with Codex agents - given their generous limits, you don’t have to worry about usage much. So “loop engineering” is no longer a mysterious thing but a tool you can throw at any task. I have some examples of “loops” in my repo as well. Project is completely open-source. Apache-2.0, npm i -g ultracodex, demo video in the README. Please check it out!

GitHub - YuanpingSong/ultracodex: Run Claude Code workflow scripts, unmodified, on the OpenAI Codex CLI — fable plans, codex executes, fable verifies. Parallel agent fleets, builder–verifier loops, token budgets, full-screen TUI. · GitHub
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
YuanpingSong
/
ultracodex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .github/ workflows .github/ workflows .ultracodex .ultracodex assets assets docs docs examples examples fixtures fixtures scripts scripts src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Run Claude Code workflow scripts, unmodified, on the OpenAI Codex CLI.
Claude Code's Workflow tool has a great orchestration format: plain-JS
scripts that fan agents out in parallel, pipeline work through stages,
enforce token budgets, and adversarially verify results. But running them
upstream — under Claude Code itself — spends Claude quota on
execution-grade work.
ultracodex is a compatible runtime for those same scripts — byte-identical,
no transpilation — that routes each agent() call to
Codex instead (or to any configured
backend, per label). The pattern it exists for:
fable plans, codex executes, fable verifies
(fable = Claude's frontier model; substitute your planner of choice)
Your most capable model authors and judges the work; cheaper/faster coding
agents do the bulk of it; one script, one journal, one budget.
Claude Code (right) is asked: "Write an essay on the meaning of life — actor–critic loop, 3 rounds. Run it with ultracodex." It authors the workflow, kicks it off, the TUI (left) watches Codex agents execute the loop, and the result JSON lands back in Claude. Sped up 1.3×. ( HD video )
Quota arbitrage. Offload execution to Codex (or another backend);
keep Claude for judgment. Zero Claude tokens spent on implementation.
Cross-vendor verification. Route critique:* to a different model
family than the one that did the work. Same-vendor self-review tends to
rubber-stamp; making the judge a different model family is one config
line here.
Durable, inspectable runs. No daemon. A run is a detached process and
a directory of plain files; an append-only journal is the single source of
truth for the TUI, the CLI, and machine consumers alike. Close your
terminal; the run doesn't care.
A script is an ES module: a pure-literal meta export, then a plain-JS
async body over eight injected globals. No imports, no TypeScript.
export const meta = {
name : 'review-files' ,
description : 'Fan out reviewers, verify findings, report' ,
phases : [ { title : 'Review' } , { title : 'Verify' } ] ,
}
const FILES = args ?. files ?? [ 'src/auth.ts' , 'src/api.ts' ]
phase ( 'Review' )
const findings = ( await parallel ( FILES . map ( f => ( ) =>
agent ( `Review ${ f } for bugs. Return via the schema.` , {
label : `review: ${ f } ` ,
schema : { type : 'object' , properties : { bugs : { type : 'array' , items : { type : 'string' } } } , required : [ 'bugs' ] } ,
} )
) ) ) . filter ( Boolean ) // failed agents are null, never throws
phase ( 'Verify' )
const verified = await pipeline ( // no barrier — each item flows independently
findings . flatMap ( f => f . bugs ) ,
( bug ) => agent ( `Try to refute: ${ bug } ` , { label : 'verify' } ) ,
( verdict , bug ) => ( { bug , verdict } ) , // verdict may be null (failed verifier) — check it
)
return { verified : verified . filter ( v => v && v . verdict ) }
global
what it does
agent(prompt, opts?)
run one agent; resolves final text, a schema-validated object, or null on failure (never rejects — except budget/caps, which throw)
parallel(thunks)
barrier over concurrent thunks; a thrown thunk becomes null
pipeline(items, ...stages)
per-item stage chains, no cross-item barrier; stages get (prev, item, index) ; the first stage's prev is the item itself. A stage that throws drops its item to null ; a stage that resolves null (e.g. a failed agent) flows onward — later stages must null-check prev
phase(title)
progress grouping for subsequent agents
log(msg)
narrator line in the TUI / --watch output
args
the run's --args input, verbatim
budget
{ total, spent(), remaining() } — output-token ceiling; exceeding it makes further agent() calls throw
workflow(name, args?)
run a saved workflow inline (one nesting level)
agent() opts: label (display + routing), phase , schema (JSON Schema),
model / effort (advisory tiers, mapped in config), isolation: 'worktree'
(fresh git wo
[truncated]
Scripts are plain JavaScript, so loops are native — and builder→verifier
loops are the pattern this tool is built to make easy. The runtime gives
loops their guardrails: a hard output-token budget (further agent()
calls throw once it's spent), a 1000-agent lifetime cap, and live
pause/skip/stop controls on every run.
let attempt = null , feedback = 'none — first attempt'
while ( budget . remaining ( ) > 20_000 ) { // rail 1: budget governor
attempt = await agent ( `Build X. Reviewer feedback: ${ feedback } ` , { label : 'build' } )
if ( attempt === null ) continue // rail 2: builders can fail
const verdict = await agent ( `Try to refute this: ${ attempt } ` , {
label : 'verify:attempt' , // route to another vendor in config
schema : { type : 'object' , properties : { pass : { type : 'boolean' } , issues : { type : 'array' , items : { type : 'string' } } } , required : [ 'pass' , 'issues' ] } ,
} )
if ( verdict ?. pass ) break // rail 2 again: judge can fail too
feedback = verdict ? verdict . issues . join ( '; ' ) : 'verifier unavailable'
}
return { attempt }
Three axes, one format: parallel() is breadth, pipeline() is flow,
loops are depth — iterate until verified good, with budget as the
governor and the run's live controls (pause/skip/stop) as the brakes. The
full pattern — max rounds, feedback threading, cross-vendor judging — is
examples/03-builder-verifier.js . Route
"verify:*" = "claude" in config and the builder's judge is a different
model family; the loop doesn't change.
The full normative definition — grammar, semantics, loop patterns,
conformance — is in docs/agent_script_spec.md .
The same file runs under Claude Code's Workflow tool and ultracodex;
ultracodex validate --strict checks a script stays in the portable subset.
Prerequisites: Node ≥ 20, pnpm , and the
Codex CLI installed and authenticated
( codex login ). Developed and tested against codex-cli 0.142.4
( ultracodex doctor prints your version). Optional: Claude Code for
claude-routed agents.
npm install -g ultracodex # or: pnpm add -g ultracodex
ultracodex doctor # checks node, codex, auth, config
From source instead:
git clone https://github.com/YuanpingSong/ultracodex && cd ultracodex
pnpm install && pnpm build
pnpm link --global # → `ultracodex` on PATH (or keep using node dist/cli.js)
The examples ship with the package — run them first (from the checkout, or
from the installed package via $(npm root -g)/ultracodex/ ):
ultracodex run examples/01-hello.js --watch # one agent, streamed events
ultracodex run examples/02-fanout-critique.js # fan-out + schemas, opens the TUI
ultracodex run examples/03-builder-verifier.js --watch --budget 200k # the loop
ultracodex ls # every run, pid-liveness checked
ultracodex show < runId > --json # machine-readable result
For real work, run from your project's root — agents work in your cwd.
run launches a detached runner — quitting the TUI never kills a run.
--json blocks and prints the result (this is what a driving LLM calls);
--watch streams events line-by-line; --detach just prints the runId and
exits; --budget takes output tokens as an integer with optional k / m
suffixes ( --budget 500k ).
Driving ultracodex from Claude Code
That's what the demo above shows, and it needs one command of setup:
ultracodex sync-skills # installs the `ultracodex` skill for Claude Code
The skill teaches Claude the whole contract — author the script exactly as
for the Workflow tool, execute with ultracodex run <file> --json [--budget 500k] , relay the result verbatim. After that, the prompt is just
the task:
Write an essay on the meaning of life — actor–critic loop, 3 rounds.
Run it with ultracodex.
Without the skill, prompt it explicitly — two variants depending on what's
in the session's context:
Option A — schema-native (session has the Workflow tool, e.g. ultracode
mode; Claude already knows the format as that tool's schema):
Author the workflow exactly as you would for the Workflow tool — same
script, byte for byte — but instead of invoking the tool, save it to a
file and run ultracodex run <file> --json --budget 300k . Relay the
result JSON verbatim; if the run fails, report the failure.
Option B — teach from the spec (plain sessions, headless claude -p ,
other agents entirely): as Option A, but point at the format docs —
"using the Agent Script format in docs/agent_script_spec.md" — and state
the task. (That's how this repo validated itself: a fresh headless Claude
authored staged build workflows from the spec and drove them through this
CLI to rebuild the project with Codex agents.)
Drop scripts in .ultracodex/workflows/<name>.js and they become runnable
by name ( ultracodex run <name> ) and visible in the TUI launcher (bare
ultracodex ). Then:
ultracodex sync-skills
generates a Claude Code skill per workflow, so Claude can trigger your
saved workflows by name without being asked — the fully-automatic tier of
the same integration.
.ultracodex/config.toml (project) or ~/.ultracodex/config.toml (global):
[ route ] # first match wins: label, then phase
"critique:*" = " claude " # judgment goes to Claude
"verify:*" = " claude " # loop verifiers too, if you want cross-vendor judging
"*" = " codex " # execution goes to Codex
[ backends . codex ]
sandbox = " workspace-write "
default_model = " gpt-5.5 "
default_effort = " xhigh "
service_tier = " standard " # never inherit fast mode from ~/.codex
model_map = { opus = " gpt-5.5 " , sonnet = " gpt-5.4 " , haiku = " gpt-5.4-mini " }
[ run ]
concurrency = 6 # default: min(16, cores-2)
Routing lives in config, never in scripts — that's what keeps scripts
portable across runtimes and backends.
ultracodex TUI home: saved workflows + recent runs
ultr

[truncated]
