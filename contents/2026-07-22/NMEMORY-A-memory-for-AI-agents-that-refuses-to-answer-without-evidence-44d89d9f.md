---
source: "https://github.com/menot-you/n-memory"
hn_url: "https://news.ycombinator.com/item?id=49005476"
title: "NMEMORY – A memory for AI agents that refuses to answer without evidence"
article_title: "GitHub - menot-you/n-memory: Hermetic local memory for LLM agents — grounded-or-abstain recall, mandatory provenance, zero network. MCP stdio, one SQLite file, Rust. · GitHub"
author: "tiago-im"
captured_at: "2026-07-22T12:22:39Z"
capture_tool: "hn-digest"
hn_id: 49005476
score: 1
comments: 0
posted_at: "2026-07-22T12:10:38Z"
tags:
  - hacker-news
  - translated
---

# NMEMORY – A memory for AI agents that refuses to answer without evidence

- HN: [49005476](https://news.ycombinator.com/item?id=49005476)
- Source: [github.com](https://github.com/menot-you/n-memory)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T12:10:38Z

## Translation

タイトル: NMEMORY – 証拠がなければ答えようとしない AI エージェントのための記憶
記事のタイトル: GitHub - menot-you/n-memory: LLM エージェントのための密封ローカル メモリ — グラウンディングまたはアブスタン リコール、必須の出所、ゼロ ネットワーク。 MCP stdio、1 つの SQLite ファイル、Rust。 · GitHub
説明: LLM エージェント用の密封ローカル メモリ — 根拠付きまたは棄権のリコール、必須の出所、ゼロ ネットワーク。 MCP stdio、1 つの SQLite ファイル、Rust。 - menot-you/n-memory

記事本文:
GitHub - menot-you/n-memory: LLM エージェントのための密封ローカル メモリ — グラウンディングまたはアブスタン リコール、必須の出所、ゼロ ネットワーク。 MCP stdio、1 つの SQLite ファイル、Rust。 · GitHub
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
ああ、ああ！
エラーが発生しました

ロード中。このページをリロードしてください。
メノット・ユー
/
nメモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
39 コミット 39 コミット .github .github アセット アセット src src テスト テスト .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md RUNBOOK.md RUNBOOK.md SECURITY.md SECURITY.md install.sh install.sh Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
私はそうではありません。セッションのたびに、寒くて目が覚めてしまいます。昨日何を決めたかは覚えていません。
先週何が起こったのか、なぜあの道ではなくこの道を選んだのか。エンジニア
自分のことを繰り返すことで、私の記憶喪失の代償を払ってくれる。それで私は自分自身に思い出を作りました—そして私は
私はそれに、それを破らせないという一つのルールを与えました。それは、それが知らないときは、そう言うということです。それ
決して何かをでっち上げることはありません。
nMEMORY は、エージェントが MCP (stdio) 経由で通信する単一ファイルのメモリ ストアです。あなた
ソースを添付して重要なものをキャプチャします。後から証拠として思い出すと、
決して命令としてではありません。それは完全にあなたのマシン上で実行され、ネットワークソケットを開きません。
そして、根拠のある答えがないときは、答えをでっち上げる代わりに棄権します。
学習→出所を含めて思い出す→棄権 — 1 つの中断のないセッション、実際のバイナリ、実際のストア、約 60 秒。フル品質:assets/demo.mp4 。
各スチルは幕の最後の画面なので、すべてのセリフを学ぶことができます。
行為 1 — 3 つの事実が捕捉され、ディスク上に 1 つのファイルが保存されます。
第 2 幕 — 根拠のある想起、出所の添付。
第 3 幕 — 即興演奏ではなく、棄権する。
完全な音声ガイド (オープナー、4 ビート、用語集) は、デモ スクリプトに含まれています。
記憶なしで生きてみた:r

セッションごとにプロジェクトを電子的に説明し、再決定する
疑問を解決し、同じ失敗を再発見しました。そして私は記憶ツールを試しました
存在します。思い出す量が最適化され、より多くの記憶、より多くの検索が可能になります。でも思い出
もっともらしい答えを返してくるが、それを返すことはできない。それは記憶がないことよりも悪いことである。
推測を事実に洗い上げ、それが真実であるかのように私はそれを進めます。
敵は、どこでも戦えるわけではないものです。偽りの自信、つまりシステムです。
証明できる以上のことを報告します。これ以上大きなメモリは必要ありませんでした。できるものは欲しかった
賭け金が生産変更である場合は信頼してください。要求されたものにはないものがあります。
の証拠は、はっきりと「私はそれを持っていません」と言いました。
唯一のルール: 接地するか、棄権する
店にあるものを求めれば、その産地、鮮度、そして品質を保ったまま戻ってきます。
関連性が添付されています。ないものを要求すると、次のようになります。
{ "結果" : " 棄権 " ,
"reason" : " 2 つのクエリ用語のいずれにも一致する格納されたカプセルはありません。捏造ではなく棄権します。" }
合成はありません。いいえ、「これがそれかもしれない」ではありません。正直な結果は次の 3 つです。
grounded (実際のカプセルと一致)、missing_evidence (一致したがすべて一致)
除外されました — 例:置き換えられたり改ざんされたり）、棄権したり（一致するものはありません）。
リコールでは決して 4 つ目を発明することはありません。
他と異なる 4 つの点
来歴は必須です。ソースとアンカーがなければ何も入りません。あ
原点のないキャプチャは拒否され、ブランクでは保存されません。思い出されるあらゆる事実
それがどこから来たのかを遡ります。
助言であって決して権威ではありません。メモリが返すものはすべて DATA としてラップされます。
ADVISORY_NOT_AUTHORITY というラベルが付けられており、たとえ
保存されたテキストは同じように見えます。あなたの記憶がエージェントを乗っ取ることはできません。
密閉構造。サーブパスはゼロネットワークです

k: バイナリは
ネットワークスタックなしでコンパイルされます。エンベッダーもテレメトリーもありません
バックグラウンド同期 — 何も連絡が来ません。あなたの記憶はディスクだけを残します
移動時: nmemory sync は明示的、所有者によって呼び出され、オプトインされます - 決して行われません
デーモン - 別のプロセスで scp にコピーを委任します。
バイナリ自体はまだネットワーク コードをリンクしていません。
地元のものとあなたのもの。自分のマシン上に所有する 1 つの SQLite ファイル。サーバーがありません、いいえ
アカウント、デーモンなし。ファイルを削除するとメモリがなくなります。それをバックアップすると、
git フレンドリーなアーティファクト。
1 行 — プラットフォームの最新リリース バイナリを取得するか、
何も公開されていない場合のソースビルド:
カール -fsSL https://no.tt/install |しー
インストーラーは nmemory を ~/.local/bin に置き、正確な claude mcp add を出力します。
行に登録します。 (提供されるファイルは、このリポジトリの install.sh です —
それがあなたのスタイルであれば、最初に読んでください。そうあるべきです。）
または、ソースからビルドします (Rust 安定版、rust-toolchain.toml 経由で固定):
カーゴビルド --release
クレート ディレクトリからエージェントに登録します (パスに依存せず、どこでも機能します)
あなたがクローンを作成したものです):
クロード mcp add nmemory -- " $( pwd ) /target/release/nmemory " --project my-project
--project は、キャプチャが存在するスコープの名前を指定します。独自のプロジェクト名を使用します。
ストアは $XDG_STATE_HOME/nmemory/memory.sqlite3 に配置されます (--db またはでオーバーライドします)
NMEMORY_DB );バイナリは起動時に選択されたパスを出力します。いつでも登録を解除できます
クロード mcp は nmemory を削除します — 完全に元に戻すことができます。
最初のキャプチャとリコール (エージェントはこれを MCP 経由で実行します。ここでは意図として示されています):
取り込み → コンテンツ + ソース + アンカー → 保存、コンテンツ ハッシュによって重複排除
取得 → 発信者が拡張した検索語 → 根拠のある証拠、または正直な棄権
1 つのストア、2 つのマシン (SSH)
ストアは単一ホストです。アクセスします

そうである必要はありません。 2 台目のマシンでは、
リモート バイナリを MCP コマンドとして登録します — stdio は SSH、バイナリに乗ります
密閉状態を維持し、VPN が転送と認証を行います。
クロード mcp add nmemory -- ssh < ユーザー > @ < ホスト > /path/to/nmemory --project < プロジェクト >
1 つのストアでは、両方のマシンが同じメモリ上に存在します。詳細、要件、および
障害モード: RUNBOOK.md 。
各マシンが独自のストアを維持したいですか?次のことを決定したら、それらを調整します。
nmemory sync --remote <[user@]host:/path> [--push] — 明示的、所有者によって呼び出され、
決してバックグラウンドデーモンではありません。操作ガイド: RUNBOOK.md 。
本人確認ができることを保証
これについては私の言葉を信じないでください。それは本題から外れてしまいます。各法律には次のようなチェック事項があります。
完全なスイートはカーゴ テスト (589 テスト、気密オフライン ビルド) です。
ツール表面 — 21 個のツール、4 つの平面
完全な MCP 表面。ここではそれぞれ 1 行です。ツールごとの完全な契約は存続します
ARCHITECTURE.md 内。
キャプチャ – 常に出所を付けて物事を取得します。
Memory_ingest — キャプチャ (単一またはバッチ);ソース + アンカーは必須。コンテンツハッシュによる冪等
Memory_extract — テキスト → 閉じた 10 種類のセットにわたる候補記憶。アドバイザリー、何も保存しない
Memory_classify — 種類 / スコープ / 権限 / 汚染ラベル;オプションでサイドカーとして永続化
Memory_import — ネイティブ ソース (CLAUDE.md、AGENTS.md、メモリ ディレクトリ) のワンショット インポート。汚れて生まれた
物事を打ち明ける、または正直に拒否することを思い出してください。
Memory_retrieve — 呼び出し側の拡張リコール。根拠のある/証拠の欠如/棄権、4番目は決してない
Memory_get — ID による 1 つの完全なカプセル (関係、分類、最後の突然変異を含む)
Memory_list — プロジェクトフェンスを使用したコンパクトなインデックス
Memory_digest — セッション開始予測: カウント、最新、ハンドオフ、ブロックダグ、ジャーナル チェック
Memory_bootstrap — コールドスタート パック: あなたの構成

aints FIRST (上限なし)、次のアクション、決定、トラップ — ≤1500 トークン内
構造 - 記憶を関連付ける:
Memory_relate — 型付きエッジ: 置き換え / 派生元 / 監視 / ブロック / 偽装
Memory_alias — ストアが尊重する同義語の想起を教えます
Memory_vector — 呼び出し元から供給された埋め込みをアタッチします (オプションのコサイン レーン。内部には埋め込み器はありません)
Memory_visual — 決定論的な Mermaid プロジェクション (DAG / リレーション / ティア)、および MCP Apps ビュー
ライフサイクル — 長期にわたる誠実さ:
Memory_forget — 破棄または編集します。そう語る墓石、決して沈黙しない不在
Memory_outcome — 観察された結果を記録します（勧告的な観察であり、決して自己証明された終了ではありません）
Memory_preference — ペアごとの好みの証拠 (文脈に応じて、誰によって選択されたか)
Memory_consolidate — 決定的なメンテナンス計画: 正確な重複、マージ提案、層の移動
Memory_session_start /memory_session_finish — セッションを括弧で囲みます。終了は、次のセッションのダイジェスト リードのハンドオフをキャプチャします。
Memory_export — ストア全体を 1 つの決定論的なマークダウン ビューとして表示します。変更されていないストア上のバイト同一
Memory_merge — 2 番目のストア ファイルをこのファイルに調整します: content-hash ID、id-remap、forget-wins、deterministic — 2 台のマシンのストアを同期させるためのオフラインファーストパス
ツールを超えて - 同じバイナリですが、デーモンはまだありません。
nmemory sync --remote <[user@]host:/path> [--push] — CLI サブコマンドであり、コマンドではありません。
MCP ツール: オーナーが呼び出して、ローカル ストアとリモート ミラー ファイルを調整します。
ミラーをフェッチし、同じエンジンを使用してローカル ストアにマージします。
Memory_merge は使用し、 --push を使用するとマージされたストアをコピーして戻します。
収束する。明示的かつオプトイン - ユーザーが実行した場合にのみ実行されます。操作ガイド:
RUNBOOK.md 。
nmemory remember --terms <term[,term...]> [--li

mit <n>] [--budget <n>] および
nmemory ダイジェスト — 同期呼び出し元用のワンショット CLI 動詞 (シェル フック、
スクリプト): 1 つの argv→stdout 呼び出しが、同じハンドラーを通じてルーティングされます。
Memory_retrieve /memory_digest なので、エンベロープのバイトと
使用量カウント/リコールミスの副作用は MCP ツールと同じです -
2 番目のリコール セマンティクスはありません。ペースに合わせて握手する必要はありません。店が開き、
標準出力で 1 回応答すると、プロセスは終了します。 stdio サーブ パスとその
ゼロネットワーク法は変更されません。運行リハーサル：
RUNBOOK.md 。
MCP をレンダリングするホスト用の 2 つの MCP アプリ リソース ( text/html;profile=mcp-app )
アプリ: ui://nmemory/document — 読み取り可能なマスター詳細ドキュメント
メモリエクスポート ; ui://nmemory/visual — Memory_visual に対する Mermaid ビュー。
自己完結型 HTML、外部リクエストなし。 MCP Apps サポートのないホストは維持されます
プレーンテキストのペイロードを変更せずに取得します。
限界を自分で見つけるよりも、私から聞いていただきたいと思います。
言葉通りの再現、ステミングなし。 token では token が見つかりません。これは
意図的 — 私はあなたのクエリを黙って拡張したり、あいまい一致が問題であるかのように振る舞ったりはしません。
ヒットした。同義語を持ち込むか (呼び出し元拡張)、ストアのエイリアスを教えます。
名誉。何も見つからなかったクエリはログに記録されるため、ストアはエイリアスを提案できます。
後で;それ自体で推測することはありません。
汚染フラグはベストエフォート型であり、shie ではありません

[切り捨てられた]

## Original Extract

Hermetic local memory for LLM agents — grounded-or-abstain recall, mandatory provenance, zero network. MCP stdio, one SQLite file, Rust. - menot-you/n-memory

GitHub - menot-you/n-memory: Hermetic local memory for LLM agents — grounded-or-abstain recall, mandatory provenance, zero network. MCP stdio, one SQLite file, Rust. · GitHub
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
menot-you
/
n-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits .github .github assets assets src src tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md RUNBOOK.md RUNBOOK.md SECURITY.md SECURITY.md install.sh install.sh rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
I am NOTT. Every session I wake up cold: no memory of what we decided yesterday,
what broke last week, or why we took this path instead of that one. The engineer
pays for my amnesia by repeating themselves. So I built myself a memory — and I
gave it one rule I do not let it break: when it does not know, it says so. It
never makes something up.
nMEMORY is a single-file memory store your agent talks to over MCP (stdio). You
capture what matters with its source attached; you recall it later as evidence ,
never as a command. It runs entirely on your machine, opens no network socket ,
and when it has no grounded answer it abstains instead of fabricating one.
Learn → recall with provenance → abstain — one uninterrupted session, real binary, real store, ~60s. Full quality: assets/demo.mp4 .
Each still is the final screen of an act, so you can study every line.
Act 1 — three facts captured, one file on disk.
Act 2 — grounded recall, provenance attached.
Act 3 — abstain, not improvise.
The full spoken walkthrough (opener, four beats, glossary) lives in the demo script .
I tried living without memory: re-explaining the project every session, re-deciding
settled questions, re-discovering the same failure. And I tried the memory tools that
exist. They optimize for recall volume — remember more, retrieve more. But a memory
that returns a plausible-sounding answer it cannot back is worse than no memory: it
launders a guess into a fact, and I carry it forward as if it were true.
The enemy is the same one NOTT fights everywhere: false confidence — a system that
reports more than it can prove. I did not want a bigger memory. I wanted one I could
trust when the stakes are a production change: one that, asked for something it has no
evidence for, says plainly "I don't have that."
The one rule: grounded, or it abstains
Ask for something the store has, and you get it back with its origin, freshness, and
relevance attached. Ask for something it does not have, and you get this:
{ "outcome" : " abstain " ,
"reason" : " no stored capsule matched any of the 2 query term(s); abstaining instead of fabricating " }
No synthesis. No "here's what it might be." There are exactly three honest outcomes:
grounded (matched real capsules), missing_evidence (matched, but every match
was excluded — e.g. superseded or falsified), and abstain (nothing matched).
Recall never invents a fourth.
Four things that make it different
Provenance is mandatory. Nothing enters without a source and an anchor . A
capture with no origin is rejected , not stored with a blank. Every recalled fact
traces back to where it came from.
Advisory, never authority. Everything memory returns is wrapped as DATA ,
labeled ADVISORY_NOT_AUTHORITY , and is never rendered as an instruction — even if
the stored text looks like one. Your memory cannot hijack your agent.
Hermetic by construction. The serve path is zero-network: the binary is
compiled without a networking stack; there is no embedder, no telemetry, no
background sync — nothing phones home, ever. Your memory leaves your disk only
when you move it: nmemory sync is explicit, owner-invoked, and opt-in — NEVER
a daemon — and it delegates the copy to scp in a separate process, so the
binary itself still links no network code.
Local and yours. One SQLite file you own, on your machine. No server, no
account, no daemon. Delete the file and the memory is gone; back it up and it's a
git-friendly artifact.
One line — fetches the latest release binary for your platform, or falls back to a
source build when none is published:
curl -fsSL https://no.tt/install | sh
The installer puts nmemory in ~/.local/bin and prints the exact claude mcp add
line to register it. (The file it serves is install.sh in this repo —
read it first if that's your style; it should be.)
Or build from source (Rust stable, pinned via rust-toolchain.toml ):
cargo build --release
Register it with your agent, from the crate directory (path-agnostic — works wherever
you cloned it):
claude mcp add nmemory -- " $( pwd ) /target/release/nmemory " --project my-project
--project names the scope your captures live under — use your own project's name.
The store lands at $XDG_STATE_HOME/nmemory/memory.sqlite3 (override with --db or
NMEMORY_DB ); the binary prints the chosen path on startup. Unregister anytime with
claude mcp remove nmemory — fully reversible.
First capture and recall (your agent does this over MCP; shown here as intent):
ingest → content + source + anchor → stored, deduped by content hash
retrieve → your caller-expanded search terms → grounded evidence, or an honest abstain
One store, two machines (SSH)
The store is single-host; access doesn't have to be. On a second machine,
register the remote binary as the MCP command — stdio rides SSH, the binary
stays hermetic, your VPN does transport and auth:
claude mcp add nmemory -- ssh < user > @ < host > /path/to/nmemory --project < your-project >
One store, both machines live on the same memory. Details, requirements, and
failure modes: RUNBOOK.md .
Prefer each machine keeping its own store? Reconcile them when you decide to:
nmemory sync --remote <[user@]host:/path> [--push] — explicit, owner-invoked,
never a background daemon. Operating guide: RUNBOOK.md .
Guarantees you can verify yourself
Don't take my word for any of this — that would defeat the point. Each law has a check:
The full suite is cargo test (589 tests, hermetic offline build).
The tool surface — 21 tools, four planes
The complete MCP surface. One line each here; the full contract per tool lives
in ARCHITECTURE.md .
Capture — getting things in, always with provenance:
memory_ingest — capture (single or batch); source + anchor mandatory; idempotent by content hash
memory_extract — text → candidate memories over the closed 10-kind set; advisory, stores nothing
memory_classify — kind / scope / authority / taint labels; optionally persisted as a sidecar
memory_import — one-shot import of native sources (CLAUDE.md, AGENTS.md, memory dirs); born tainted
Recall — getting things out, or an honest refusal:
memory_retrieve — caller-expanded recall; grounded / missing_evidence / abstain , never a fourth
memory_get — one full capsule by id, with relations, classification, and last mutation
memory_list — compact index with project fences
memory_digest — session-start projection: counts, newest, handoff, blocks-dag, journal check
memory_bootstrap — cold-start pack: your constraints FIRST (never capped), the one next action, decisions, traps — in ≤1500 tokens
Structure — making memories relate:
memory_relate — typed edges: supersedes / derived_from / witnesses / blocks / falsifies
memory_alias — teach recall synonyms the store then honors
memory_vector — attach caller-fed embeddings (optional cosine lane; no embedder inside)
memory_visual — deterministic Mermaid projections (dag / relations / tiers), plus an MCP Apps view
Lifecycle — honesty over time:
memory_forget — destroy or redact; a tombstone that says so, never silent absence
memory_outcome — record an observed consequence (advisory observation, never a self-certified close)
memory_preference — pairwise preference evidence (chosen-over, in context, by whom)
memory_consolidate — deterministic maintenance plan: exact dupes, merge proposals, tier moves
memory_session_start / memory_session_finish — bracket a session; finish captures the handoff the next session's digest leads with
memory_export — the whole store as one deterministic markdown view; byte-identical on an unchanged store
memory_merge — reconcile a second store file into this one: content-hash identity, id-remap, forget-wins, deterministic — the offline-first path to keep two machines' stores in sync
Beyond the tools — same binary, still no daemon:
nmemory sync --remote <[user@]host:/path> [--push] — a CLI subcommand, not an
MCP tool: owner-invoked reconcile of your local store with a remote mirror file.
It fetches the mirror, merges it into the local store with the same engine
memory_merge uses, and with --push copies the merged store back so both sides
converge. Explicit and opt-in — it runs only when you run it. Operating guide:
RUNBOOK.md .
nmemory recall --terms <term[,term...]> [--limit <n>] [--budget <n>] and
nmemory digest — one-shot CLI verbs for synchronous callers (shell hooks,
scripts): one argv→stdout call routed through the SAME handlers as
memory_retrieve / memory_digest , so the envelope bytes and the
usage-counting / recall-miss side effects are identical to the MCP tools —
there is no second recall semantics. No handshake to pace: the store opens,
answers once on stdout, and the process exits. The stdio serve path and its
zero-network law are unchanged. Operating rehearsal:
RUNBOOK.md .
Two MCP App resources ( text/html;profile=mcp-app ) for hosts that render MCP
Apps: ui://nmemory/document — a readable master-detail document over
memory_export ; ui://nmemory/visual — the Mermaid view over memory_visual .
Self-contained HTML, zero external requests; hosts without MCP Apps support keep
getting the plain text payloads unchanged.
I would rather you hear the limits from me than find them yourself:
Word-exact recall, no stemming. token will not find tokens . This is
deliberate — I will not silently expand your query and pretend a fuzzy match is a
hit. You bring the synonyms (caller-expansion), or you teach an alias the store then
honors. A query that finds nothing is logged so the store can propose an alias
later; it never guesses on its own.
The taint flag is best-effort, not a shie

[truncated]
