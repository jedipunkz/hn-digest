---
source: "https://www.salvager.sh/"
hn_url: "https://news.ycombinator.com/item?id=48578204"
title: "Salvager, Undo for AI agents that edit your files"
article_title: "Salvager — Your agent broke it. Get it back."
author: "gorkamb"
captured_at: "2026-06-18T01:02:42Z"
capture_tool: "hn-digest"
hn_id: 48578204
score: 1
comments: 0
posted_at: "2026-06-17T23:00:45Z"
tags:
  - hacker-news
  - translated
---

# Salvager, Undo for AI agents that edit your files

- HN: [48578204](https://news.ycombinator.com/item?id=48578204)
- Source: [www.salvager.sh](https://www.salvager.sh/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T23:00:45Z

## Translation

タイトル: Salvager、ファイルを編集する AI エージェントを元に戻す
記事のタイトル: Salvager — あなたのエージェントはそれを破りました。取り戻してください。
説明: AI コーディング エージェントのためのファイルシステム セーフティ ネット。パッシブ ウォッチャーはすべてのファイル リビジョンを自動的に保存するため、エージェントが破損したものはすべて取り戻すことができます。

記事本文:
Salvager — あなたのエージェントがそれを破りました。取り戻してください。
コンテンツにスキップ
サルベージャー 仕組み エージェント向け クイックスタート FAQ エンタープライズ GitHub
あなたのエージェントがそれを破りました。
取り戻してください。
Salvager は、AI コーディング エージェントのためのファイルシステム レベルのセーフティ ネットです。パッシブ
ウォッチャーは、あなたとエージェントの作業中にすべてのファイル リビジョンを保存します。
エージェントが中断しても、戻ることができます。
GitHub で見る
仕組み
設定ゼロ。任意のプロジェクト ルートで実行します。
コピーサルベージャーウォッチの問題
エージェントは、差分を読み取るよりも速くファイルを書き換えることができます。
AI コーディング エージェントはディスク上の実際のファイルを編集します。ほとんどの場合、それは
まさにあなたが望むもの。しかし、間違いは沈黙します。破壊された関数、
削除されたブロック、行き過ぎたリファクタリング、そして気づいたときには、
以前のバージョンはすでになくなっている可能性があります。
回復はエージェントが注意するかどうかに依存するべきではなく、あなたが
適切なタイミングでコミットされたか、たまたま開いていたエディタ上でコミットされました。
エージェントの下、ファイルシステムに存在する保証が必要です。
受動的に監視するものであり、覚えておく必要があるワークフローではありません。
AI エージェントのローカル ヒストリーを考える: キャプチャは自動、リカバリは自動
CLI、MCP、素手など、何でも構いません。二人は
意図的に分離されているため、あなたを保護する部分は決して依存しません。
失敗する可能性がある部分。
任意のプロジェクト ルートで Salvager watch を実行します。設定もアカウントもクラウドもありません。早速木の観察を始めます。
ユーザーのキーストロークであれ、エージェントの編集であれ、ファイルが変更されるたびに、Salvager はファイルごとの新しいリビジョンを .salvager/ に保存します。自動的に。継続的に。
CLI、MCP 経由、または単純な ls と cat を使用して手動でファイルのタイムラインを参照します。すべてのバージョンが存在し、それらを区別するためのコンテンツ信号が付いています。
ファイルを以前のリビジョンにロールバックします。復元自体はまず現在の状態を保存します

, リバーシブルでもあります。何も行き止まりになることはありません。
あなたのセーフティネットが捕まえるために構築されたものではありません。
エージェントが実際の作業を行って破壊し、その git (ファイル システム) を見つけるのを見てください。
そして編集者は何も返すことができません。それからサルベージャーがそれを手渡すのを見て、
無傷。ケースを選択してください。
git replace --hard で破棄されたコミットされていない作業を復元します。
インプレース上書き後にほとんどの行が削除された後、追跡されていないデータセットを回復します。
エージェントが弱い散文で書き直したドキュメントセクションを復元する。
後のパスでビルドが中断された後、複数ファイルのリファクタリングの最後に正常に実行されたパスを回復すると。
エージェントが自ら実行できる回復レイヤー。
サルベージャーはMCPに関する歴史を暴露する。何かを壊すエージェントは、
タイムラインを発見し、過去のリビジョンを検査し、適切なリビジョンを復元します。
ループ内に人間はいません。まず読んでから復元してください。
ファイルの保存されたリビジョンを最新のものから順にリストします。それぞれがそのコンテンツ信号 (行、行デルタ、開始署名) を伝送するため、エージェントはすべてのバージョンを再読することなく、正しいものを選択できます。
1 つのリビジョンの正確な内容を取得します。読み取り専用検査。エージェントは何かを変更する前に、必要なバージョンが見つかったことを確認します。
ファイルを選択したリビジョンに復元します。まず、現在の状態を復元前のリビジョンとして保存するため、復元自体は元に戻すことができます。
コピー {
"mcpサーバー": {
「サルベージャー」: {
"コマンド": "サルベージャー",
"args": ["mcp"],
"cwd": "/パス/への/プロジェクト"
}
}
}
MCP を介した削除は不可
MCP 経由で公開されるパージまたは削除ツールはありません。セーフティネット
壊れそうなものでは消せない。エージェントができることは、
作品は復元できますが、その記録は絶対に破棄しないでください。
保証はその構造により有効です。
全てを捉える→全てが可逆的。デザインは、
セーフティ ネットが通常失敗する方法: 構成の欠落、部分的なカバレッジ、
破壊的回復、

ロックイン。
任意のプロジェクトのルートでサルベージャーを監視します。設定ファイルもアカウントもセットアップの儀式も必要ありません。それはただあなたを守り始めるだけです。
cgo、ランタイム、依存関係を持たない 1 つの静的 Go バイナリ。それを放り込んで実行してください。 Linux と macOS、amd64 と arm64。
履歴は .salvager/ の下にあるプレーン ファイルです。バイナリを紛失しましたか?裸の犬と猫は何でも手で回収します。ロックインや独自のフォーマットはありません。
OS のリアルタイム監視の上限に達すると、オーバーフロー サブツリーは自動的にポーリング スイープに戻ります。カバー範囲は完全なままであり、黙って部分的なものになることはありません。
すべての復元では、まず現在の状態が復元前のリビジョンとして保存されます。復元自体は元に戻すことができます。一方通行のドアはありません。
MCP 経由ではパージや削除は公開されません。コードを破壊する可能性のあるエージェントは、コードを回復できるようにするレコードを消去できません。
すべてがマシン上に残ります。テレメトリー、アップロード、サードパーティはありません。 Apache-2.0 ライセンスを取得しており、完全に自己ホスト可能です。
既知のギャップを使用して実行することは可能ですが、それは明示的なオプトインとしてのみ可能です。
--allow-partial 。背中の後ろでカバー範囲が減少することはありません。
1 分以内に起動して実行できます。
インストール スクリプトは、単一の静的バイナリを PATH にドロップします。それ
sudo を使用せず、テレメトリを送信せず、シェルには触れません
構成。 macOS と Linux。
コピーcurl -fsSL https://raw.githubusercontent.com/usesalvager/salvager/main/install.sh |しー
Homebrew (macOS/Linux) の場合:
brew install usesalvager/tap/salvager をコピーします
自分で構築したいですか? Go ツールチェーンを使用すると、
build -o salvager に進みます。からの
クローン。 Linux および macOS (amd64 + arm64) 用のビルド済みバイナリ。それぞれ
SHA-256 チェックサムは、
リリースページ。 Windows はソースからのベストエフォート型です。
開始するコマンドは 1 つです。そこからのすべてはあなたの条件に応じて回復されます。
Copy # ワンコマンドオンボーディング: MCP サーバーを登録し、エージェント構成を更新します
サルベージャー・イニ

t
# 現在のプロジェクトの監視を開始します (キャプチャはすぐに開始されます)
サルベージャーウォッチ
# ファイルの記録されたバージョンをリストする
サルベージャー履歴 config.json
# 1 つのバージョンのコンテンツを検査します (履歴からのタイムスタンプ)
サルベージャー ショー config.json 1718312445
# ファイルをあるバージョンに復元します (それ自体は元に戻すことができます)
サルベージャー復元 config.json 1718312445
MCP 経由でエージェントに接続するには、次を参照してください。
上記のエージェント向け。
編集が行われると、ウォッチャーはすでにファイルのリビジョンを自動的に保存しています。バージョンを一覧表示し、適切なバージョンを復元すれば、元に戻ります。復元自体はまず現在の状態を保存するため、それさえも元に戻すことができます。
コミットもステージングも必要なく、覚えておくべきことは何もありません。 Salvager は、コミットしていない作業も含め、すべての保存をバックグラウンドでキャプチャし、git を置き換えるのではなく補完します。 git は永続的な共有履歴です。 Salvager は、その下にある常時稼働のローカル セーフティ ネットです。
いいえ。任意のプロジェクト ルートでサルベージャー ウォッチを実行します。ゼロ構成: アカウント、クラウド、構成ファイルはありません。起動時にすべての追跡ファイルの初期リビジョンを記録し、その後のすべての変更をキャプチャします。
プロジェクト内のプレーンな .salvager/ ディレクトリ内。オブジェクトはコンテンツ ハッシュによって重複排除され、リビジョンは単純なログにリストされるため、裸の ls と cat は手動ですべてを回復します。独自のフォーマットやロックインはありません。
いいえ。MCP サーバーは、バージョンのリスト、バージョンの取得、および復元という 3 つのツールを公開します。 MCP にはパージや削除がないため、機能を壊す可能性のあるエージェントによってセーフティ ネットを消去することはできません。
ランタイムも依存関係も持た​​ない単一の静的 Go バイナリ。 macOS および Linux (amd64 および arm64)、ソースから Windows のベストエフォートを使用。 Apache-2.0 ライセンスを取得しました。

## Original Extract

A filesystem safety net for AI coding agents. A passive watcher saves every file revision automatically, so anything an agent breaks, you can get back.

Salvager — Your agent broke it. Get it back.
Skip to content
salvager How it works For agents Quickstart FAQ Enterprise GitHub
Your agent broke it.
Get it back.
Salvager is a filesystem-level safety net for AI coding agents. A passive
watcher saves every file revision as you and your agent work, so anything
the agent breaks, you can get back.
View on GitHub
How it works
Zero config. Run it in any project root:
Copy salvager watch The problem
An agent can rewrite a file faster than you can read the diff.
AI coding agents edit real files on disk. Most of the time that's
exactly what you want. But mistakes are silent. A clobbered function, a
deleted block, a refactor that ran too far, and by the time you notice,
the previous version may already be gone.
Recovery shouldn't depend on the agent being careful, on you having
committed at the right moment, or on an editor that happened to be open.
You need a guarantee that lives below the agent, down at the filesystem.
A passive watcher, not a workflow you have to remember.
Think Local History for AI agents: capture is automatic, recovery is
whatever you want, whether CLI, MCP, or your bare hands. The two are
decoupled on purpose, so the part that protects you never depends on the
part that might fail.
Run salvager watch in any project root. No config, no accounts, no cloud. It starts observing the tree immediately.
Each time a file changes, whether your keystrokes or an agent's edit, Salvager saves a new per-file revision into .salvager/. Automatically. Continuously.
Browse the timeline for any file from the CLI, over MCP, or by hand with plain ls and cat. Every version is there, with a content signal to tell them apart.
Roll any file back to any earlier revision. The restore itself first saves the current state, so it's reversible too. Nothing is ever a dead end.
The work your safety net was never built to catch.
Watch an agent do real work, destroy it, and find that git, the filesystem,
and the editor have nothing to give back. Then watch Salvager hand it over,
unharmed. Pick a case.
Recovering uncommitted work that git reset --hard destroyed.
Recovering an untracked dataset after an in-place overwrite dropped most rows.
Restoring a documentation section an agent rewrote with weaker prose.
Recovering the last-good pass of a multi-file refactor after a later pass broke the build.
A recovery layer your agent can drive itself.
Salvager exposes the history over MCP. An agent that breaks something can
discover the timeline, inspect past revisions, and restore the right one,
no human in the loop. Read first, then restore.
List the saved revisions for a file, newest first. Each carries its content signal (lines, line delta, start signature) so the agent can pick the right one without re-reading every version.
Fetch the exact contents of one revision. Read-only inspection, so the agent confirms it found the version it wants before changing anything.
Restore a file to a chosen revision. First saves the current state as a pre-restore revision, so the restore is itself undoable.
Copy {
"mcpServers": {
"salvager": {
"command": "salvager",
"args": ["mcp"],
"cwd": "/path/to/project"
}
}
}
No delete over MCP
There is no purge or delete tool exposed over MCP. The safety net
can't be erased by the thing that might break it. An agent can
recover work, but never destroy the record of it.
The guarantee holds because of how it's built.
Captures everything → everything is reversible. The design removes the
ways a safety net usually fails: missing config, partial coverage,
destructive recovery, lock-in.
salvager watch in any project root. No config files, no accounts, no setup ritual. It just starts protecting you.
One static Go binary with no cgo, no runtime, no dependencies. Drop it in and run. Linux and macOS, amd64 and arm64.
History is plain files under .salvager/. Lost the binary? A bare ls and cat recovers anything by hand. No lock-in, no proprietary format.
When the OS real-time watch ceiling is hit, overflow subtrees fall back to a polling sweep automatically. Coverage stays complete, never silently partial.
Every restore first saves the current state as a pre-restore revision. Restoring is itself undoable. There are no one-way doors.
No purge or delete is exposed over MCP. The agent that might break your code cannot erase the record that lets you recover it.
Everything stays on your machine. No telemetry, no upload, no third party. Apache-2.0 licensed and fully self-hostable.
Running with known gaps is possible, but only as an explicit opt-in via
--allow-partial . Coverage is never reduced behind your back.
Up and running in under a minute.
The install script drops a single static binary on your PATH. It
never uses sudo, sends no telemetry, and doesn't touch your shell
config. macOS and Linux.
Copy curl -fsSL https://raw.githubusercontent.com/usesalvager/salvager/main/install.sh | sh
On Homebrew (macOS/Linux):
Copy brew install usesalvager/tap/salvager
Prefer to build it yourself? With a Go toolchain,
go build -o salvager . from a
clone . Prebuilt binaries for Linux and macOS (amd64 + arm64), each with a
SHA-256 checksum, are on the
releases page . Windows is best-effort from source.
One command to start. Everything from there is recovery on your terms.
Copy # One-command onboarding: registers the MCP server and updates your agent config
salvager init
# Start watching the current project (capture begins immediately)
salvager watch
# List the recorded versions of a file
salvager history config.json
# Inspect one version's content (timestamp from history)
salvager show config.json 1718312445
# Restore a file to a version (itself reversible)
salvager restore config.json 1718312445
To wire it into an agent over MCP, see
For agents above.
The watcher has already saved a revision of the file, automatically, as the edit happened. List the versions, restore the good one, and you're back. The restore itself first saves the current state, so even that is reversible.
No commits, no staging, nothing to remember. Salvager captures every save in the background, including work you never committed, and complements git rather than replacing it. git is your durable, shared history; Salvager is the always-on local safety net underneath it.
No. Run salvager watch in any project root. Zero configuration: no accounts, no cloud, no config files. It records an initial revision of every tracked file on startup, then captures every change thereafter.
In a plain .salvager/ directory inside your project. Objects are deduplicated by content hash and revisions are listed in plain logs, so a bare ls and cat recover anything by hand. No proprietary format, no lock-in.
No. The MCP server exposes exactly three tools: list versions, get a version, and restore. There is no purge or delete over MCP, so the safety net can't be erased by the agent that might break things.
A single static Go binary with no runtime and no dependencies. macOS and Linux (amd64 and arm64), with Windows best-effort from source. Apache-2.0 licensed.
