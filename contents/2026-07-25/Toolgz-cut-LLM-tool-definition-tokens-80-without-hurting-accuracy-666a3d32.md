---
source: "https://github.com/dperussina/toolgz"
hn_url: "https://news.ycombinator.com/item?id=49051573"
title: "Toolgz – cut LLM tool-definition tokens ~80% without hurting accuracy"
article_title: "GitHub - dperussina/toolgz: Compress your tool calls and save tokens over time. · GitHub"
author: "dperussina"
captured_at: "2026-07-25T21:44:29Z"
capture_tool: "hn-digest"
hn_id: 49051573
score: 2
comments: 0
posted_at: "2026-07-25T21:05:43Z"
tags:
  - hacker-news
  - translated
---

# Toolgz – cut LLM tool-definition tokens ~80% without hurting accuracy

- HN: [49051573](https://news.ycombinator.com/item?id=49051573)
- Source: [github.com](https://github.com/dperussina/toolgz)
- Score: 2
- Comments: 0
- Posted: 2026-07-25T21:05:43Z

## Translation

タイトル: Toolgz – 精度を損なうことなく LLM ツール定義トークンを最大 80% 削減
記事のタイトル: GitHub - dperussina/toolgz: ツール呼び出しを圧縮し、時間をかけてトークンを節約します。 · GitHub
説明: ツール呼び出しを圧縮し、時間をかけてトークンを保存します。 - GitHub - dperussina/toolgz: ツール呼び出しを圧縮し、時間をかけてトークンを節約します。

記事本文:
GitHub - dperussina/toolgz: ツール呼び出しを圧縮し、時間をかけてトークンを節約します。 · GitHub
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
ペルシーナ
/
ツールグズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github/ workflows .github/ workflows ベンチ ベンチ 脳 脳 ドキュメント ドキュメント 仕様/ 001-core-compression スペック/ 001-core-compression src src テスト テスト .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.mdライセンス ライセンス通知 通知 README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントは、ユーザーが単語を入力する前に、ツール定義に 30 ～ 50,000 のコンテキスト トークンを費やします。 toolsgz はその最大 80% を取り戻します。
420 実行のクロスプロバイダー スイープ ·
4つのフロンティアモデル・
実行時の依存関係がゼロ ·
前後に生成される
npmインストールツールgz
問題
いくつかの MCP サーバーを接続します。それぞれ 20 ～ 50 個のツールが同梱されます。すべてのツールは JSON スキーマであり、
パラメータごとの散文。このブロックは、すべての単一リクエストの先頭でレンダリングされます。
現実的なツール定義は最大 420 個のトークンであり、そのうち約 400 個がモデルの散文です
正しく選択するためには必要ありません。ツール 50 個は 20,000 トークンです。 100は40kです。
迅速なキャッシュにより、トークンが安価になります。占有スペースが少なくなるわけではありません。
部屋を取り戻すというのがこれです。
import { compress , forAnthropic } from "toolgz" ;
const c = 圧縮 (myTools) ; // 既存の MCP/SDK ツール配列
const {ツール、システム} = forAnthropic(c); // 代わりにこれらを送信します
次に、ディスパッチする前にモデルのコールバックを変換します。
const r = c 。解決 (ブロック . 名前 , ブロック . 入力 ) ;
if ( r . kind === "call" ) await myDispatch ( r . name , r . args ) ; // 実際の名前、実際の引数
myDispatch は、以前に取得したものとまったく同じものを取得します。下流側では何も変わりません。
4 つのフロンティア モデル、7 つの戦略、5 つのツール選択

タスク、3 回 —
現在のスイープでは 420 回実行され、全ラウンドで 1,200 回以上実行されます。すべての実行ごとの生のレコードは、
ベンチ/結果/でコミットされました。任意の数値を再計算します
npx tsx ベンチ/analyze-multi.ts 。
プロバイダー
モデル
ツールブロック
プロンプトトークン
コスト
レイテンシー
タスク
人間的
クロード作品5
9,242 → 1,284
30,817 → 4,628 (−85%)
−78%
15.0秒→12.1秒
15/15
xAI
グロク-4.5
6,421 → 775
17,522 → 2,663 (−85%)
−70%
6.1秒→4.6秒
15/15
Google
gemini-3.1-pro-プレビュー
5,264 → 732
10,948 → 2,302 (−79%)
−62%
5.6秒→5.5秒
15/15
OpenAI
gpt-5.6-ソル
2,752 → 573
7,694 → 2,196 (−71%)
−7%
6.8秒→5.6秒
15/15
推論は高い努力をすれば 4 つすべてで有効になるため、これは同様のフロンティアです
比較。 60/60 のタスクが完了、幻覚ツール名なし、不正な形式なし
引数 - そして、どのプロバイダーでも非圧縮よりも高速です。
モデルを悪くするものではありません
それは反証すべきことであり、私たちはそうしようと懸命に努力しました。タスクスイートは以下から構築されます
意図的に混同しやすいツール クラスタ — search_issues と list_issues 、 comment-vs-update 、
承認とマージ、同じ 3 つの製品を並べて表示 - 正しい選択が有効になります
圧縮によって削除されるツール名。
モデルは選択する能力を失わず、想起の問題を検索の問題に変換します。
問題が発生し、必要なものを調べます。赤いセルがあるため、デフォルトのマップ スタイルが存在します。
ベアツール名は grok-4.5 で決定的に失敗し、1 つのシナリオで 3 回の試行中 3 回失敗しました。
ツール呼び出しはゼロで、エラーも発生しませんでした。必要な引数に名前を付けることで問題が修正されました。
コストと、正しく理解するのに 2 ラウンドかかった理由
最初のプロバイダー間スイープでは、コンテキストが低下しているにもかかわらず、OpenAI でコストが 15% 上昇していることが判明しました
69%。ディスパッチャーは余分なターンを費やしており、推論モデルでは、各ターンでコストが支払われます。
これの新鮮なラウンド

キング。
そこで、推測ではなく拒否されている通話をキャプチャし、3 つのバグを発見しました。
このライブラリでは、モデルはクエリを q という名前のパラメーターに渡します (18 個中 14 個が拒否)。
マップ コードをツール名として呼び出す場合もあり、代わりに引数をフラットに渡す場合もあります。
ネストされた。 3 つすべてを修正すると、OpenAI は +15% から -7% になり、不正な引数が減少しました。
すべてのプロバイダーでゼロ。
OpenAI の -7% は依然として最小の節約額であり、正直に言ってそうです。推論出力が OpenAI の大半を占めています。
したがって、プロンプトが小さいほど合計が少なくなります。コンテキスト ウィンドウの占有率は維持されます。
プライマリクレーム — 推論設定に応じた金額でコストが計算されます。
図書館に聞いてください。 1 または 3 を返しますが、2 は返しません。次のように説明されています。
"toolgz" から { recommendLevel } をインポートします。
const { レベル , 理由 } = recommendLevel (myTools) ;
レベル
送信します
本名
プロバイダースキーマの適用
いつ使用しますか
1
それぞれ 1 つのネイティブ ツール、署名行の説明
はい
はい
デフォルト。小さい、または広くてまばらなツールセット。測定されたマイナス面はゼロ。
2
名前空間ごとに 1 つの複合ツール
はい
いいえ
回線上には読みやすい op 名が必要です。それ以外の場合はスキップします。
3
1 つのディスパッチャー + 1 つのルックアップ ツール
コード
いいえ
大きくて深いツールセット。上記の 80% という数字。
(レベル 0 は、独自のアプリ内での A/B テスト用のパススルーです。)
レベル 1 は無料です - 測定値: トークンが少なく、不正な引数がゼロ、余分なターンがゼロ、
遅延も悪くありません。レベル 2 は、より多くの生産を行うなど、あらゆる軸でレベル 3 によって支配されています。
不正な引数。それは踏み台ではありません。
レベル 3 は実際にはどのようなものですか
開始するツールの数に関係なく、ネットワーク上の 2 つのツールと、システム プロンプト内のマップ
キャッシュ ブレークポイントの背後:
<ツールマップ>
a0 github_create_issue 所有者、リポジトリ、タイトル
a1 github_search_issues q
b0lack_post_message チャネル、テキスト
</ツールマップ>
モデルは t(f="a0", a

={…}) 、および q(c="a0") は、次の場合にコードを完全な署名に展開します。
オプションのパラメータが必要です。
これらの検索を完全に削除したい場合は、署名全体をマップに配置します。
compress ( myTools , { level : 3 , mapStyle : "signature" } ) ;
// a0 github_create_issue(owner,repo,title,body?,labels?)
測定結果: ルックアップがゼロになり、OpenAI (4.0 秒、
−17%のコスト）。これはわずかに大きく、xAI ではデフォルトよりも悪かったため、
デフォルトではなくオプションです。
取引: レベル 2 ～ 3 では、モデルは汎用引数オブジェクトを埋めるため、プロバイダーの
サンプラーはスキーマを強制しなくなりました。 toolgz は元のスキーマに対して検証し、
代わりに、モデルが読み取り可能なエラーを返します。これが、validate がデフォルトでオンになっている理由です。オンのままにしておきます。
上記のすべてのアーティファクトは、ライブラリを実行することによって生成されます。「」を参照してください。
完全なツール配列とシステム プロンプトについては docs/BEFORE-AFTER.md、
前後、すべてのレベルで、実際のトークン数とライブエンコード→デコードラウンドを使用
旅行。テストではファイルがコードと一致することが確認されるため、ドリフトすることはありません。
完全ガイド
インストール→エージェントの動作ループ。プロンプト キャッシュ、MCP アグリゲーション、トラブルシューティングの 4 つすべてに対するプロバイダーごとのセットアップ。ここから始めましょう。
前 / 後
生成されますが、図示されていません。どちらのアーティファクトも、toolgz があらゆるレベルで変更します。
完全な結果
あらゆる数字、方法論、そしてそれが確立していないもの。
発表草案
HN および LinkedIn にすぐに投稿できる書き込みと、行ってはいけない主張。
解放する
有効期間の長いトークンを使用せずに、OIDC 経由で npm にパブリッシュします。
プロバイダー
import { forAnthropic , forOpenAI , forOpenAIResponses , forGemini } from "toolgz" ;
純粋な関数。彼らはあなたが渡したものを決して変更しません。
xAI は OpenAI と互換性があります。baseURL: "https://api.x.ai/v1" で forOpenAI を使用します。
コスト削減の規模

トークンの保存サイズではありません。測定値 62 ～ 78%
3 つのプロバイダーでは安くなりますが、推論出力が大半を占める OpenAI ではわずか 7% です。
請求書。主張はコンテキストウィンドウの占有です。コストは可変量で続きます。
ツールブロックのサイズに関しては、Anthropic のネイティブ ツール検索に勝るものはありません。で構成されます
それは、同等のものが存在しない場所でも機能し、フロンティア層の下でより信頼性が高くなります。
defer_loading は、Haiku 4.5 では 6/30 タスクのみをサイレントに完了しました。
ツールを検出するかどうかを選択します。ディスパッチャは、ディスカバリをエントリ ポイントにします。
レベル 3 の非フロンティア モデルでは測定されていません。 Haiku 4.5 では、引数
エラーが急激に増加しました (30 実行中 17 実行) - すべてキャッチされて再試行され、タスクは失われませんでしたが、それは
既知のエッジ。
それは十の道具を使った魔法ではありません。ツールが 15 個未満では、再利用できるものはほとんどありません。
recommendLevel() がそのことを教えてくれます。
compress() は参照透過的です。同じツールが入力され、バイト同一のペイロードが出力されます。ツール
ソートされ、反復順序がそのままになることはありません。
これは正確性のプロパティであり、整理整頓ではありません。プロンプトのキャッシュはプレフィックスの一致であるため、
reordered ツールは、プロンプト全体をサイレントに再請求します。バイト安定性を主張するテストがあります。
そして削除されません。
npm test # 131 テスト、オフライン、無料
npm run build # tsc → dist/ with .d.ts
npx tsx bench/harness/run-multi.ts --provider=all --reps=3 --variants # コストがかかります
npx tsxベンチ/analyze-multi.ts
npx tsx docs/generate-examples.ts
方法論とリポジトリの規則: AGENTS.md 。
原則の仕様は docs/CONSTITUTION.md に対してチェックされます。
Apache-2.0 — 「ライセンス」と「通知」を参照してください。
MIT ではなく Apache-2.0 を意図的に使用しています。これには明示的な特許付与があり、
特許報復条項。これは、特許を導入する図書館にとって重要です。
単なるグルーコードではなくテクニック、そしてそれがライセンスです。

OST企業
依存関係にあることを好みます。
ツール呼び出しを圧縮し、時間をかけてトークンを節約します。
Readme Apache-2.0 ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Compress your tool calls and save tokens over time. - GitHub - dperussina/toolgz: Compress your tool calls and save tokens over time.

GitHub - dperussina/toolgz: Compress your tool calls and save tokens over time. · GitHub
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
dperussina
/
toolgz
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github/ workflows .github/ workflows bench bench brain brain docs docs specs/ 001-core-compression specs/ 001-core-compression src src tests tests .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE NOTICE NOTICE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Your agent spends 30–50k tokens of context on tool definitions before the user types a word. toolgz gets ~80% of it back.
420-run cross-provider sweep ·
4 frontier models ·
zero runtime dependencies ·
generated before/after
npm install toolgz
The problem
You connect a few MCP servers. Each ships 20–50 tools. Every tool is a JSON Schema with a
sentence of prose per parameter. That block renders at the front of every single request .
A realistic tool definition is ~420 tokens, and roughly 400 of them are prose the model
doesn't need in order to pick correctly. Fifty tools is 20k tokens. A hundred is 40k.
Prompt caching makes those tokens cheap . It does not make them take up less room .
Reclaiming the room is what this does.
import { compress , forAnthropic } from "toolgz" ;
const c = compress ( myTools ) ; // your existing MCP/SDK tool array
const { tools , system } = forAnthropic ( c ) ; // send these instead
Then translate the model's call back before you dispatch:
const r = c . resolve ( block . name , block . input ) ;
if ( r . kind === "call" ) await myDispatch ( r . name , r . args ) ; // real name, real args
myDispatch gets exactly what it got before. Nothing downstream changes.
Four frontier models, seven strategies, five tool-selection tasks, 3 reps —
420 runs on the current sweep, 1,200+ across all rounds. Every raw per-run record is
committed in bench/results/ ; recompute any figure with
npx tsx bench/analyze-multi.ts .
Provider
Model
Tool block
Prompt tokens
Cost
Latency
Tasks
Anthropic
claude-opus-5
9,242 → 1,284
30,817 → 4,628 (−85%)
−78%
15.0s → 12.1s
15/15
xAI
grok-4.5
6,421 → 775
17,522 → 2,663 (−85%)
−70%
6.1s → 4.6s
15/15
Google
gemini-3.1-pro-preview
5,264 → 732
10,948 → 2,302 (−79%)
−62%
5.6s → 5.5s
15/15
OpenAI
gpt-5.6-sol
2,752 → 573
7,694 → 2,196 (−71%)
−7%
6.8s → 5.6s
15/15
Reasoning is enabled on all four at high effort, so this is a like-for-like frontier
comparison. 60/60 tasks completed, zero hallucinated tool names, zero malformed
arguments — and it is faster than uncompressed on every provider.
It does not make the model worse
That was the thing to disprove, and we tried hard to. The task suite is built from
deliberately confusable tool clusters — search_issues vs list_issues , comment-vs-update,
approve-vs-merge, the same three products side by side — where the correct choice turns on
the tool name that compression takes away.
The model doesn't lose the ability to choose — it converts a recall problem into a retrieval
problem and looks up what it needs. The default map style exists because of the red cell:
bare tool names failed on grok-4.5 deterministically , 3 of 3 attempts on one scenario,
answering with zero tool calls and no error raised. Naming the required arguments fixed it.
Cost, and why it took two rounds to get right
The first cross-provider sweep found cost going up 15% on OpenAI even while context fell
69%. The dispatcher was spending extra turns, and on a reasoning model every turn pays for a
fresh round of thinking.
So we captured the calls that were being rejected instead of guessing, and found three bugs
in this library : models pass query to a parameter named q (14 of 18 rejections), they
sometimes call the map code as the tool name, and they sometimes pass arguments flat instead
of nested. Fixing all three took OpenAI from +15% to −7% and drove malformed arguments to
zero on every provider .
OpenAI's −7% is still the smallest saving, and honestly so: reasoning output dominates its
bill, so a smaller prompt moves the total less. Context-window occupancy remains the
primary claim — cost follows from it, by an amount that depends on your reasoning settings.
Ask the library. It returns 1 or 3, never 2, and explains itself:
import { recommendLevel } from "toolgz" ;
const { level , reason } = recommendLevel ( myTools ) ;
Level
Sends
Real names
Provider schema enforcement
Use when
1
one native tool each, signature-line descriptions
yes
yes
default. Small or wide-and-sparse tool sets. Zero measured downside.
2
one compound tool per namespace
yes
no
you need readable op names on the wire. Otherwise skip.
3
one dispatcher + one lookup tool
codes
no
large, deep tool sets. The 80% number above.
(Level 0 is a passthrough, for A/B testing inside your own app.)
Level 1 is free — measured: fewer tokens, zero malformed arguments, zero extra turns,
latency no worse. Level 2 is dominated by level 3 on every axis, including producing more
malformed arguments; it is not a stepping stone.
What level 3 actually looks like
Two tools on the wire regardless of how many you start with, and a map in the system prompt
behind a cache breakpoint:
<toolmap>
a0 github_create_issue owner,repo,title
a1 github_search_issues q
b0 slack_post_message channel,text
</toolmap>
The model calls t(f="a0", a={…}) , and q(c="a0") expands a code to its full signature when
it needs the optional parameters.
If you want to remove those lookups entirely, put the whole signature in the map:
compress ( myTools , { level : 3 , mapStyle : "signature" } ) ;
// a0 github_create_issue(owner,repo,title,body?,labels?)
Measured: lookups drop to zero and it was the fastest and cheapest arm on OpenAI (4.0s,
−17% cost). It is slightly larger, and on xAI it was worse than the default, so it is an
option rather than the default.
The trade: at levels 2–3 the model fills a generic argument object, so the provider's
sampler no longer enforces your schema. toolgz validates against your original schema and
returns a model-readable error instead. That is why validate defaults to on — leave it on.
Every artifact above is generated by running the library — see
docs/BEFORE-AFTER.md for the full tools array and system prompt,
before and after, at every level, with real token counts and a live encode → decode round
trip. A test asserts that file matches the code, so it cannot drift.
Complete guide
Install → working agent loop. Per-provider setup for all four, prompt caching, MCP aggregation, troubleshooting. Start here.
Before / after
Generated, not illustrated. Both artifacts toolgz modifies, at every level.
Full results
Every number, the methodology, and what it does not establish.
Announcement drafts
Ready-to-post write-ups for HN and LinkedIn, plus claims not to make.
Releasing
Publishing to npm over OIDC, with no long-lived token.
Providers
import { forAnthropic , forOpenAI , forOpenAIResponses , forGemini } from "toolgz" ;
Pure functions; they never mutate what you pass them.
xAI is OpenAI-compatible — use forOpenAI with baseURL: "https://api.x.ai/v1" .
The size of the cost saving is not the size of the token saving. Measured 62–78%
cheaper on three providers but only 7% on OpenAI, where reasoning output dominates the
bill. The claim is context-window occupancy; cost follows, by a variable amount.
It does not beat Anthropic's native tool search on tool-block size. It composes with
it, works where there is no equivalent, and is more reliable below the frontier tier —
defer_loading completed only 6/30 tasks on Haiku 4.5, silently, because it lets the model
choose whether to discover tools. A dispatcher makes discovery the entry point.
It has not been measured on a non-frontier model at level 3. On Haiku 4.5, argument
errors rose sharply (17 of 30 runs) — all caught and retried, no task lost, but that is the
known edge.
It is not magic on ten tools. Under ~15 tools there is little to reclaim;
recommendLevel() will tell you so.
compress() is referentially transparent: same tools in, byte-identical payload out. Tools
are sorted, never left in iteration order.
This is a correctness property, not tidiness — prompt caching is a prefix match, so one
reordered tool silently re-bills your whole prompt. There is a test asserting byte-stability,
and it does not get deleted.
npm test # 131 tests, offline, no cost
npm run build # tsc → dist/ with .d.ts
npx tsx bench/harness/run-multi.ts --provider=all --reps=3 --variants # costs money
npx tsx bench/analyze-multi.ts
npx tsx docs/generate-examples.ts
Methodology and repo conventions: AGENTS.md .
Principles specs are checked against: docs/CONSTITUTION.md .
Apache-2.0 — see LICENSE and NOTICE .
Apache-2.0 rather than MIT deliberately: it carries an express patent grant and
a patent-retaliation clause, which matters for a library implementing a
technique rather than just glue code, and it is the license most enterprises
prefer in a dependency.
Compress your tool calls and save tokens over time.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
