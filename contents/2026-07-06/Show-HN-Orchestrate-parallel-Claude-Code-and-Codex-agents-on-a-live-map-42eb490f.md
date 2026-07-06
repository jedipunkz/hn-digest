---
source: "https://github.com/Latand/live-log-viewer-next"
hn_url: "https://news.ycombinator.com/item?id=48804797"
title: "Show HN: Orchestrate parallel Claude Code and Codex agents on a live map"
article_title: "GitHub - Latand/live-log-viewer-next · GitHub"
author: "latand6"
captured_at: "2026-07-06T15:02:35Z"
capture_tool: "hn-digest"
hn_id: 48804797
score: 2
comments: 0
posted_at: "2026-07-06T14:06:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Orchestrate parallel Claude Code and Codex agents on a live map

- HN: [48804797](https://news.ycombinator.com/item?id=48804797)
- Source: [github.com](https://github.com/Latand/live-log-viewer-next)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T14:06:25Z

## Translation

タイトル: HN の表示: ライブ マップ上で並列クロード コードとコーデックス エージェントを調整する
記事のタイトル: GitHub - Latand/live-log-viewer-next · GitHub
説明: GitHub でアカウントを作成して、Latand/live-log-viewer-next の開発に貢献します。
HN テキスト: こんにちは、HN!画面全体で 1 つの会話が行われる、AI エージェント用の単一のチャット インターフェイスになぜ誰もが平気でいるのか (そして実際に何度もコピーしようとするのか)、私にはまったく理解できませんでした。私は多くのコーディング エージェント (Claude Code、Codex) を並行して実行することがよくありますが、しばらくすると、どのサブエージェントがどこにあり、誰が誰を生成したのか、どれが機能していてどれが機能していないのかがわからなくなります。エージェントが終了するか静かに死ぬかは 30 分後にわかります。エージェントに関する情報はすでにディスク上にありました。すべてのセッションは ~/.claude または ~/.codex の下に JSONL トランスクリプトを書き込み、すべてのサブエージェントが独自のファイルを取得します。そこで、これらのファイルを読み取り、選択したファイルを拡張可能なツールコール カードを備えたライブ チャット フィードとして表示するローカル Web UI を構築しました。データベースはほとんどなく、すべてが 127.0.0.1 で実行されます。最も困難な部分は (今もそうですが) 親子関係ツリーです。スキャナーはトランスクリプトを読み取り、リンクを回復し、ルートの会話から最後のバックグラウンド シェル タスクまで、プロジェクトごとに 1 つのツリーを構築します。何が起こっているかを確認するために、すでにクロード コードとコーデックスを使用している場合はすでに存在する実際のプロセスと実際のログから構築されたレイヤーが必要でした。しかし、ビューアはワークフロー内で何も変更できないという点だけにとどまりませんでした。オーケストレーションは引き続きヘッドレスで実行され、アプリを開くことはできず、とにかく存在する同じ tmux ペインとログを読み取ります。しかし、サブエージェントがすでに作業を行っていて、もう 1 つの修正が必要な場合は、そのペインにメッセージを直接投げることができます。そのため、サブエージェントを自由に再利用できます。また、クロード コードと

サイドバーのコーデックス制限ペイン、音声転写、レビュー サイクル ループ、およびワークフロー (ただし、最新のものは適切にテストしていません)。今後もその取り組みを続けます。主にモデルをテストするために、Claude Fable 5 と Codex を使用して全体をバイブコーディングしました。以前に、同様のアイデアのより単純なバージョンで Codex を苦しめたところ、非常に苦労したため、そうしてよかったと思います。寓話はそれを釘付けにしました おそらく私がこれを作成したのは、私が ADHD を持っていて、頭の中に 10 人の並行エージェントを保持するのが難しいからですが、何かが私に一人ではないことを教えてくれます 試してみて、貢献してください: 'bunx Agent-log-viewer とセットアップ' をエージェントに送信してください。 Linux 用で、macOS のサポートは限定されていますが、Windows はまだサポートされていません。

記事本文:
GitHub - Latand/live-log-viewer-next · GitHub
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
ラタン語
/
ライブログビューア次へ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 他のアクションを開く m

enu フォルダーとファイル
201 コミット 201 コミット .claude/ skill/ review-loop .claude/ skill/ review-loop .github/ workflows .github/ workflows bin bin docs docs public public scripts scripts src src .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md bun.lock bun.lock eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package.json package.json postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Agent-log-viewer は、生の Codex / Claude Code エージェントを変換するローカル Web UI です
読み取り可能なライブ更新チャット フィードにログインします。あらゆるセッションを発見し、
サブエージェント、Codex コンパニオン ジョブ、およびマシン上のバックグラウンド シェル タスク、リンク
それらを親→子ツリーに変換し、選択したツリーをリアルタイムで追跡します。
すでにディスク上にあるファイルに対してすべてがローカルで実行されます。データベースはありませんし、
外部サービスなし - アプリは ~/.claude と ~/.codex を読み取り、何をレンダリングしますか
それは見つかります。
クロード コード セッション ( ~/.claude/projects/**/*.jsonl ) とそのセッション
サブエージェント、チャットとしてレンダリング: ユーザーのバブル、アシスタントの散文、ツール呼び出し
✓/✗ ステータスのカードと拡張可能な出力。
コマンドによる Codex CLI セッション ( ~/.codex/sessions/**/rollout-*.jsonl )
カード、パッチ、サービス イベント。
Codex コンパニオン ジョブ ( ~/.claude/plugins/data/codex-openai-codex/state )
ワンクリックで、各ジョブの背後にある完全なロールアウト セッションにジャンプします。
バックグラウンド シェル タスク (OS では claude-<uid>/**/tasks/*.output)
temp dir — Linux の場合は /tmp、macOS の場合は $TMPDIR) — 元の Bash
コマンドはセッション記録から復元され、上に表示されます。
端子出力。
親子ツリー: セッション → サブエージェント → コーデックス ジョブ → ロールアウト → バックグラウンド
タスク、ビルド

サーバー側でトランスクリプトをスキャンする（追加のみの増分、
キャッシュ済み — ウォーム /api/files ポーリングは約 100 ミリ秒のままです)。
ライブアクティビティ: コンテンツベースのバッジ - 作業中のトランスクリプト読み取り
それはターンの途中であり、最後のアシスタントメッセージが到着すると完了します。
ディープリンク: すべての選択は URL ( #f=<path> ) に反映されるため、
リンクをクリックすると、その正確なログが開きます。
プロジェクト スキーム: 各プロジェクトはパン、ズーム可能な図、つまりルートです
上に会話、1 世代下に生成されたエージェント、色付きの矢印
エンジン。静かなブランチとタスクは、最も近い祖先の下に折りたたまれます。
英語またはウクライナ語の UI、モデル チップ ( fable-5 、 gpt-5.5 、 Sonnet …)、
永続的な状態を持つ折りたたみ可能なツリー、フォローモードの自動スクロール、サービスイベント
トグルとラインフィルター。
ルート会話を生成されたエージェントに接続するセッション親子関係ツリー:
パッケージは npm に公開されるため、クイックスタートにはクローンは必要ありません。
bunx エージェント ログ ビューア
# または
npx エージェント ログ ビューア
これにより、127.0.0.1:8898 でサーバーが起動し、ブラウザが開きます。
バンインストール
バンランビルド
bun start --port 8898 --hostname 127.0.0.1
# http://127.0.0.1:8898/ を開く
npm の場合 (開始スクリプトにフラグを転送する -- に注意してください):
npmインストール
npm ビルドを実行する
npm start -- --port 8898 --hostname 127.0.0.1
start は最後の build の出力を提供するため、最初に build を実行します。のために
開発、bun dev はホットリロードでアプリを実行します (高 OS が必要です)
大規模なホーム ディレクトリのファイル監視制限)。
前提条件: Node ≥ 20.9、および bun または npm/pnpm。 tmux はオプションです — を参照してください。
プラットフォームのサポート。
エージェントログビューア [オプション]
オプション
説明
-p、--ポート<n>
ローカルサーバーのポート (デフォルトは 8898 )。
-H、--ホスト名 <h>
バインドアドレス (デフォルトは 127.0.0.1 )。
--テールスケール
テールネット内でビューアを公開します (下記を参照)。
--新しいトークン
新しいアクセスキーを生成し、

古いクッキーを検証します。
--no-open
起動時にブラウザを開かないでください。
-v、--バージョン
バージョンを印刷します。
-h、--ヘルプ
使用状況を表示します。
プラットフォームのサポート
Linux はネイティブ ターゲットです。プロセス検出は /proc を直接読み取ります。 macOSは
ps および lsof をシェルアウトするポータブル バックエンドを通じてサポートされます。
代わりに - 同じライブプロセス検出、tmux Composer ターゲティング、エージェント
spawn/kill とバックグラウンド タスクの検出、サブプロセスのオーバーヘッドが少しだけ増加
スキャンごとに。バックエンドは process.platform によって自動的に選択されます (「
src/lib/proc/ ); VIEWER_PROC_BACKEND=portable はポータブル パスを強制的にオンにします
Linux もテスト用に。
このパッケージは Linux と macOS をサポートしています。パッケージの os フィールドが Windows をブロックする
EBADPLATFORM とともにインストールされます。 WSLはLinuxとして動作します。
tmux はオプションです。それがなければ、ログ表示、親子関係ツリー、ライブアクティビティ
ディープリンクはすべて機能します。コンポーザー、エージェントのスポーン/キル、およびペインへの再開
機能には tmux が必要です (macOS では brew install tmux、またはディストリビューションのパッケージが必要です)
Linux)。
UI はデフォルトで英語になり、プロジェクト内にコンパクトな EN/UK 切り替えが表示されます。
リストのヘッダー。ロケールは最初に localStorage キー llv_lang として解決されます。
次にブラウザ言語 (ブラウザが優先する場合はウクライナ語)、次に英語です。
CLI メッセージはデフォルトでは英語ですが、次のコマンドを使用するとウクライナ語に切り替わります。
LLV_LANG=uk または LANG / LC_ALL の uk_* 値。
エージェントと会話する作曲家には、メッセージを口述するためのマイク ボタンがあります。によって
デフォルトの文字起こしは、より高速なウィスパー経由で完全にローカルで実行されます。音声は残りません。
機械。 script/setup-whisper.sh を 1 回実行して、ローカル エンジンをインストールします。
2 つのクラウド バックエンドは、マシンごとの明示的なオプトインとして利用できます (UI ではありません)
トグル): ChatGPT (ローカルの Codex ログインを再利用) および Celebrities Scribe (
ライブのストリーミング文字起こしが含まれるのは 1 つだけです）。で選択
LLV_TRANSCRIBE_BACKEND=local|chatgpt|elevenlabs または

バックエンドを書くことで
~/.config/agent-log-viewer/transcribe-backend に名前を付けます。ローカルがデフォルトです。
セットアップ、キーの場所については、docs/transcription.md を参照してください。
そしてトラブルシューティング。
ビューアは実装→レビューのサイクルを調整します: 寿命の長い実装者
tmux のエージェント、完全な差分にわたるラウンドごとの新しい読み取り専用レビュアー、
結果の自動中継と、スキーム ビューの評決デッキ。一つ始めてください
会話ペインの上のフローチップから。プリセットペアエンジンと
役割ごとの推論の労力 (例: Codex high → Fable )。
ラウンドプロトコルについては docs/review-loop.md を参照してください。
プリセット、HTTP オートメーション API、およびトラブルシューティング。クロード・コードのスキル
エージェントからのフローを駆動するためのパッケージは .claude/skills/review-loop/ にあります —
クローン内で動作するエージェントは、それを自動的に取得します。
bunx エージェント ログ ビューア --tailscale
--tailscale は 127.0.0.1 でローカル サーバーを起動し、それを内部で公開します。
フォアグラウンド テールスケール経由のテールネットは、<port> プロセスを提供します。一般の人々
インターネット (ファネル) は決して使用されません。
CLI は 32 文字のアクセス キーを生成し、それをテールネット URL に追加します。
?k=... 、最初のアクセスの後、サーバーは llv_auth Cookie を 30 に設定します。
日々。 --new-token は新しいキーを生成し、古いキーをすべて即座に無効にします
cookie — 各リクエストはハッシュを現在のトークンと比較します。
以前のキーで作成されたキーは通過しなくなりました。
端末は、電話でスキャンするための QR コードとともにテールネット URL を印刷します。
同じ QR を Web UI 内で使用できます: プロジェクトの QR アイコン ボタン
リストヘッダーは、コードとリンクをテキストとして含むポップオーバーを開きます（コピー付き）
ボタン）。 QR は完全にクライアント側でレンダリングされます (qrcode パッケージ、
外部リクエスト）、すでに承認されたクライアントにのみ提供されます - 同じ
src/proxy.ts のトークン ゲートも /api/access を保護します。いつ

彼はサーバー
--tailscale なしで実行すると、ボタンに開始のヒントが表示されます
bunx エージェントログビューア --tailscale 。
この URL にテールネットでアクセスできる人は誰でも、すべてのエージェントのトランスクリプトを読むことができます。
セッションに到達した機密データが含まれ、コマンドを実行できる
/api/tmux および /api/spawn を通じて。テールネット URL をシークレットとして扱います。
他の人に転送しないでください。
ログ API は、ホワイトリストに登録されているパスのいずれかに解決されないパスを拒否します。
ログルート ( src/lib/scanner/roots.ts を参照)。変異するエンドポイントが存在します:
/api/tmux は tmux セッションにキーを送信し、/api/spawn はコマンドを開始します。
デフォルトでは、CLI は 127.0.0.1 にバインドされます。 --tailscale を使用すると、アクセスが公開されます
テールスケール サーブを介してテールネット内にあり、トークン ゲートによって保護されています
src/proxy.ts 。非ループバック バインドでもトークン モードが強制されます。あらゆる URL を扱う
?k= を資格情報として含みます。
すべてオプションです。転写変数は完全に文書化されています。
docs/transcription.md 。
ビューアは、その状態を標準の XDG ディレクトリに保存します。
パッケージ:
~/.config/agent-log-viewer/ — アクセストークン、transcribe-backend、および
elevenlabs-API キー 。
~/.cache/agent-log-viewer/whisper-venv — ローカルの転写
仮想環境。
従来のライブ ログ ビューア パスは引き続き読み取り専用フォールバックとして受け入れられるため、
既存のセットアップは変更せずに機能し続けます。
ARCHITECTURE.md : src/app/api/* のルート ハンドラーを参照してください。
src/lib/scanner/* の下の純粋なスキャナー パイプライン (検出 → 記述 → アクティビティ)
→ モデル → リンク)、 src/components/* の下の React コンポーネント。キャッシュは存続します
globalThis と、dev ホットリロードを生き残る。
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
1
v0.8.0 — ボード高速パス、モバイル シェル、コピー アフォーダンス
最新の
2026 年 7 月 6 日
パッケージ
0
読み込み中にエラーが発生しました。これをリロードしてください

のページ。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Latand/live-log-viewer-next development by creating an account on GitHub.

Hi HN! I never understood why everybody is fine with (and actually try to copy it multiple times) a single chat interface for ai agents, where you have one conversation on the whole screen. I frequently run many coding agents in parallel (Claude Code, Codex), and after some time I lose the picture of where which subagent is and who spawned whom, which is working and which it not. Agents finish or die quietly and I find out half an hour later And the information about the agents was already sitting on the disk. Every session writes a JSONL transcript under ~/.claude or ~/.codex, every subagent gets its own file. So I built a local web UI that reads these files and shows the selected one as a live chat feed with expandable tool-call cards. There is almost no database and everything runs on 127.0.0.1 The hardest part was (still is) the parentage tree. The scanner reads the transcripts, recovers the links and builds one tree per project, from the root conversation down to the last background shell task I wanted a layer built from real processes and real logs, that already exist if you are already using Claude Code and Codex, to see what's happening. But I didn't stop on just that The viewer can change nothing in your workflow. Orchestration still runs headless, you can never open the app, it reads the same tmux panes and logs that exist anyway. But when a subagent already did the work and I need one more fix, I can throw a message straight into its pane - so I'm able to reuse subagents freely I also have added Claude Code and Codex limits pane on a sidebar, voice transcription, Review cycle loop, and workflows (haven't tested properly the latest one, though). Will continue working on that. I vibecoded the whole thing with Claude Fable 5 and Codex, mostly to test the model. Glad I did, because earlier I tortured Codex with much simpler versions of a similar idea and it struggled a lot. Fable nailed it Possibly I created it because I have ADHD and holding ten parallel agents in my head is hard, but something tells me I am not the only one Try it, contribute: send 'bunx agent-log-viewer and set it up' to your agent. its for Linux, with limited macOS support, and no Windows support yet.

GitHub - Latand/live-log-viewer-next · GitHub
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
Latand
/
live-log-viewer-next
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
201 Commits 201 Commits .claude/ skills/ review-loop .claude/ skills/ review-loop .github/ workflows .github/ workflows bin bin docs docs public public scripts scripts src src .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md bun.lock bun.lock eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package.json package.json postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json View all files Repository files navigation
agent-log-viewer is a local web UI that turns raw Codex / Claude Code agent
logs into a readable, live-updating chat feed. It discovers every session,
subagent, Codex companion job and background shell task on your machine, links
them into a parent→child tree, and tails the selected one in real time.
Everything runs locally against files already on disk. There is no database and
no external service — the app reads ~/.claude and ~/.codex and renders what
it finds.
Claude Code sessions ( ~/.claude/projects/**/*.jsonl ) and their
subagents, rendered as a chat: user bubbles, assistant prose, tool-call
cards with ✓/✗ statuses and expandable output.
Codex CLI sessions ( ~/.codex/sessions/**/rollout-*.jsonl ) with command
cards, patches and service events.
Codex companion jobs ( ~/.claude/plugins/data/codex-openai-codex/state )
with a one-click jump to the full rollout session behind each job.
Background shell tasks ( claude-<uid>/**/tasks/*.output under the OS
temp dir — /tmp on Linux, $TMPDIR on macOS) — the originating Bash
command is recovered from the session transcript and shown above the
terminal output.
Parentage tree : session → subagents → codex jobs → rollouts → background
tasks, built server-side by scanning transcripts (append-only incremental,
cached — the warm /api/files poll stays around 100 ms).
Live activity : content-based badges — a transcript reads working while
it is mid-turn and done once the final assistant message lands.
Deep links : every selection is reflected in the URL ( #f=<path> ), so a
link opens that exact log.
Project scheme : each project is a pannable, zoomable diagram — the root
conversation on top, spawned agents one generation below, arrows colored by
engine. Quiet branches and tasks collapse under their nearest ancestor.
English or Ukrainian UI , model chips ( fable-5 , gpt-5.5 , sonnet …),
collapsible tree with persisted state, follow-mode autoscroll, service-event
toggle, and a line filter.
The session parentage tree, wiring root conversations to their spawned agents:
The package is published to npm, so the quickstart needs no clone:
bunx agent-log-viewer
# or
npx agent-log-viewer
This starts the server on 127.0.0.1:8898 and opens your browser.
bun install
bun run build
bun start --port 8898 --hostname 127.0.0.1
# open http://127.0.0.1:8898/
With npm (note the -- that forwards flags to the start script):
npm install
npm run build
npm start -- --port 8898 --hostname 127.0.0.1
start serves the output of the last build , so run build first. For
development, bun dev runs the app with hot reload (it needs a high OS
file-watch limit for large home directories).
Prerequisites: Node ≥ 20.9, and bun or npm/pnpm. tmux is optional — see
Platform support .
agent-log-viewer [options]
Option
Description
-p, --port <n>
Port for the local server (default 8898 ).
-H, --hostname <h>
Bind address (default 127.0.0.1 ).
--tailscale
Expose the viewer inside your tailnet (see below).
--new-token
Generate a fresh access key and invalidate old cookies.
--no-open
Don't open the browser on start.
-v, --version
Print the version.
-h, --help
Show usage.
Platform support
Linux is the native target: process discovery reads /proc directly. macOS is
supported through a portable backend that shells out to ps and lsof
instead — same live-process detection, tmux composer targeting, agent
spawn/kill and background-task discovery, just a bit more subprocess overhead
per scan. The backend is chosen automatically by process.platform (see
src/lib/proc/ ); VIEWER_PROC_BACKEND=portable forces the portable path on
Linux too, for testing.
The package supports Linux and macOS. The package's os field blocks Windows
installs with EBADPLATFORM . WSL works as Linux.
tmux is optional. Without it, log viewing, the parentage tree, live activity
and deep links all work; the composer, agent spawn/kill and resume-into-pane
features need tmux ( brew install tmux on macOS, or your distro's package on
Linux).
The UI defaults to English and shows a compact EN/UK toggle in the project
list header. The locale is resolved as localStorage key llv_lang first,
then the browser language (Ukrainian if the browser prefers it), then English.
CLI messages are English by default, and switch to Ukrainian with
LLV_LANG=uk or a uk_* value in LANG / LC_ALL .
Composers that talk to agents have a mic button for dictating messages. By
default transcription runs fully locally via faster-whisper — no audio leaves
the machine. Run scripts/setup-whisper.sh once to install the local engine.
Two cloud backends are available as an explicit per-machine opt-in (never a UI
toggle): ChatGPT (reuses your local Codex login) and ElevenLabs Scribe (the
only one with live, streaming transcription). Select with
LLV_TRANSCRIBE_BACKEND=local|chatgpt|elevenlabs or by writing the backend
name to ~/.config/agent-log-viewer/transcribe-backend ; local is the default.
See docs/transcription.md for setup, key locations,
and troubleshooting.
The viewer orchestrates implement→review cycles: a long-lived implementer
agent in tmux, a fresh read-only reviewer per round over the full diff,
automatic relay of findings, and a verdict deck in the scheme view. Start one
from the Flow chip above a conversation pane; presets pair engines and
reasoning efforts per role (e.g. Codex high → Fable ).
See docs/review-loop.md for the round protocol,
presets, the HTTP automation API, and troubleshooting. A Claude Code skill
for driving flows from an agent ships in .claude/skills/review-loop/ —
agents working in a clone pick it up automatically.
bunx agent-log-viewer --tailscale
--tailscale starts the local server on 127.0.0.1 and exposes it inside your
tailnet through a foreground tailscale serve <port> process. The public
internet (Funnel) is never used.
The CLI generates a 32-character access key, appends it to the tailnet URL as
?k=... , and after the first visit the server sets an llv_auth cookie for 30
days. --new-token generates a fresh key and immediately invalidates every old
cookie — each request compares the hash against the current token, so a cookie
minted with a previous key no longer passes.
The terminal prints the tailnet URL along with a QR code to scan with a phone.
The same QR is available inside the web UI: the QR-icon button in the project
list header opens a popover with the code and the link as text (with a copy
button). The QR is rendered entirely client-side (the qrcode package, no
external requests) and is served only to already-authorized clients — the same
token gate from src/proxy.ts also protects /api/access . When the server
runs without --tailscale , the button shows a hint to start
bunx agent-log-viewer --tailscale .
Anyone with tailnet access to this URL can read all agent transcripts,
including any sensitive data that landed in a session, and can execute commands
through /api/tmux and /api/spawn . Treat the tailnet URL as a secret — do
not forward it to anyone else.
The log APIs refuse any path that does not resolve into one of the whitelisted
log roots (see src/lib/scanner/roots.ts ). Mutating endpoints exist:
/api/tmux sends keys to tmux sessions and /api/spawn starts commands.
By default the CLI binds to 127.0.0.1 . With --tailscale , access is exposed
inside the tailnet via tailscale serve and guarded by the token gate in
src/proxy.ts . Non-loopback binds also force token mode. Treat any URL
containing ?k= as a credential.
All optional. Transcription variables are documented in full in
docs/transcription.md .
The viewer keeps its state under the standard XDG directories, named after the
package:
~/.config/agent-log-viewer/ — access token , transcribe-backend , and
elevenlabs-api-key .
~/.cache/agent-log-viewer/whisper-venv — the local transcription
virtualenv.
Legacy live-log-viewer paths are still honored as read-only fallbacks, so
existing setups keep working without changes.
See ARCHITECTURE.md : route handlers under src/app/api/* , a
pure scanner pipeline under src/lib/scanner/* (discover → describe → activity
→ model → links), React components under src/components/* . Caches live on
globalThis and survive dev hot-reload.
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
1
v0.8.0 — board fast path, mobile shell, copy affordances
Latest
Jul 6, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
