---
source: "https://github.com/fayzan123/rungraph"
hn_url: "https://news.ycombinator.com/item?id=49321388"
title: "Show HN: Rungraph – See your AI coding-agent runs as a graph"
article_title: "GitHub - fayzan123/rungraph: See your agent runs as a graph — zero-setup, agent-first visualizer for Claude Code sessions and workflows · GitHub"
author: "FayzanMalik"
captured_at: "2026-08-16T17:11:56Z"
capture_tool: "hn-digest"
hn_id: 49321388
score: 1
comments: 0
posted_at: "2026-08-16T16:18:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rungraph – See your AI coding-agent runs as a graph

- HN: [49321388](https://news.ycombinator.com/item?id=49321388)
- Source: [github.com](https://github.com/fayzan123/rungraph)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T16:18:51Z

## Translation

タイトル: HN の表示: Rungraph – AI コーディング エージェントの実行をグラフとして表示します。
記事のタイトル: GitHub - fayzan123/rungraph: エージェントの実行をグラフとして表示 — クロード コードのセッションとワークフローのためのゼロセットアップ、エージェントファーストのビジュアライザー · GitHub
説明: エージェントの実行をグラフとして表示 — クロード コードのセッションとワークフローのためのゼロセットアップ、エージェントファーストのビジュアライザー - fayzan123/rungraph

記事本文:
GitHub - fayzan123/rungraph: エージェントの実行をグラフとして表示 — クロード コードのセッションとワークフローのためのゼロセットアップ、エージェントファーストのビジュアライザー · GitHub
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
フェイザン123
/
ラングラフ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
38 コミット 38 コミット .github/ workflows .github/ workflows bin bin docs docs フロントエンド フロントエンド スクリプト スクリプト src src テスト テスト .gitignore .gitignore CLAUDE.md

CLAUDE.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SCHEMA.md SCHEMA.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントの実行をグラフで確認します。
これは、この機能を構築したライブ セッション、つまりストリップを観察しているラングラフです
何が問題だったのかが表示され、ワンクリックでそれが意味するノードが点灯します。
コーディング エージェントは、その実行内容をすべてすでに書き留めています。ラングラフはそれらを変える
インタラクティブな有向エージェント グラフへのトランスクリプト — オーケストレーター、
サブエージェントとノードとしてのツール。関係をエッジとして生成/返します。の
コース変更の瞬間（拒否、回答、再試行）がパス上にマークされます。それは機能します
これまでに実行したすべてのセッションに遡って、フックもラッパーもなし
セットアップ、テレメトリなし。
npxラングラフ
これがクイックスタート全体です。 ~/.claude/projects をスキャンし、ローカルのプロジェクトを開始します。
サーバーにアクセスし、ブラウザを開きます。実行中の実行を含む実行を選択します
現時点では、グラフはエージェントの動作に応じてライブに成長します (ファイル監視のみ)。
グラフは、話しかけることもできます。MCP サーバーをコンピューターに接続します。
コーディング エージェントに実行について尋ねます。「なぜ token.js の編集が維持されたのか
失敗してる？」 — 回答が端末に届き、エージェントが強調表示します
応答時に開いたグラフ上に記述されている正確なノード。参照
ランニングについてエージェントに尋ねてください。
ここは新しいですか？ docs/GUIDE.md で全体を説明します —
グラフを読み、各信号が何を意味するのか、それを自分のエージェントに接続するのか、そして何を意味するのか
何かが壊れているように見えるときに行うこと。
エージェント セッションは少し前に会話ではなくなりました。それらは実行です:
オーケストレーターがサブエージェントを生成する、ワークフローがファンアウトするレビューアー、ツール
失敗しては再試行し、人間は時々「ノー」と言います。ラングラフはそれを描きます
構造的には 4,000 行のトランスクリプトになる

実際に読めるものは何かあります:
時間が流れていきます。プロンプトはバックボーンです。並列エージェントがファンアウトする
隣り合ったレーンに移動し、結果を受け取ったターンに戻ります。
ツールノードは、どのツールを実行したかだけでなく、何を実行したかを示します: Bash · npm test ×12 、
編集 · Canvas.jsx 、Grep · waitForURL 。同じ内容の連続呼び出し
ツールは 1 つのノードに折りたたまれるため、テストと修正のループが毛玉にならないようになります。
ノードをクリックすると、ターンのプロンプトと応答の全文が表示されます。ごとに
呼び出しの入力、出力、エラー、ツールのタイミング。完全な
サブエージェントのトランスクリプト。ツールノードには、エージェントの理由も表示されます。
電話をかける直前の自分のナレーション (「これからテストを再実行します)
チェックしてください…」）。
人間の介入は第一級のノードです。拒否された許可、
質問への回答、ターン途中の割り込み - これらはランの瞬間です
方向を変える、そしてそれに続くエッジが理由を運ぶ
(権限拒否後、失敗後、Bash エラー後に再試行)。
ワークフローの実行 (マルチエージェント オーケストレーション) は、単一のノードとして表示されます。
独自のグラフ、フェーズ ボックスなど、すべてにリンクされた再試行をドリルダウンできます。
彼らが置き換えた試み。
トークン、期間、およびモデルはノードに注釈を付けます。の全実行合計
ヘッダー。
すべてを等しい重みで描画するグラフは、何もポイントを持たない:
2 秒のファイル読み取りと 40 分間の再試行スパイラルは同じように見えます。それで
ラングラフには意見があります。ランニングからシグナルを導き出し、それらを
キャンバスの上でストリップします。クリーンな実行では、ストリップの高さはゼロになります。
信頼できないマーカーはマーカーがないより悪いからです。
信号をクリックすると、グラフがフォーカスされます。それらのノードが点灯し、他のすべてのノードが点灯します。
4 分の 1 に薄暗くなります — 薄暗くなり、決して隠されることはありませんので、すでに記憶した形状がそのまま表示されます
そのままです。 Esc または空のキャンバスをクリックすると、キャンバスがクリアされます。
同じフォーカスメカ

nism は、ノードを指す他のすべてをサポートします。
検索 ( / ) — ノード ラベルおよび各ノードのファイル上のプレーンな部分文字列
触れた。モデルもネットワークもサブプロセスもありません。ブラウザでフィルタリングします。
ファイル — ツールとエージェントのノードは、作業を含む、それらが接触したパスを運びます。
サブエージェント内で行われ、実際の編集はここで行われます。の
インスペクターは、実行で触れられたすべてのファイルをカウントとともにリストします。 1 つをクリックして表示します
どのステップがそれに影響したかを正確に示します。
ライブ エスカレーション — ライブテール更新ごとにシグナルが再導出されます。やってみろよ
エージェントの作業中に別のことを行う。ストリップは何かがあったときだけ大音量になります
新しいことは実際には間違っていました。
ダッシュボードはあなたのためのものです。 MCP サーバーはエージェント用です。それらはの両端です
2 つの製品ではなく、1 つのループです。
npx rungraph mcp --install # 1 回、その後クロード コードを再起動します
npx rungraph mcp --check # 動作していますか?修正すべき内容を正確に出力します
次に、クロード コードで、トランスクリプトで実際にどのような質問ができるかを尋ねます。
答え:
前回の実行で失敗した編集はどれですか?壊れたままになっているものはありますか?
サブエージェントを含めて、どのステップが src/auth.js に影響しましたか?
「監査認証モジュール」エージェントは何を見つけましたか?
その赤いノードの背後にある実際のエラーは何でしたか?
実際にテストを実行したのか、それともテストを実行したとだけ言ったのか?
クロードは find_nodes / get_graph / get_detail を呼び出して、
ターミナル — モデル、セッション、完全に検査可能。それからそれは呼び出します
focus_nodes 、開いているグラフの正確なノードが点灯します。
答えは、適切な実行に切り替えるか、ブラウザのタブを開くかです。
そうする必要があり、同じハイライトを復元するディープリンクを返します。
貼り付ける人は誰でも。
質問を考案する必要もありません: インスペクターの下部
コピー ボタンを使用すると、見ている実行からそれらの情報が自動的に書き込まれます。
何も固定されず、プロンプトもプロキシもされません

: rungraph はグラフに寄与しますが、
会話。読み取り専用ツールは、サーバーがまったく実行されていない状態でも機能します。
複数のダッシュボードがライブになっています - あなたのものに加えて、誰かがあなたに送ったバンドル
(下) — MCP はそれらを集約します: list_runs はすべてのサーバーの実行をマージします。
どこから来たのかというタグが付けられ、他のすべてのツールは実行 ID によってルーティングされます。
実際にその走行を示すダッシュボード。
実行は、あなたの条件に応じて、ファイルとしてマシンから離れることができます。ビラルの代理人が行ったと言う
横向きに見ると手伝うこともできますし、同僚に機能の場所を見せたい場合も
実際に建てられました。
ビラルの輸出。ダッシュボードから、または実行ペインで共有…を確認します。
オフラン、出発予定の内容を確認するか、エージェントに次のように尋ねます。
rungraph エクスポート --last 2 --as Bilal
# rungraph: インベントリのエクスポート (完全なコンテンツ):
# 2 実行 · 143 ノード · 12 個のプロンプトが含まれる
タッチされたファイルの数: 24
# rungraph: acme-2026-08-15.rungraph を書き込みました (412,882 バイト)
インベントリは毎回印刷されます。人々は、ある場所にどれだけの人が住んでいるのかを理解していません。
トランスクリプトなので、ツールは終了する前にそれを表示します。そして、ブロックをエクスポートする場合は、
信頼性の高いシークレット (AWS キー、GitHub/Slack/API トークン、秘密キー) を見つけます。
ブロック - アンカーされたパターン、誤検知がほぼゼロになるように調整されています）、リスト
それぞれが正確にどこにいるのか。 --redact-secrets で解決します (プレースホルダー、
それ以外はすべてそのまま)、 --struction-only (グラフの形状、ツール名、ファイル)
およびタイミング - プロンプトなし、出力なし)、またはフィクスチャの場合は --allow-secrets
チェックしたキー。
ファイルは転送です。すでにあるものの上に .rungraph を送信します
信頼 — Slack、AirDrop、リポジトリ。 rungraph 自体はネットワークに触れることはありません。
npx rungraph チームワーク.rungraph を開く
これにより、独自の一時的なダッシュボードでバンドルが提供されます。何もコピーされません。
どこでも。プロセスを閉じるとプロセスは終了します。再度開くためにファイルを保存してください

何もない
時間。すべてのランにはその由来が込められています (「Bilal · Team-work.rungraph が共有」)。
ループ全体がそれに対して動作します。信号はラングラフと独自のものから派生します。
エージェントはビラルの実行を指摘される可能性があります - 「バンドルで何が間違っていたのかビラル
私を送ってくれた？」 — あなた自身の隣にあります。
バンドルにはベンダー中立の IR が含まれるため、Codex を実行するとエクスポートして開きます。
クロード コードのものと同じで、バンドルを開くときにアダプターは必要ありません。
全部。 sharedBy は表示文字列であり、ID ではありません - バンドルを信頼する方法
あなたはそれが到着したチャネルを信頼します。
表示されているものにリンクします。ヘッダー内のコピーリンクは現在のビューをキャプチャします —
実行、選択したノード、フォーカス - URL として。 focus_nodes は同じ種類のを返します
リンクをクリックすると、エージェントが PR や問題のために貼り付けられるものを渡すことができます。リンク
ロード時にクエリを再実行します (検索リンクは再検索、シグナルリンクは再検索します)
再派生)、間違ったダッシュボードに到達したリンクはワンクリックで提供されます。
ランがあるものにジャンプします。
ナビゲーションは Figma スタイルで、実際の実行時に縦長で細いグラフ向けに構築されています。
生産する：
ミニマップ (右下) には、ラン全体がドラッグ可能なストリップとして表示されます。
ビューポート — エラーは赤いビーコンとして点灯します。 1 つをクリックすると、その項目に直接ジャンプします
失敗。可読ズームで実行が開きます: 最初のプロンプトで終了した実行、ライブ
最新のアクティビティで実行され、フォローモードでビューを新しいノードとしてスライドさせます。
流れ込みます。
UI で実行できることはすべて、コーディング エージェントが CLI 経由で実行できます。ブラウザも必要ありません。
プロンプト、標準出力の JSON、標準エラー出力のログ、終了コード 0 ok / 1 エラー / 2 no
ランが見つかりました。このセクションをプロンプトに貼り付けると、エージェントがセルフサービスで対応できます。
npx ラングラフ リスト --json
# {"runs":[{"runId":"claude-code:…:5822df8b-…","kind":"session","title":"不安定な認証テストを修正する",
# "プロジェクト":"/home/you/dev/app","modifiedAt":"2026-08-11T16:31:06.055Z","active":true,…},…]}
n

px ラングラフ グラフ ' クロード コード:…:5822df8b-… ' --json
# 標準出力で実行される完全なグラフ IR:
# {"irVersion":1,"meta":{"runId":"…","kind":"session","title":"…","totals":{"tokens":184230,"toolCalls":57,"agents":4},…},
# "nodes":[{"id":"…","kind":"agent","label":"不安定なテストの調査","status":"completed",
# "ファイル":["/home/you/dev/app/src/auth/token.js"],"トークン":{…}},…],
# "edges":[{"kind":"spawn","from":"…","to":"…","label":"auth.spec.ts が不安定になる理由を調査する"},…],
# "グループ":[…],
# "signals":[{"kind":"retry-storm","severity":"high","nodeIds":["…"],"label":"編集呼び出しが 6 回失敗しました",
# "reason":"token.js の 3 つの連続するステップで編集が 6 回失敗しました、…"}]}
# → エージェントは自身の過去の実行を読み取ることができます: 何が生成されたか、何が失敗したか、人間がどこでノーと言ったのか。
npx rungraph find ' クロードコード:…:5822df8b-… ' token.js --json
# {"一致":4,"ノードID":[…],"ノード":[…]}
＃→まず狭い。大きなグラフには、1 つの質問に答えるために 20,000 以上のコンテキスト トークンが含まれます。
npx rungraphserve --no-open
# {"url":"http://127.0.0.1:4321"} (サーバーはフォアグラウンドに留まり、HTTP + SSE ライブ テール経由で同じデータ)
同じサーフェスを MCP ツールとして利用できます。「ランについてエージェントに問い合わせる」を参照してください。
上記、または rungraph mcp --install 。
IR はバージョン管理され、 SCHEMA.md に文書化されます。それは
ベンダー中立、2 つのアダプター: Claude Code (セッション

[切り捨てられた]

## Original Extract

See your agent runs as a graph — zero-setup, agent-first visualizer for Claude Code sessions and workflows - fayzan123/rungraph

GitHub - fayzan123/rungraph: See your agent runs as a graph — zero-setup, agent-first visualizer for Claude Code sessions and workflows · GitHub
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
fayzan123
/
rungraph
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits .github/ workflows .github/ workflows bin bin docs docs frontend frontend scripts scripts src src tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SCHEMA.md SCHEMA.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
See your agent runs as a graph.
That's rungraph watching the live session that built this feature — the strip
says what went wrong, and one click lights up the nodes it means.
Your coding agent already wrote down everything it did. rungraph turns those
transcripts into an interactive directed agentic graph — orchestrator,
subagents, and tools as nodes; spawn/return relationships as edges; the
course-change moments (denials, answers, retries) marked on the path. It works
retroactively on every session you've ever run : no hooks, no wrappers, no
setup, no telemetry.
npx rungraph
That's the whole quickstart. It scans ~/.claude/projects , starts a local
server, and opens your browser. Pick a run — including one that's running
right now : the graph grows live as the agent works (file watching only).
The graph is also something you can talk to : wire the MCP server into your
coding agent and ask it about a run — "why did the Edit on token.js keep
failing?" — the answer arrives in your terminal, and the agent highlights
the exact nodes it's describing on the open graph as it answers. See
Ask your agent about a run .
New here? docs/GUIDE.md walks through the whole thing —
reading the graph, what each signal means, wiring it to your own agent, and what
to do when something looks broken.
Agent sessions stopped being conversations a while ago. They're runs : an
orchestrator spawning subagents, workflows fanning out reviewers, tools
failing and retrying, a human occasionally saying no. rungraph draws that
structure so a 4,000-line transcript becomes something you can actually read:
Time flows down. Your prompts are the backbone; parallel agents fan out
into side-by-side lanes and return to the turn that collected their result.
Tool nodes say what ran , not just which tool: Bash · npm test ×12 ,
Edit · canvas.jsx , Grep · waitForURL . Consecutive calls of the same
tool collapse into one node so a test-fix loop doesn't become a hairball.
Click any node for the full story: prompt and response for turns; every
call's inputs, outputs, errors, and timing for tools; the complete
transcript for subagents. Tool nodes also show the why — the agent's
own narration from just before the call ("Now I'll rerun the tests to
check…").
Human interventions are first-class nodes. A denied permission, an
answered question, a mid-turn interrupt — these are the moments a run
changes direction, and the edges that follow them carry the reason
( after permission denial , retry after failure , after Bash error ).
Workflow runs (multi-agent orchestrations) appear as single nodes you
can drill into: their own graph, phase boxes and all, retries linked to the
attempts they replaced.
Tokens, durations, and models annotate nodes; whole-run totals in the
header.
A graph that renders everything with equal weight points at nothing: a
two-second file read and a forty-minute retry spiral look identical. So
rungraph has an opinion. It derives signals from the run and puts them in a
strip above the canvas — and on a clean run that strip costs zero height,
because a marker you can't trust is worse than no marker.
Click a signal and the graph focuses : those nodes light up, everything else
dims to a quarter — dimmed, never hidden, so the shape you already memorized
stays put. Esc or a click on empty canvas clears it.
The same focus mechanism backs everything else that points at nodes:
Find ( / ) — plain substring over node labels and the files each node
touched . No model, no network, no subprocess; it filters in the browser.
Files — tool and agent nodes carry the paths they touched, including work
done inside subagents , which is where a lot of real editing happens. The
inspector lists every file the run touched with a count; click one to see
exactly which steps touched it.
Live escalation — signals are re-derived on every live-tail update. Go do
something else while the agent works; the strip goes loud only when something
new has actually gone wrong.
The dashboard is for you; the MCP server is for your agent. They are two ends of
one loop, not two products.
npx rungraph mcp --install # one time, then restart Claude Code
npx rungraph mcp --check # is it working? prints exactly what to fix
Then ask, in Claude Code, the kinds of questions a transcript can actually
answer:
which edits in my last run failed — and did any stay broken?
which steps touched src/auth.js , subagents included?
what did the "audit auth module" agent find?
what was the actual error behind that red node?
did it actually run the tests, or just say it did?
Claude calls find_nodes / get_graph / get_detail and answers in your
terminal — your model, your session, fully inspectable. Then it calls
focus_nodes , and the graph you have open lights up the exact nodes the
answer is about — switching to the right run, or opening a browser tab, if
it has to — and hands back a deep link that restores the same highlight for
anyone you paste it to.
You don't have to invent the questions, either: the bottom of the inspector
writes them for you, from the run you're looking at, with a copy button.
Nothing is pinned, prompted, or proxied: rungraph contributes the graph, not the
conversation. The read-only tools work with no server running at all.
With more than one dashboard live — yours, plus a bundle someone sent you
(below) — the MCP aggregates them: list_runs merges every server's runs,
tagged with where they came from, and every other tool routes by run id to the
dashboard actually showing that run.
A run can leave the machine — as a file, on your terms. Say Bilal's agent went
sideways and you could help, or you want to show a colleague where a feature
was actually built.
Bilal exports. Either from the dashboard — share… in the runs pane, check
off runs, review what's about to leave — or by asking his agent:
rungraph export --last 2 --as Bilal
# rungraph: export inventory (full content):
# 2 runs · 143 nodes · 12 of your prompts included
# files touched: 24
# rungraph: wrote acme-2026-08-15.rungraph (412,882 bytes)
The inventory prints every time: people don't realize how much lives in a
transcript, so the tool shows it before it leaves. And export blocks if it
finds a high-confidence secret (AWS keys, GitHub/Slack/API tokens, private-key
blocks — anchored patterns, calibrated for near-zero false positives), listing
exactly where each one is. Resolve with --redact-secrets (placeholders,
everything else verbatim), --structure-only (graph shape, tool names, files
and timings — no prompts, no outputs), or --allow-secrets if they're fixture
keys you've checked.
The file is the transfer. Send the .rungraph over whatever you already
trust — Slack, AirDrop, a repo. rungraph itself never touches a network.
npx rungraph open team-work.rungraph
That serves the bundle on its own ephemeral dashboard — nothing is copied
anywhere; close the process and it's gone; keep the file to re-open it any
time. Every run wears its provenance ("shared by Bilal · team-work.rungraph"),
and the whole loop works on it: signals derive on your rungraph, and your own
agent can be pointed at Bilal's runs — "what went wrong in the bundle Bilal
sent me?" — right alongside your own.
A bundle carries the vendor-neutral IR, so a Codex run exports and opens
identically to a Claude Code one, and opening a bundle needs no adapters at
all. sharedBy is a display string, not an identity — trust a bundle the way
you trust the channel it arrived on.
Link to what you see. copy link in the header captures the current view —
run, selected node, focus — as a URL; focus_nodes returns the same kind of
link, so your agent can hand you something pastable for a PR or an issue. Links
re-execute their query on load (a find link re-finds, a signal link
re-derives), and a link that lands on the wrong dashboard offers a one-click
jump to the one that has the run.
Navigation is Figma-style, built for the tall, skinny graphs real runs
produce:
A minimap (bottom-right) shows the whole run as a strip with a draggable
viewport — errors glow as red beacons; click one to jump straight to the
failure. Runs open at readable zoom: finished runs at the first prompt, live
runs at the latest activity, with follow mode sliding the view as new nodes
stream in.
Everything the UI can do, a coding agent can do over the CLI — no browser, no
prompts, JSON on stdout, logs on stderr, exit codes 0 ok / 1 error / 2 no
runs found. Paste this section into a prompt and an agent can self-serve:
npx rungraph list --json
# {"runs":[{"runId":"claude-code:…:5822df8b-…","kind":"session","title":"Fix flaky auth test",
# "project":"/home/you/dev/app","modifiedAt":"2026-08-11T16:31:06.055Z","active":true,…},…]}
npx rungraph graph ' claude-code:…:5822df8b-… ' --json
# The full Graph IR for that run on stdout:
# {"irVersion":1,"meta":{"runId":"…","kind":"session","title":"…","totals":{"tokens":184230,"toolCalls":57,"agents":4},…},
# "nodes":[{"id":"…","kind":"agent","label":"Investigate flaky test","status":"completed",
# "files":["/home/you/dev/app/src/auth/token.js"],"tokens":{…}},…],
# "edges":[{"kind":"spawn","from":"…","to":"…","label":"Investigate why auth.spec.ts flakes"},…],
# "groups":[…],
# "signals":[{"kind":"retry-storm","severity":"high","nodeIds":["…"],"label":"6 failed Edit calls",
# "reason":"Edit failed 6× across 3 consecutive steps on token.js, …"}]}
# → an agent can read its own past runs: what it spawned, what failed, where the human said no.
npx rungraph find ' claude-code:…:5822df8b-… ' token.js --json
# {"matched":4,"nodeIds":[…],"nodes":[…]}
# → narrow first. A big graph is 20k+ tokens of context to answer one question.
npx rungraph serve --no-open
# {"url":"http://127.0.0.1:4321"} (server stays in foreground; same data over HTTP + SSE live tail)
The same surface is available as MCP tools — see "Ask your agent about a run"
above, or rungraph mcp --install .
The IR is versioned and documented in SCHEMA.md . It is
vendor-neutral, with two adapters: Claude Code (sessions

[truncated]
