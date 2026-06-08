---
source: "https://sleuth-io.github.io/sx/2026/06/05/a-package-manager-for-ai-assets.html"
hn_url: "https://news.ycombinator.com/item?id=48448354"
title: "A package manager for AI assets (and why the lock file is per-user)"
article_title: "A package manager for AI assets (and why the lock file is per-user) | sx blog"
author: "detkin"
captured_at: "2026-06-08T18:57:54Z"
capture_tool: "hn-digest"
hn_id: 48448354
score: 3
comments: 0
posted_at: "2026-06-08T17:31:18Z"
tags:
  - hacker-news
  - translated
---

# A package manager for AI assets (and why the lock file is per-user)

- HN: [48448354](https://news.ycombinator.com/item?id=48448354)
- Source: [sleuth-io.github.io](https://sleuth-io.github.io/sx/2026/06/05/a-package-manager-for-ai-assets.html)
- Score: 3
- Comments: 0
- Posted: 2026-06-08T17:31:18Z

## Translation

タイトル: AI アセットのパッケージ マネージャー (およびロック ファイルがユーザーごとである理由)
記事のタイトル: AI アセットのパッケージ マネージャー (およびロック ファイルがユーザーごとである理由) | SXブログ
説明: sx が AI アセットの npm/Cargo マニフェストとロックの形状を借用する方法と、メンタル モデルが壊れる 2 つの場所: Git ID に対して解決されるユーザーごとのロック ファイルと、1 つのアセットを数十のクライアントの形式に正規化することです。

記事本文:
SXブログ
SXブログ
AI アセットのパッケージ マネージャー (およびロック ファイルがユーザーごとである理由)
過去 2 年間のある時点で、リポジトリは新しいカテゴリのファイルで静かにいっぱいになりました。コードではなく、正確な構成でもありません: プロンプト。ここに .claude/skills/ ディレクトリがあります。 .cursor/rules/ フォルダーがあります。ルートに CLAUDE.md、その隣に AGENTS.md、エージェントが呼び出しを許可されているサーバーをリストする .mcp.json があります。これらはコードベースでコーディング エージェントを役立つものであり、無数に存在します。
そのうちの 1 つがチームメイトに必要とされるほど優れていると判断した瞬間、あなたはコピー＆ペーストします。現在、コピーが 2 つあります。誰かが 1 つのバグを修正すると、コピーは漂流し、3 か月後にはどのバージョンが正規版なのか、どのリポジトリにそのバージョンがあるのか​​さえ誰もわかりません。 Git は各コピーをバージョン管理しますが、それは独自のリポジトリ内でのみです。あるリポジトリのコピーを別のリポジトリのコピーに結びつけるものは何もないので、共有される真実の情報源はなく、「プラットフォーム チームの全員がこのルールを理解するが、他の誰もが理解しない」と言う方法もありませんし、誰かが実際にあなたが書いたものを使用しているかどうかを知る方法もありません。
数週間前に私が話したあるチームは、まさにこれに当てはまる状況にあります。彼らは 5 つのリポジトリ (1 つのアプリケーションと 4 つのマイクロサービス) を実行しており、共有スキルに対するすべての修正を 5 つすべてにコピーする必要があります。彼らはそれを何か月も繰り返し、各コピーが着地するたびに少しずつ調整し、最終的にはコピーがわずかに異なる言葉でほぼ同じことを言っているほぼ重複したものになった。エンジニアの 1 人は、それらを「異なる時代の」バージョンと呼んでいました。最終的に、誰かが重複をインベントリするためだけに Confluence ページを作成しました。うち 4 人は数日ごとに集まり、誰がどのリポジトリを移行できたかを尋ねます。答えはいつも同じです。時間はありません。取引が完了したばかりで、今日出荷しなければならないものがあります。
これはツールが不足しているためではありません。クロード・コードは

プラグインとプラグイン マーケットプレイス: いくつかのスキル、コマンド、フックをバンドルして git リポジトリにプッシュすると、誰でもバンドルをインストールできます。本当に便利です。ただし、それは 1 つのクライアントであり、マーケットプレイスへのインストールは完全か無かです。それは他の誰のフォーマットではなく、Claude Code のフォーマットを作成します。また、「このルールはプラットフォーム チームのもので、これは全員のもの」という概念もありません。この問題のチーム間、クライアント間バージョンは依然として未解決のままです。
つまり、これは配布の問題であり、配布の問題には既知の形があり、パッケージ マネージャーを構築します。私たちはそうしました。それは SX と呼ばれます。そのほとんどは、マニフェスト、ロック ファイル、リゾルバーなど、誰もが予想するような退屈でよく理解された機構です。その部分は解決されており、sx は npm と Cargo がすでに決定している形状を借用するだけです。
標準のプレイブックでは、ほとんどのことは理解できますが、いくつかの間違った答えが返されます。AI アセットは、コード パッケージでは決して実行されない場所で、パッケージ マネージャーのメンタル モデルを破壊するためです。 package-lock.json で育ってきた人にとって、見出しの 1 つは異端です。ロック ファイルはコミットされておらず、チームの開発者ごとに異なります。地味なものは、誰も警告しない統合税です。十数の AI クライアントが 1 つの良いアイデアを採用し、それぞれが互換性のないオンディスク フォーマットを出荷しました。これら 2 つの問題がここでのスペースのほとんどを占めています。後半は短く、誰もコマンドを実行せずにアセットがどのようにインストールされたままになるか、ラップトップで実行されているボールトがどのようにして claude.ai などの Web クライアントにサービスを提供できるか、および設計が漏洩した場所についての正直な説明です。
sx は、npm、Cargo、uv がすべて独立して着地したマニフェストとロックの分割を借用しています。これは正しいためです。
人間が作成した信頼できる情報源であるマニフェスト sx.toml があります。すべての管理対象資産とそのバイトが存在する場所 (HTTP URL、git ref、lo

cal パス)、および誰がそれを取得する必要があります。それは金庫にコミットされています。トリミングされたアセット エントリは次のとおりです。
[[資産]]
名前 = "python-docstrings"
バージョン = "1.2.0"
type = "ルール"
[assets.source-http]
URL = "https://vault.example.com/assets/python-docstrings/1.2.0/bundle.zip"
ハッシュ = { sha256 = "e3b0c442…" }
[[アセット.スコープ]]
種類 = 「チーム」
チーム = 「プラットフォーム」
そして、コンテンツ ハッシュによって固定された、実際にインストールする完全に解決されたアーティファクトであるロック ファイルがあります。これまでのところ、何も異常はありません。そのマニフェストの興味深い単語はscopesであり、これはロックファイルがnpmのような動作を停止するものです。
ロック ファイルはあなたに対して解決されました
通常のパッケージ マネージャーでは、依存関係グラフはプロジェクトのプロパティです。リポジトリをチェックアウトする全員が同じグラフを解決し、同じ package-lock.json を取得します。ロック ファイルが共有されるのは、「このプロジェクトの依存関係は何ですか?」という質問に対する答えが 1 つだけであるためです。
質問 sx の答えは異なります。「このディレクトリにいるこの発信者は何をインストールする必要がありますか?」マニフェストは、アセットの範囲を組織、リポジトリ、リポジトリ内のパス、チーム、単一ユーザー、またはボット ID に設定します。それを解決するには、誰が尋ねているかを知る必要があります。そのため、解決には npm が必要としない 2 番目の入力、つまり呼び出し元の ID が必要となり、ロック ファイルは共有ファクトではなくマニフェストのユーザーごとの投影になります。
ID は git から来ます。 sx は、文書化されたフォールバック チェーンを使用して git config user.email にシェルを出力します。最初に SX_BOT 環境変数が優先され (CI ランナーとヘッドレス エージェントの場合)、次に明示的なオーバーライドが行われ、次に git config、そして最後の手段として合成 local:$USER@hostname が使用されます。その合成アイデンティティは意図的に二級市民です。読み取りには問題ありませんが、ミューテーションでは RequireRealIdentity が呼び出されます。

() を実行してバウンスするため、誤って監査エントリを local:nobody@laptop として書き込むことはありません。 ID の解決は git サブプロセスを意味するため、結果は (repoPath、SX_BOT) ペアごとにキャッシュされます。
リゾルバー、manifest.Resolve(manifest,actor) は、すべてのアセットのすべてのスコープを調べ、このアクターに対して、それが適用されるかどうか、および何に折りたたまれるかを決定します。注目に値するのはチームのケースです。
kind = "org" スコープはグローバルです。誰もがそれを理解しています。スコープ制限のないロック エントリに解決されます。
kind = "user" スコープは、電子メールが caller と一致する場合は global に解決され、そうでない場合は完全に削除されます。
興味深いのは、種類 = 「チーム」スコープです。チームは独自のリポジトリを持っています。したがって、チーム スコープは「チーム」に解決されません (ロック ファイルにはそのようなものはありません)。これは、呼び出し元がそのチームのメンバーである場合に限り、チームが所有するリポジトリごとに 1 つのリポジトリ スコープのエントリに展開されます。
最後のポイントは設計上の決定です。チームの名簿は常に変更されます。チームが所有するリポジトリの変更は少なくなります。チームが持つべき一連の資産は少なくとも変化します。 kind = Team をマニフェストに保存し、解決時のみ具体的なリポジトリに拡張することで、名簿の変更時に単一のアセット エントリを書き直す必要がなくなりました。メンバーシップは 1 か所 (チームの定義) に存在し、最後の瞬間にファンアウトします。
スコープが重複するため、拡張後にマージ パスが発生します。呼び出し元が同じリポジトリに対してリポジトリ全体のスコープと kind = "path" スコープ ( paths = ["docs/"] のようなものを運ぶパス スコープ) の両方を持つことになった場合、リポジトリ全体のエントリが優先され、狭い方のエントリが削除されます。リポジトリ全体のインストールでは、その下のすべてのパスがすでにカバーされているため、両方を保持すると、後で混乱する人のために冗長なエントリが残ることになります。パスは辞書順にソートされるため、出力は確定的になります。
次に、ロック f

ファイルはユーザーのキャッシュ ディレクトリに書き込まれ、ボールト URL のハッシュによってキー設定されます。そして、ここで私が気に入った最後のタッチが、回転が時間ではなくコンテンツに基づいて行われるということです。 sx は新しいロック ファイルを書き込むときにバイトをハッシュし、それらがすでにそこにあるものと一致する場合は、mtime に触れるだけです。それらが異なる場合は、新しいファイルがロードされる前に、古いファイルの名前が ISO-8601 タイムスタンプ (コロンはハイフンに置き換えられます、Windows のおかげです) で変更されます。したがって、lockfiles/<vault>.lock は常に最新であり、lockfiles/<vault>-2026-01-14T09-31-07Z.lock は、先週の火曜日にバイト単位でインストールしたものとまったく同じです。 1 か月間再解決されなかった古い CI ボックスには、実行内容の再現可能な記録がまだ残っています。セマンティックを変更せずにマニフェストを書き換えても、コンテンツ ハッシュは変更されないため、回転ノイズは発生しません。
npm が当然だと思っていることを放棄します。 2 人の開発者のインストールを git diff して、なぜ異なるのかを説明することはできません。ロックは最初から git に存在しなかったためです。それは共有される成果物ではないため、人々の間で比較するものは何もありません。返されるのは、もう一方の軸に沿った再現性です。回転された履歴は、時間の経過とともに 1 人の人物に限定された差分です。このドメインにとって、それは正しい取引です。 「どうしてアリスには私にはないルールがあるの？」答えは、一致するはずのない 2 つのロック ファイルの差分ではなく、マニフェスト内の彼女のチーム メンバーシップによって決まります。
すべてのクライアントが独自のファイル形式を発明しました
これは、解決するために最も多くのコードを必要とする問題であり、人生のほとんどの事柄に当てはまることですが、xkcd はそれをそう呼んでいました。
「ルール」は単純なアイデアです。いくつかの命令と、それらが適用されるファイル グロブのセットです。すべての AI クライアントはアイデアに同意し、それを異なる方法で実装しました。
Claude Code は、paths: キーを使用した YAML フロントマターを含む Markdown ファイルを必要とします。
カーソルは fro を含む .mdc ファイルを必要としています

ntmatter は globs: キーを使用し、さらに glob がない場合は alwaysApply: true フラグを使用します。 glob が 1 つある場合、glob 値は裸の文字列ですが、複数ある場合は YAML リストになります。
Gemini はルールごとのファイルをまったく実行しません。すべてはスコープごとに 1 つの GEMINI.md に格納されます。
そのため、正規のアセットは形式に依存せずに一度保存され、各クライアントはインストール時にそれをそのクライアントの方言にレンダリングするハンドラーを所有します。このアーキテクチャは、正しい意味で意図的に愚かです。資産タイプ A をクライアント B に変換する方法を認識する中央ルーターはありません。各クライアントは、1 つのインターフェイスを実装する密閉されたボックスです。
タイプ クライアント インターフェイス {
ID()文字列
SupportsAssetType ( t 資産 . Type ) bool
InstallAssets (ctx context . Context , req InstallRequest ) (InstallResponse , error )
// …
}
それぞれは、インポート時にプレーンな func init() を使用して自身を登録します。オーケストレーターは、インストールされているすべてのクライアントに同時にファンアウトし、サポートを宣言した資産のサブセットを各クライアントに渡し、ファイルシステムに対して必要なことを何でも実行できるようにします。クライアントの追加は、switch ステートメントの編集ではなく、パッケージの追加です。
機能宣言は、README のサポート マトリックスが実際に存在する場所です。クロード コードは、asset.AllTypes() を使用してそれ自体を構築します。 Gemini は、表現できる 5 つのタイプを正確にリストしており、エージェントやプラグインを黙って受け取ることはありません。マトリックスは、コードから逸脱する可能性のあるドキュメントではありません。それがコードです。
このレイヤーの 2 つの詳細は注目に値します。
ジェミニマーカーのトリック。 Gemini はすべてのルールを 1 つの共有 GEMINI.md にパックするため、sx は、ユーザーが手動で編集しているファイルを上書きすることなく、後で単一のルールを更新または削除する必要があります。各管理セクションを HTML コメント センチネルでラップします。
<!-- sx:python-docstrings -->
## Python ドキュメント文字列
ドキュメントストリングを適用する

すべての公的機能。
<!-- /sx:python-docstrings -->
更新はマーカー間の検索と置換です。アンインストールはそれらの間の削除です。ユーザーがマーカーの外側に書いたものは決して触れられません。これは、ドットファイル マネージャーが .bashrc の「マネージド ブロック」セクションに使用するのと同じ手法で、エージェントの指示に適用されます。
サポートされていないことは、失敗したことと同じではありません。フック イベントはクライアント間できれいにマッピングされません。クロード コードでは、ツール使用後とツール使用失敗後を区別します。 Gemini は両方を 1 つの AfterTool イベントにまとめており、 pre-compact の概念がまったくありません。ハンドラーがクライアントが表現できないイベントに遭遇すると、センチネル ErrUnsupportedEvent を返し、インストール パイプラインはその特定のエラーを StatusFailed ではなく StatusSkipped に変換します。
func TranslateInstallError (err error , successMessage string ) (ResultStatus , string , error ) {
if err == nil {
StatusSuccess 、 successMessage 、 nil を返す
}
エラーが発生した場合。 Is (err, フック .ErrUnsupportedEvent) {
StatusSkipped を返します。エラー。エラー ()、エラー
}
StatusFailed 、 fmt を返します。 Sprintf ( "インストールに失敗しました: %v" , err ), err
}
したがって、6 つのイベントのフック バンドルを Gemini にインストールすると、失敗するのではなく、「インストールされました、2 つのイベントがスキップされました」と報告されます。その違いが重要なのです。部分的なインストールごとにハード障害を報告するツールは無視され、それを報告するツールは無視されます。

[切り捨てられた]

## Original Extract

How sx borrows the npm/Cargo manifest-and-lock shape for AI assets — and the two places that mental model breaks: a per-user lock file resolved against your git identity, and normalizing one asset into a dozen clients’ formats.

sx blog
sx blog
A package manager for AI assets (and why the lock file is per-user)
Sometime in the last two years your repos quietly filled up with a new category of file. Not code, not config exactly: prompts. A .claude/skills/ directory here. A .cursor/rules/ folder there. A CLAUDE.md at the root, an AGENTS.md next to it, a .mcp.json listing the servers your agent is allowed to call. These are the things that make a coding agent useful on your codebase, and they’re sprawling.
The moment one of them is good enough that a teammate wants it, you copy-paste. Now there are two copies. Someone fixes a bug in one, the copies drift, and three months later nobody can tell which version is canonical or which repos even have it. Git is versioning each copy, but only inside its own repo. Nothing connects the copy in one repo to the copy in another, so there’s no shared source of truth, no way to say “everyone on the platform team gets this rule but nobody else,” and no way to know if anyone is actually using the thing you wrote.
A team I talked to a few weeks ago lives exactly this. They run five repos: one application and four microservices, and every fix to a shared skill has to be copied into all five. They’d been doing that for months, tweaking each copy a little as it landed, until the copies had drifted into near-duplicates that say roughly the same thing in slightly different words. One of their engineers called them versions “from different eras.” Someone eventually built a Confluence page just to inventory the duplication. Four of them meet every couple of days to ask who managed to migrate which repo, and the answer is always the same: no time, a deal just closed, something has to ship today.
This isn’t for lack of tooling. Claude Code has plugins and a plugin marketplace: bundle up some skills, commands, and hooks, push them to a git repo, and anyone can install the bundle. It’s genuinely useful. But it’s one client, and a marketplace install is all-or-nothing. It writes Claude Code’s formats and no one else’s, and it has no concept of “this rule for the platform team, that one for everybody.” The cross-team, cross-client version of the problem is still wide open.
So it’s a distribution problem, and distribution problems have a known shape: you build a package manager. We did; it’s called sx . Most of it is the boring, well-understood machinery you’d expect: a manifest, a lock file, a resolver. That part is solved, and sx just borrows the shape npm and Cargo already settled on.
The standard playbook gets you most of the way and then gives a few wrong answers, because AI assets break the package-manager mental model in places code packages never do. The headline one is heretical if you grew up on package-lock.json : the lock file is not committed, and it is different for every developer on the team. The unglamorous one is the integration tax nobody warns you about. A dozen AI clients took one good idea and each shipped an incompatible on-disk format for it. Those two problems get most of the space here. The back half is shorter: how the assets stay installed without anyone running a command, how a vault running on your laptop can serve the web clients like claude.ai, and an honest accounting of where the design leaks.
sx borrows the manifest-and-lock split that npm, Cargo, and uv all landed on independently, because it’s correct.
There’s a manifest, sx.toml , which is the human-authored source of truth. It lists every managed asset, where its bytes live (an HTTP URL, a git ref, a local path), and who should get it . It’s committed to the vault. Here’s a trimmed asset entry:
[[assets]]
name = "python-docstrings"
version = "1.2.0"
type = "rule"
[assets.source-http]
url = "https://vault.example.com/assets/python-docstrings/1.2.0/bundle.zip"
hashes = { sha256 = "e3b0c442…" }
[[assets.scopes]]
kind = "team"
team = "platform"
And there’s a lock file, the fully-resolved artifact you actually install from, pinned by content hash. So far, nothing unusual. The interesting word in that manifest is scopes , and it’s what makes the lock file stop behaving like npm’s.
The lock file is resolved against you
In a normal package manager the dependency graph is a property of the project. Everyone who checks out the repo resolves the same graph and gets the same package-lock.json . The lock file is shared precisely because the question it answers, “what are this project’s dependencies?”, has exactly one answer.
The question sx answers is different: “what should this caller, standing in this directory have installed?” The manifest scopes assets to an org, a repo, a path within a repo, a team, a single user, or a bot identity. Resolving that requires knowing who’s asking. So resolution takes a second input that npm never needs, the caller’s identity, and the lock file becomes a per-user projection of the manifest rather than a shared fact.
Identity comes from git. sx shells out to git config user.email , with a documented fallback chain: an SX_BOT environment variable wins first (for CI runners and headless agents), then explicit overrides, then the git config, then a synthetic local:$USER@hostname as a last resort. That synthetic identity is deliberately a second-class citizen. It’s fine for reads, but any mutation calls RequireRealIdentity() and bounces it, so nobody accidentally writes audit entries as local:nobody@laptop . Resolving identity means a git subprocess, so the result is cached per (repoPath, SX_BOT) pair.
The resolver, manifest.Resolve(manifest, actor) , walks every scope on every asset and decides, for this actor, whether it applies and what it collapses to. The team case is the one worth looking at:
A kind = "org" scope is global. Everyone gets it; it resolves to a lock entry with no scope restriction.
A kind = "user" scope resolves to global if the email matches the caller , and is dropped entirely otherwise.
A kind = "team" scope is the interesting one. Teams own repositories. So a team scope doesn’t resolve to “team” (there’s no such thing in a lock file). It expands into one repo-scoped entry per repository the team owns , but only if the caller is a member of that team.
That last point is the design decision. Team rosters change constantly; repositories a team owns change less; the set of assets a team should have changes least of all. By storing kind = team in the manifest and expanding it to concrete repos only at resolution time, a roster change never requires rewriting a single asset entry. The membership lives in one place (the team definition) and fans out at the last possible moment.
There’s a merge pass after expansion, because scopes overlap. If a caller ends up with both a whole-repo scope and a kind = "path" scope for the same repo (the path scope carrying something like paths = ["docs/"] ), the repo-wide entry wins and the narrower one is dropped. A repo-wide install already covers every path under it, so keeping both would leave a redundant entry for someone to puzzle over later. Paths are sorted lexicographically so the output is deterministic.
Then the lock file is written to the user’s cache directory, keyed by a hash of the vault URL, and here’s the last touch I liked: rotation is keyed on content, not time. When sx writes a new lock file it hashes the bytes, and if they match what’s already there it just touches the mtime. If they differ, the old file is renamed with an ISO-8601 timestamp (colons swapped for hyphens, thanks Windows) before the new one lands. So lockfiles/<vault>.lock is always current, and lockfiles/<vault>-2026-01-14T09-31-07Z.lock is the exact thing you had installed last Tuesday, byte for byte. A stale CI box that hasn’t re-resolved in a month still has a reproducible record of what it’s running. Rewriting the manifest with no semantic change produces no rotation noise, because the content hash is unchanged.
You do give up something npm takes for granted. You can’t git diff two developers’ installs to explain why they differ, because the lock was never in git to begin with. It isn’t a shared artifact, so there’s nothing to compare across people. What you get back is reproducibility along the other axis: the rotated history is the diff, scoped to one person over time. For this domain that’s the right trade. “Why does Alice have a rule I don’t?” is answered by her team membership in the manifest, not by diffing two lock files that were never meant to match.
Every client invented its own file format
Here’s the problem that requires the most code to solve and, as is true with most things in life, xkcd called it.
A “rule” is a simple idea: some instructions plus a set of file globs they apply to. Every AI client agreed on the idea and then implemented it differently.
Claude Code wants a Markdown file with YAML frontmatter using a paths: key.
Cursor wants a .mdc file with frontmatter using a globs: key, plus an alwaysApply: true flag when there are no globs. The glob value is a bare string when there’s one glob but a YAML list when there are several.
Gemini doesn’t do per-rule files at all. Everything goes into one GEMINI.md per scope.
So the canonical asset is stored once, format-neutral, and each client owns a handler that renders it into that client’s dialect on install. The architecture is deliberately dumb in the right way: there’s no central router that knows how to translate asset type A for client B. Each client is a sealed box implementing one interface:
type Client interface {
ID () string
SupportsAssetType ( t asset . Type ) bool
InstallAssets ( ctx context . Context , req InstallRequest ) ( InstallResponse , error )
// …
}
Each one registers itself at import time with a plain func init() . The orchestrator fans out across every installed client concurrently, hands each one the subset of assets it declared support for, and lets it do whatever it wants to the filesystem. Adding a client is adding a package, not editing a switch statement.
The capability declaration is where the README’s support matrix actually lives. Claude Code constructs itself with asset.AllTypes() ; Gemini lists exactly the five types it can represent and silently never receives an agent or a plugin. The matrix isn’t documentation that can drift from the code. It is the code.
Two details from this layer are worth calling out.
The Gemini marker trick. Since Gemini packs every rule into one shared GEMINI.md , sx needs to update or remove a single rule later without clobbering a file the user might also be hand-editing. It wraps each managed section in HTML comment sentinels:
<!-- sx:python-docstrings -->
## Python docstrings
Enforce docstrings on all public functions.
<!-- /sx:python-docstrings -->
Update is a find-and-replace between the markers; uninstall is a delete between them; anything the user wrote outside the markers is never touched. It’s the same technique your dotfile manager uses for “managed block” sections of .bashrc , applied to agent instructions.
Unsupported is not the same as failed. Hook events don’t map cleanly across clients. Claude Code distinguishes post-tool-use from post-tool-use-failure ; Gemini collapses both into one AfterTool event and has no concept at all of pre-compact . When a handler hits an event the client can’t express, it returns a sentinel ErrUnsupportedEvent , and the install pipeline translates that specific error into a StatusSkipped rather than a StatusFailed :
func TranslateInstallError ( err error , successMessage string ) ( ResultStatus , string , error ) {
if err == nil {
return StatusSuccess , successMessage , nil
}
if errors . Is ( err , hook . ErrUnsupportedEvent ) {
return StatusSkipped , err . Error (), err
}
return StatusFailed , fmt . Sprintf ( "Installation failed: %v" , err ), err
}
So installing a six-event hook bundle onto Gemini reports “installed, two events skipped” instead of failing. That difference matters. A tool that reports a hard failure on every partial install gets ignored, and a tool tha

[truncated]
