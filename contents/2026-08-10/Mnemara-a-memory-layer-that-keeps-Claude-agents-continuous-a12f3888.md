---
source: "https://github.com/mekickdemons-creator/mnemara"
hn_url: "https://news.ycombinator.com/item?id=49242743"
title: "Mnemara – a memory layer that keeps Claude agents continuous"
article_title: "GitHub - mekickdemons-creator/mnemara: Controlled rolling-context conversation runtime for Claude. Stop your agent from running away with your bill. · GitHub"
author: "Mekickdemons"
captured_at: "2026-08-10T12:44:23Z"
capture_tool: "hn-digest"
hn_id: 49242743
score: 1
comments: 1
posted_at: "2026-08-10T12:26:38Z"
tags:
  - hacker-news
  - translated
---

# Mnemara – a memory layer that keeps Claude agents continuous

- HN: [49242743](https://news.ycombinator.com/item?id=49242743)
- Source: [github.com](https://github.com/mekickdemons-creator/mnemara)
- Score: 1
- Comments: 1
- Posted: 2026-08-10T12:26:38Z

## Translation

タイトル: Mnemara – クロード エージェントの継続性を維持する記憶層
記事のタイトル: GitHub - mekickdemons-creator/mnemara: クロード用の制御されたローリング コンテキスト会話ランタイム。エージェントが請求書を持ち逃げするのを阻止してください。 · GitHub
説明: クロードの制御されたローリング コンテキスト会話ランタイム。エージェントが請求書を持ち逃げするのを阻止してください。 - mekickdemons-creator/mnemara

記事本文:
GitHub - mekickdemons-creator/mnemara: クロードのローリングコンテキスト会話ランタイムを制御しました。エージェントが請求書を持ち逃げするのを阻止してください。 · GitHub
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
mekickdemons-クリエイター
/
ムネマラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
187 コミット 187 コミット .github .github 例 例 ノート ノート src/ mnemara src/ mnemara テスト テスト .gitignore .gitignore CONTRI

BUTORS.md CONTRIBUTORS.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
注: deliberium による mnemara と呼ばれる Rust プロジェクトもあります。これは組み込み/サービス システム用のメモリ エンジンです。別のプロジェクト、並行した命名 (どちらも Mnemosyne からインスピレーションを得たもの)。それを探して到着した場合は、そこに向かいましょう。
クロードのエージェントが請求書を持ち逃げするのを阻止してください。
Mnemara は、ロール ドキュメントが再読みされるクロード会話ランタイムです。
毎ターン、システムプロンプトとして固定されます。それはあなたが知るまでは退屈に聞こえるでしょう
それによって何ができるかを理解する: エージェントがあらゆる場合に従わなければならないルールを作成する
最初の API 呼び出しだけでなく、単一の API 呼び出しを実行すると、エージェントがそれらの呼び出しを強制的に実行します。
それ自体。
経験的チェック (2026-05-08): 101 ターンのコーディング セッションで実行
Mnemara 自体、パネルはレーン外書き込みもプロトコルもゼロでした
その役割に対する違反文書。 1ターン目に発動するルールはそのまま
101ターン目に発射。
代表的な例は、examples/roles/sentinel.md です。
それをドロップすると、エージェントは 4 つの失敗について自身の実行を監視します。
エージェントのセッションを高額な事故に変えるモード:
進行なし — N+ は状態を変更せずに同じサブ目標をオンにします。
ポーリング — 同じツール、同じ引数、連続 3 回以上。
ドリフト — ユーザーが求めていないことをしようとしている。
おべっか — 口調だけの反発を受けて結論を覆そうとしている。
いずれかが発火すると、エージェントは発火する代わりに停止して尋ねます。
さらに N ターンの API 予算がかかります。ロール ドキュメントはプレーンな Markdown です - 編集してください
実際に発生する障害モードと一致するようにします。
pip インストール ムネマラ
mnemara init --インスタンス スクラッチ
mnemara ロール --instanceScratch --set-from-url \
https://raw.githubusercontent.com/mekickdemons-creator/mnemara/main/examples/roles/sentinel.md
ムネマル

実行 -- インスタンス スクラッチ
Claude Agent SDK に基づいて構築されています。
Mnemara は、透明なファイルベースのコンテキスト レイヤーで SDK をラップします。
あらゆるターンでモデルが見ているものを正確に見て、形作ることができます。
API 呼び出しごとにロール ドキュメントが再読み込みされ、システム プロンプトとして固定されます。
(Sentinel を機能させるビット)。
最近のターンの構成可能なローリング ウィンドウ (FIFO、行数または
トークンの予算）。
ネイティブ ツールの使用 — Bash、Read、Edit、Write — およびインプロセス WriteMemory
ツールが SDK MCP サーバーとして登録されます。
オプションの MCP ワイヤースルー: 設定で stdio MCP サーバーを宣言し、
Claude Agent SDK はそれらをモデルに公開します。
テキスト TUI ( mnemara run ) と裸のプロンプト ツールキット REPL フォールバック。
~/.mnemara/<instance>/ の下のインスタンスごとのファイルのみの状態 — デーモンなし、
サービスも隠蔽状態もありません。
オプションのメモリ/wiki + LanceDB RAG + Kuzu プロパティ グラフ バックエンド、および
mnemara ウィキ ページの下書きを作成する統合プリミティブを再生し、
クラスター化されたメモリアトムからの役割修正提案。
pip インストール ムネマラ
または、ハッキングしたい場合はソースから:
git clone https://github.com/mekickdemons-creator/mnemara.git
cdムネマラ
Python -m venv .venv
ソース .venv/bin/activate
pip install -e 。
プラットフォームのサポート
Linux と macOS でテスト済み。 Windows は WSL を通じて動作します。
Claude Agent SDK のツール呼び出しでは bash が使用されますが、これは
ネイティブ Windows シェル (cmd / PowerShell)。 Windows を使用している場合は、実行します
Mnemara を WSL ディストリビューション内に配置し、それを「Linux」環境として扱います。
Mnemara は、Anthropic API と通信する Claude Agent SDK 上で実行されます。の
最も簡単な方法は、API キーを設定することです。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
https://console.anthropic.com/ でキーを取得します。 SDK はローカルの
Claude Code がある場合のフォールバックとしての claude CLI のサブスクリプション認証
インストールしてログインしましたが、ドキュメントは

記載されているパスは API キーです。
最初の実行の前に: 権限モデルをざっと読んでください。
セクション。 Mnemara は、LLM Bash/読み取り/書き込み/編集アクセス権を、
それを実行するマシン。サンドボックスではありません。
mnemara init --インスタンス スクラッチ
# (ロールのドキュメント パスの入力を求めるプロンプトが表示されます。空白のままにして後で設定できます)
mnemara role --instanceScratch --setexamples/roles/sentinel.md
mnemara run --インスタンス スクラッチ
--set 引数はロール doc (次のような Markdown ファイル) を指します。
エージェントのシステム プロンプト。リポジトリには、examples/roles/sentinel.md が付属しています。
出発点として。何を入力するかについては、以下のロールのドキュメントを参照してください
あなた自身の。
デフォルトでは、mnemara を実行すると、テキスト チャット パネル (TUI) が開きます。パス
ベアプロンプトツールキット REPL の --no-tui (または MNEMARA_NO_TUI=1 を設定) —
スクリプトまたは非 TTY コンテキストに役立ちます。
+--------------------------------------------------------------+
|ムネマラ: スクラッチ モデル=claude-opus-4-7 役割=... |ヘッダー
+--------------------------------------------------------------+
| |
|あなた: リースのタイムアウトを確認するにはどうすればよいですか? |
|アシスタント:server.py を開いて ... を grep します。
| > ツール: Read(file_path=server.py) |
|結果: ... |
| |チャットログ
+--------------------------------------------------------------+
|ターン: 12/100 |トークン: 14K/200K |モデル: クロード-作品-4-7 |ステータス
+--------------------------------------------------------------+
| > _ |入力
+--------------------------------------------------------------+
入力領域は複数行です。 「スラッシュコマンド」を参照してください。
完全なキーバインド テーブルとスラッシュ コマンド サーフェスについては以下を参照してください。
( /role_doc 、 /context 、 /clear 、 /evict 、 /inbox 、 /models 、 /swap 、
/tokens 、 /export 、 /import 、 /compress reads 、 /skeleton 、 /name など)。
ロール ドキュメントは、エージェントのシステムとなるプレーンな Markdown ファイルです。
プロンプト。 Mnemara は API 呼び出しごとにそれを再読み込みし、スロット 0 に固定します。

f
メッセージ — それは最初のターンだけでなく、すべてのターンに適用されることを意味します。
セッション中にファイルを編集すると、次のターンでファイルが取得されます。
変化します。
これは、エージェントに対する最も強いステアリング信号です。使ってください。
適切な役割ドキュメントは、短い散文ドキュメント (数百から数千) です。
言葉）、順番に答えます：
エージェントとは誰ですか - その身元とこの文書での指示
インスタンス。 (「あなたは Acme リポジトリのコードレビュー担当者です。」
税法を研究する研究助手です。」)
何をすべきか、何をすべきではないか — 範囲、厳しい制約、
避けるべきアンチパターン。
何か問題が発生したときにどのように動作すべきか、いつ停止するか、
いつ助けを求めるべきか、いつエスカレーションすべきか。
スタイルメモ (「簡潔に、謝罪しないでください」)、ツールを含めることができます。
規約 (「コードを書いた後は常にテストを実行する」)、またはドメイン
用語集。必要なスキーマはありません。唯一の機械的要件
ファイルが存在し、読み取り可能であることを意味します。
ループ/ドリフト問題の解決
インタラクティブ エージェント セッションがバーンスルーする最も一般的な理由
トークンには何も表示されません。エージェントがスタックしているためです。
出力が変更されるのを待って、同じツールを繰り返し呼び出します。
ユーザーの実際のリクエストから、隣接するうさぎの穴へと漂流します。
ユーザーが押し返した瞬間に、正しい結論が覆されます。
これらはロールドキュメント形式の問題です。ロール doc は、
エージェントのスパイラルを防ぐルール。それらのルールが含まれていない場合は、
システム プロンプトでは、それらは一貫して適用されるわけではなく、次の場合にのみ再表示されます。
ユーザーはエージェントに忘れずに通知します。
Examples/roles/sentinel.md は
自己監視の役割に関するドキュメント。それをインスタンスのロールとしてドロップすると、
エージェントは、上記の障害モードについて自身の実行を監視します。
(タイムアウト / 進行なし、ポーリング、セマンティック ドリフト、おべっか

逆転)
さらに N ターンを費やすのではなく、停止してユーザーに質問します。
暴走ループ。
mnemara role --instance my-agent --set Examples/roles/sentinel.md
または、リポジトリを複製せずに GitHub から直接ダウンロードします。
mnemara ロール --instance my-agent --set-from-url \
https://raw.githubusercontent.com/mekickdemons-creator/mnemara/main/examples/roles/sentinel.md
--set-from-url はドキュメントを 1 回フェッチします (https のみ、1 MB 上限、UTF-8)。
それを ~/.mnemara/<instance>/role.md に保存します。ムネマラは決して再フェッチしない
実行時の URL — 保存されたローカル コピーが毎回再読み込みされます。
回ります。ローカル コピーを編集してカスタマイズします。
モニタリング風味の作業にそのまま使用するか、テンプレートとして扱います。
ファイルをコピーし、故障モードに一致するようにトリガー条件を編集します。
関心がある場合は、そのコピーをインスタンスに指定します。
Sentinel スタイルの自己監視ルールを上に重ねることもできます。
タスク固有のロール ドキュメント — Sentinel について特別なことは何もありません。それは
エージェントが毎ターン読み取る Markdown ファイル内のテキストだけです。
リポジトリには、examples/roles/ の下に 3 つのロール ドキュメントが同梱されています。
これらは出発点です。それらのいずれかをコピーし、一致するようにルールを編集します
あなたの作品をコピーし、そのコピーをインスタンスに指定します。
自己監視ロールを使用したセッションがどのように見えるかに関するドキュメント
具体例。ロールのドキュメントは、ポーリングを行う場合にエージェントを停止するように指示します。
同じツールを 3 回使用します。ユーザーはエージェントにビルドを待つよう依頼します
終了するには:
あなた: ビルドを実行して、完了したら教えてください
エージェント: > Bash(npm run build)
(出力: 建物...)
ビルドが実行中です。確認させてください。
> Bash(npm run build)
(出力: 建物...)
まだ構築中です。
> Bash(npm run build)
(出力: 建物...)
[センチネル停止]
トリガー: ポーリング
ターン: 4
context: `npm run build` を 3 回連続で呼び出しました。
呼び出しごとに、レポートではなく新しいビルドが開始されました。
1つを実行しています。私'

実際には監視していません。新しく始めます
構築します。
提案された次のステップ: を介して既存のビルド プロセスを待機します。
`wait` するか、存在する場合は `npm run build:status` を確認します。どれ
よろしいでしょうか？
ロール doc がなければ、同じエージェントは喜んで npm run build を呼び出します。
連続 20 回、それぞれが次のいずれかの前に新しいビルドを開始します。
ユーザーが気づいたか、API の予算がなくなった。ロール doc は、
エージェントはターン 4 に通知し、チェックインします。
インスタンスのすべては ~/.mnemara/<instance>/ の下に存在します。
~/.mnemara/<インスタンス>/config.json :
フィールド
意味
役割ドキュメントのパス
ロールドキュメントへの絶対パス。 API 呼び出しごとに再読み取りします。システムプロンプトとして固定されます。
モデル
クロード モデル ID (例: claude-opus-4-7 、claude-sonnet-4-6 、claude-haiku-4-5 )。
max_window_turns
ローリング ウィンドウ サイズ (FIFO)。デフォルトは 20。ユーザーとアシスタントの両方のターンをカウントします。
max_window_tokens
トークンの予算の上限。合計トークンがこれを超えると、ウィンドウは FIFO でトリミングされます。
allowed_tools
{tool, mode, allowed_patterns} ポリシーのリスト。モード ∈ 許可/質問/拒否。
mcp_servers
モデルに配線された標準出力 MCP サーバーのリスト。
ストリーム
true の場合、モデルのテキスト デルタが到着したときにレンダリングします。
bash_timeout_秒
Bash コマンドのタイムアウト。デフォルトは 60。
file_tool_home_only
true の場合、読み取り/書き込み/編集は拒否されます

[切り捨てられた]

## Original Extract

Controlled rolling-context conversation runtime for Claude. Stop your agent from running away with your bill. - mekickdemons-creator/mnemara

GitHub - mekickdemons-creator/mnemara: Controlled rolling-context conversation runtime for Claude. Stop your agent from running away with your bill. · GitHub
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
mekickdemons-creator
/
mnemara
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
187 Commits 187 Commits .github .github examples examples notes notes src/ mnemara src/ mnemara tests tests .gitignore .gitignore CONTRIBUTORS.md CONTRIBUTORS.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
Note: There is also a Rust project called mnemara by deliberium — a memory engine for embedded/service systems. Different project, parallel naming (both inspired by Mnemosyne). If you arrived looking for that one, head over there.
Stop your Claude agent from running away with your bill.
Mnemara is a Claude conversation runtime where the role doc is re-read on
every turn and pinned as the system prompt. That sounds boring until you
realize what it lets you do: write rules the agent has to obey on every
single API call — not just the first one — and the agent enforces them on
itself.
Empirical check (2026-05-08): in a 101-turn coding session run on
Mnemara itself, the panel made zero out-of-lane writes and zero protocol
violations against its role doc. The rule that fires on turn 1 still
fires on turn 101.
The flagship example is examples/roles/sentinel.md .
Drop it in, and the agent watches its own execution for the four failure
modes that turn agent sessions into expensive accidents:
No progress — N+ turns on the same sub-goal with no state change.
Polling — same tool, same args, 3+ times in a row.
Drift — about to do something the user didn't ask for.
Sycophancy — about to reverse a conclusion under tone-only pushback.
When any one fires, the agent halts and asks , instead of burning
another N turns of API budget. The role doc is plain Markdown — edit it
to match the failure modes you actually see.
pip install mnemara
mnemara init --instance scratch
mnemara role --instance scratch --set-from-url \
https://raw.githubusercontent.com/mekickdemons-creator/mnemara/main/examples/roles/sentinel.md
mnemara run --instance scratch
Built on the Claude Agent SDK .
Mnemara wraps the SDK with a transparent, file-based context layer so you
can see and shape exactly what the model sees on every turn.
A role doc re-read on every API call and pinned as the system prompt
(the bit that makes Sentinel work).
A configurable rolling window of recent turns (FIFO, by row count or
token budget).
Native tool use — Bash, Read, Edit, Write — plus an in-process WriteMemory
tool registered as an SDK MCP server.
Optional MCP wire-through : declare stdio MCP servers in config and the
Claude Agent SDK exposes them to the model.
A Textual TUI ( mnemara run ) and a bare prompt-toolkit REPL fallback.
Per-instance, file-only state under ~/.mnemara/<instance>/ — no daemon,
no service, no hidden state.
Optional memory/wiki + LanceDB RAG + Kuzu property graph backends, and a
mnemara replay consolidation primitive that drafts wiki pages and
role-amendment proposals from clustered memory atoms.
pip install mnemara
Or from source if you want to hack on it:
git clone https://github.com/mekickdemons-creator/mnemara.git
cd mnemara
python -m venv .venv
source .venv/bin/activate
pip install -e .
Platform support
Tested on Linux and macOS. Windows works through WSL — the
Claude Agent SDK's tool calls use bash , which is not available on
native Windows shells (cmd / PowerShell). If you're on Windows, run
Mnemara inside a WSL distro and treat that as your "Linux" environment.
Mnemara runs on the Claude Agent SDK, which talks to the Anthropic API. The
easiest way is to set your API key:
export ANTHROPIC_API_KEY=sk-ant-...
Get a key at https://console.anthropic.com/ . The SDK also supports the local
claude CLI's subscription auth as a fallback if you have Claude Code
installed and logged in — but the documented path is the API key.
Before your first run: skim the Permissions model
section. Mnemara gives an LLM Bash/Read/Write/Edit access to the
machine you run it on. It is not a sandbox.
mnemara init --instance scratch
# (prompts for role doc path; you can leave it blank and set it later)
mnemara role --instance scratch --set examples/roles/sentinel.md
mnemara run --instance scratch
The --set argument points at a role doc — a Markdown file that becomes
the agent's system prompt. The repo ships with examples/roles/sentinel.md
as a starting point; see Role docs below for what to put in
your own.
By default mnemara run opens the Textual chat panel (TUI). Pass
--no-tui (or set MNEMARA_NO_TUI=1 ) for the bare prompt-toolkit REPL —
useful for scripting or non-TTY contexts.
+------------------------------------------------------------+
| mnemara: scratch model=claude-opus-4-7 role=... | header
+------------------------------------------------------------+
| |
| you: how do I check the lease timeout? |
| assistant: open server.py and grep for ... |
| > tool: Read(file_path=server.py) |
| result: ... |
| | chat log
+------------------------------------------------------------+
| turns: 12/100 | tokens: 14K/200K | model: claude-opus-4-7 | status
+------------------------------------------------------------+
| > _ | input
+------------------------------------------------------------+
The input area is multi-line. See Slash commands
below for the full keybinding table and the slash-command surface
( /role_doc , /context , /clear , /evict , /inbox , /models , /swap ,
/tokens , /export , /import , /compress reads , /skeleton , /name , etc.).
The role doc is a plain Markdown file that becomes the agent's system
prompt . Mnemara re-reads it on every API call and pins it at slot 0 of
the messages — meaning it applies to every turn, not just the opening one,
and you can edit the file mid-session and the next turn picks up the
changes.
This is the strongest steering signal you have over the agent. Use it.
A good role doc is a short prose document (a few hundred to a few thousand
words) that answers, in order:
Who the agent is — its identity and standing instructions in this
instance. ("You are a code reviewer for the Acme repo." "You are a
research assistant working on tax law.")
What it should and shouldn't do — scope, hard constraints,
anti-patterns to avoid.
How it should behave when something goes wrong — when to halt,
when to ask for help, when to escalate.
You can include style notes ("be terse, no apologies"), tooling
conventions ("always run the tests after writing code"), or domain
glossaries. There is no required schema. The only mechanical requirement
is that the file exists and is readable.
Solving the looping / drift problem
The most common reason an interactive agent session burns through
tokens with nothing to show for it is that the agent gets stuck :
It calls the same tool over and over waiting for output to change.
It drifts from the user's actual request into adjacent rabbit holes.
It reverses a correct conclusion the moment the user pushes back.
These are role-doc-shaped problems. The role doc is where you encode the
rules that keep the agent from spiraling . If those rules aren't in the
system prompt, they aren't applied consistently — they reappear only when
the user remembers to remind the agent.
examples/roles/sentinel.md is a
self-monitoring role doc. Drop it in as your instance's role and the
agent will watch its own execution for the failure modes above
(timeout / no progress, polling, semantic drift, sycophantic reversal)
and halt to ask the user rather than spending another N turns on a
runaway loop.
mnemara role --instance my-agent --set examples/roles/sentinel.md
Or download it directly from GitHub without cloning the repo:
mnemara role --instance my-agent --set-from-url \
https://raw.githubusercontent.com/mekickdemons-creator/mnemara/main/examples/roles/sentinel.md
--set-from-url fetches the doc once (https only, 1 MB cap, UTF-8) and
saves it into ~/.mnemara/<instance>/role.md . Mnemara never re-fetches
the URL at runtime — the saved local copy is what gets re-read each
turn. Edit the local copy to customize.
Use it as-is for monitoring-flavored work, or treat it as a template:
copy the file, edit the trigger conditions to match the failure modes
you care about, and point your instance at the copy.
You can also layer Sentinel-style self-monitoring rules on top of a
task-specific role doc — there's nothing special about Sentinel; it's
just text in a Markdown file that the agent reads on every turn.
The repo ships three role docs under examples/roles/ :
These are starting points. Copy any of them, edit the rules to match
your work, and point your instance at the copy.
What a session looks like with a self-monitoring role doc
Concrete example. Your role doc tells the agent to halt if it polls
the same tool 3 times. The user asks the agent to wait for a build
to finish:
you: run the build and tell me when it's done
agent: > Bash(npm run build)
(output: building...)
Build is running. Let me check.
> Bash(npm run build)
(output: building...)
Still building.
> Bash(npm run build)
(output: building...)
[SENTINEL HALT]
trigger: polling
turn: 4
context: I've called `npm run build` three times in a row and
each call started a new build rather than reporting on the
running one. I'm not actually monitoring; I'm starting fresh
builds.
proposed next step: wait on the existing build process via
`wait` or check `npm run build:status` if it exists. Which
would you like?
Without the role doc, the same agent would happily call npm run build
twenty times in a row, each one starting a fresh build, before either
the user noticed or the API budget ran out. The role doc is what makes
the agent notice on turn 4 and check in.
Everything for an instance lives under ~/.mnemara/<instance>/ :
~/.mnemara/<instance>/config.json :
Field
Meaning
role_doc_path
Absolute path to the role doc. Re-read on every API call. Pinned as the system prompt.
model
Claude model id (e.g. claude-opus-4-7 , claude-sonnet-4-6 , claude-haiku-4-5 ).
max_window_turns
Rolling-window size (FIFO). Default 20. Counts both user and assistant turns.
max_window_tokens
Token-budget cap. The window is FIFO-trimmed once total tokens exceed this.
allowed_tools
List of {tool, mode, allowed_patterns} policies. mode ∈ allow / ask / deny .
mcp_servers
List of stdio MCP servers wired through to the model.
stream
If true, render the model's text deltas as they arrive.
bash_timeout_seconds
Bash command timeout. Default 60.
file_tool_home_only
If true, Read/Write/Edit refuse pa

[truncated]
