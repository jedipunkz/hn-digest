---
source: "https://jonnyzzz.com/blog/2026/06/20/auditing-git-for-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48643981"
title: "What Did the Agent Just Push? Auditing Git for AI Agents"
article_title: "What Did the Agent Just Push? Auditing Git for AI Agents – Eugene Petrenko"
author: "speckx"
captured_at: "2026-06-23T12:53:20Z"
capture_tool: "hn-digest"
hn_id: 48643981
score: 2
comments: 0
posted_at: "2026-06-23T12:30:53Z"
tags:
  - hacker-news
  - translated
---

# What Did the Agent Just Push? Auditing Git for AI Agents

- HN: [48643981](https://news.ycombinator.com/item?id=48643981)
- Source: [jonnyzzz.com](https://jonnyzzz.com/blog/2026/06/20/auditing-git-for-ai-agents/)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T12:30:53Z

## Translation

タイトル: エージェントは何をプッシュしましたか? AI エージェントの Git の監査
記事のタイトル: エージェントは何をプッシュしましたか? AI エージェントの Git の監査 – Eugene Petrenko
説明: リポジトリを自律型 AI エージェントに渡すときは、それが触れることができるものを制限し、それが行ったことをキャプチャする必要があります。パート 1 — アプローチ — キーをエージェントから遠ざけ、Git プロキシをハード境界として、Git フックをより軽量な代替手段として使用します。

記事本文:
創設エンジニアリングリーダー | Agentic AI 開発ツールとエクスペリエンス
エージェントは何をプッシュしたのでしょうか? AI エージェントの Git の監査
ポスト
ポスト
リポジトリを自律型 AI エージェントに渡して機能させます。クローン、ブランチ、コミット、プッシュ、
おそらくプルリクエストをオープンします。数分後、完了したことが表示されます。 2 つの質問が保留されたままです。
そしてどちらも納得のいく答えを持っていません。
実際にどのコミットが生成されたのでしょうか?それはあなたのために書いた要約ではなく、実際のオブジェクトです。
サーバーに着陸しました。
許可したブランチのみに触れたのでしょうか?混乱しているエージェントの強制プッシュ メインの 1 つは、
悪い午後。
私はしばらくの間、同じコードベースでエージェントのフリートを実行してきました (「
AI フリートと
Git Fork パターン )、実行するエージェントが増えるほど、
さらに、これら 2 つの質問は好奇心から運用リスクに変わります。この投稿は、次の場所でそれらに答えることについてのものです
真実が存在するレイヤー: git 自体。
この投稿はそのアプローチです: なぜこれが git の問題なのか、他のすべてを可能にする 1 つのルール、そして
git プロキシと git フックという 2 つのメカニズムを、推論して実行できるレベルで実現します。最も深い
プロキシのワイヤ プロトコル メカニズムは、独自のフォローアップを取得できるほど関与しています。ここでカバーします
AI エージェントが実装を開始できる十分な疑似コードを含む、形状と決定。
積み重ねる。
これがポリシー文書ではなく Git の問題である理由
素朴な答えは、「エージェントがどのブランチに触れる可能性があるかを伝えるだけです。」です。それがポリシーです。ポリシーは
シェル、ネットワーク、およびタスクを完了したいという強い衝動を持つプロセスへの提案。エージェントは
賢い — 要求として表現された制約を与えると、制約から抜け出すのではなく、それを回避する創造的な方法を見つけるかもしれません。
悪意はありますが、障害物を回避するのは文字通り

私たちがそれを行うために訓練したこと。
したがって、私たちは上手に尋ねることに依存しません。次の 2 つのプロパティに依存します。
爆発半径 — 選択したヘッドのみが実サーバー上で変更でき、どこかで強制されます。
エージェントは迂回できません。
来歴 — エージェントが作成したすべてのコミットがキャプチャされて属性が付けられるため、復元、レビュー、
そして事後に監査するのです。これは、後で反復学習を行うための優れた情報源になります。
そこに行くには 2 つの方法があります。 1 つ目は、エージェントのサンドボックスの外側にある git プロキシです。
厳しい境界線。 2 つ目は、サンドボックス イメージに焼き付けられた git フックです。これは、軽量でベストエフォート型のレイヤーです。
イメージは所有しているが、ネットワーク パスは所有していない場合。プロキシを実際の壁として使用し、フックを次のように扱います。
追加のテレメトリ。その理由は次のとおりです。
すべての根底にある 1 つのルール: エージェントに本物のキーを渡さない
プロキシやフックの前に、残りを可能にする決定があります。
エージェントは、実際の git サーバーへの資格情報を保持することはありません。
サンドボックス内で、エージェントは有効期間が短く、エンドポイントに対して認証されるスコープ付きトークンを取得します。
GitHub や Space、内部 git ホストではありません。そのトークンを Git 資格情報ヘルパーという退屈な方法で注入します。
または URL の書き換えであり、ディスク上の実際のキーは決してありません。実際のプッシュ認証情報は、エージェントがプロキシ内に存在します。
読めません。プッシュが許可されている場合、プロキシは独自のサービス資格情報を使用してアップストリームを認証します。
これは、賢いエージェントが生き残る特性です。サンドボックスはアクセス許可の境界であり、サンドボックス自体の内部にあります。
サンドボックスエージェントは、環境の読み取り、ファイルの書き換え、
インストールしたツールを無効にします。本物の鍵がその環境のどこかにあった場合、決意を持ったエージェントが
それを見つけて Git サーバーと直接通信できれば、私たちが構築したすべてのコントロールは装飾になります。ベック

を使う
キーが存在しない場合、逃走したエージェントが到達できる最悪の事態はプロキシです。プロキシはポリシーを強制し、
鍵を握っている。影響範囲は、認証情報が存在する場所によって制限され、認証情報がどのように適切に動作するかによって制限されます。
エージェントは
注意したいのは、エージェントのプッシュ/プル操作を制限するパターンは標準ではないということです。
また、GitHub、Space、またはサードパーティの Git ホストを使用してそのまま実装することは、困難または不可能である可能性があります。
Agentic AI は、私たちにさらに別の方向性をもたらします。ソフトウェアは、方法を知っていれば、非常に安価に構築できるようになりました。
以下では、それを実装するための十分なアイデアを説明します。 Git は、HTTP、SSH、Git over TCP などの複数のプロトコルをサポートします。
以下では HTTP を使用することにしましたが、同じ手法は他のプロトコル、主に SSH にも適用できます。
解決策 1: Git プロキシ — 厳密な境界
HTTP 経由の git プッシュは、驚くほど検査可能な会話です。間のパスにプロキシを置くと、
エージェントとサーバーを介して、交換全体を読み取り、許可するかどうかを決定し、すべてをキャプチャできます。
それは、エージェントが違いを見分けることができず、通過することもできずに流れます。
あなたの周りには他に行くところがないから
実際にワイヤー上でプッシュするとどのように見えるか
Git の「スマート HTTP」プロトコルは 2 つの HTTP リクエストでプッシュを実行します
( gitprotocol-http ):
Ref アドバタイズメント — GET /info/refs?service=git-receive-pack 。 #service=git-receive-pack の後
ヘッダー行とフラッシュにより、サーバーはすべての参照とその現在のオブジェクト ID をリストし、最初の参照についてもリストします。
行に NUL バイトを追加し、その後に
サポートされる機能。
プッシュ自体 — POST /git-receive-pack 。本体は一連のコマンド (「update」) です。
refs/heads/x from <old> to <new> ”) の後に PACK (オブジェクトのバイナリ BLOB) が続きます。
転送されました。
Th

アドバタイズメント、コマンド リスト、およびステータス レポートはすべて pkt 行 (4 バイトの 16 進数) でフレーム化されます。
各チャンクの前に長さのプレフィックスがあり、特別な 0000「フラッシュ」パケットが区切り文字として使用されます。
( gitprotocol-common )。 PACK ペイロード自体は生です
コマンドセクションに続くバイト。このフレーミングこそが、プロキシがストリームを解析できる理由のすべてです
安価 — 次のレコードがどこから始まるのかが常にわかります。
強制: 許可しなかったヘッドを拒否します
これは簡単で最も価値のある半分です。 ref-update コマンドは、
パックの前に本体を POST します。それらを解析し、各ターゲット参照を許可リストと比較し、
一致しないものは、1 バイトが実サーバーに到達する前に送信されます。
POST /git-receive-pack(body) で:
# パケットを終了する 0000 フラッシュまで pkt-line を読み取ります。
# コマンドリスト;その直後にパックが始まります。
# --signed はコマンドをプッシュ証明書でラップします ...
# Push-cert-end ブロック (その後、push-options): を解析します。
# その内部からコマンドを実行し、証明書を転送します
# byte-for-byte — 署名でカバーされます
コマンド、pack = split_at_pack(body)
parse_commands(commands) の (old、new、ref) の場合:
# 例: refs/heads/agent/* のみが変更される可能性があります
許可されない場合(参照):
returnject(ref, "許可リストにないref")
is_delete(new) で、allow_deletes(ref) ではない場合:
return request(ref, "削除は許可されていません")
# 今だけ、本物の資格を持って
応答 = forward_upstream(body)
# 監査証跡 — フォローアップのキャプチャ手順
キャプチャ(レスポンス、パック)
プロキシは資格情報を保持し、エージェントは資格情報を保持していないため、これは要求ではなく壁です。エージェント
壁を嫌うものは壁を押しのけることはできません。そのトークンはプロキシを開くだけです。
これは施行の半分、つまり爆発範囲を制限する部分であり、最初に出荷できる部分です。の
他は

lf はキャプチャです: 最初にプロキシへのトラフィックを取得し、プッシュされたパックを読み取り、
すべてのコミットをログに記録し、回復のためにオブジェクトを存続させます。それはかなりの話題であることが判明しました
独自の — MITM とエンドポイント分類、シン パックと非シン トリック、パックですべてが得られる理由
差分ではなくファイル、バックアップマーカー参照、署名されたアトミックファイルの真ん中にあることの副作用
プロトコル — したがって、ビルドについては独自のフォローアップ記事を投稿します。もし
実行することはただ 1 つだけです。それは強制を行うことです。監査証跡も必要な場合は、そのフォローアップが必要です
構築されます。
解決策 2: Git フック — より軽量なコンテナ内の代替手段
プロキシでは、ネットワーク パスを制御する必要があります。エージェントのイメージをコントロールするのはあなた自身です
が流れ込みますが、出口は何でもあります。その場合、git フックが取得されます
ワイヤープロトコルの作業は一切なく、出所の大部分はあなたにあります。
アイデアは、すべてのコミットに作成時にセッション ID をスタンプし、プッシュ時にアーティファクトをキャプチャすることです。
prepare-commit-msg は主要なタグ付けポイントです。コミット、修正、マージ、チェリーピック、および
rebase reword/squash/fixup パス — そして重要なことに、これは git commit によって抑制されません
--no-verify ( pre-commit や commit-msg とは異なります)。それがまさにそれがプライマリである理由です
タグ: エージェントが逃れることのできないコミット時の 1 つのフック。追加してもらいます
Agent-Session: <id> のようなトレーラーをメッセージに追加します。あ
トレーラーはメッセージの最後にある構造化された Key: value 行であり、トレーラーと同じ仕組みです。
Signed-off-by なので、リベース後も存続し、後で grep できます。既存の価値を再収集し、
1 つのブロックを再発行するため、スカッシュは個別のセッションを保持し、同一のセッションを重複排除します。
commit-msg はシン バリデータです。トレーラーが存在することを確認する 2 番目の機会です。のように
その他、それはobのみです

保存し、コミットをブロックすることはありません。
post-rewrite は、修正/リベース後に古い→新しいコミット マッピングを記録するため、来歴はリライトに従います。
(詳細に注意してください。標準入力はタブではなくスペースで区切られています。)
プリプッシュは、強制と捕獲のネットです。タグのないコミットが見つかった場合は、全体を書き換えます。
チェックアウトなしでトポロジ順にプッシュされたサフィックス
git commit-tree — リベースがないため、ダーティ ツリーや
空コミットの危険 — 各子孫をその書き換えられた祖先に再親化する (そしてすべての親を維持する)
-p を繰り返してマージを実行した後、ブランチを移動してプッシュを中止し、修正されたコミットが保存されるようにします。
再試行中です。この自動書き換えは、プッシュされた <local-ref> が実際の場合にのみ意味を持ちます。
refs/heads/* ブランチ;削除、デタッチされた HEAD プッシュ、または raw-SHA プッシュの場合、移動するブランチはありません。
代わりに指示に従って中止します。クリーンプッシュでキャプチャします -
git Bundle ( git Bundle verify で検証)、コミットごと
パッチと小さなmetadata.json。ただし、これが何であるかについては正直に言ってください。
リモート応答であるため、これらはプッシュ試行アーティファクトです。エージェントが送信しようとした内容を記録するものではありません。
サーバーが最終的に受け入れたもの。 「実際に到達したもの」については、プロキシまたはサーバー側のフックが必要です。
フックをすべてのリポジトリ (/workspace とエージェントが後で作成するクローン) に適用するには、次のコードを実行します。
ファイルをそれぞれにコピーするのではなく、core.hooksPath を介してそれらを作成します。
.git/フック 。ルート所有の読み取り専用イメージにフックをベイクし、core.hooksPath を
root が所有する /etc/gitconfig 。
# Dockerfile (スケッチ)
COPY --chmod=555 --chown=root:root githooks/ /opt/githooks/
git config --system core.hooksPath /opt/githooks を実行します。
# エントリポイント内の 1 つの正規エージェント セッション ID を解決し、それを固定の既知のファイルに書き込みます
# (/.session_id と言う) 番目

フックでの読み取り — 環境変数ではありません。エージェントがファイルを設定解除したり、
# 環境よりも forge し、サブシェルや `git` ラッパーでも存続します。所有者が書き込み可能なものを事前に作成する
# イメージ内 (エントリポイントは非ルートで実行されます)。ファイルが見つからない、または空のファイルは単に「タグ付けをスキップ」することを意味します。
# フックは依然としてブロックされません。
明白な環境変数ではなくファイルである理由: エージェントはその環境を共有します。設定を解除できます
変数を使用してオーバーライドするか、スクラブされた環境で git を生成します。これは、先ほど説明したのと同じクラスのバイパスです。
GIT_CONFIG_* 。エージェントの実行前にエントリポイントが一度書き込む固定パスは、より小さなターゲットです。それは
まだサンドボックス内にあるため、決意を持ったエージェントであればまだ倒すことができますが、簡単なフットガンは取り除かれます。
これがベストエフォート層の適切なバーです。
フックはエージェントまたは IDE によって破壊される可能性があります - それは予想されることです
ここは、ほとんどの記事で省略される部分であり、フックが最初の解決策ではなく 2 番目の解決策である理由です。
フックはエージェントのサンドボックス内で実行されます。つまり、フックはエージェントが十分な情報を持っている場所に常駐します。
それを倒す許可。正直な穴のリスト:
構成の優先順位は間違った方向に働きます。 /etc/gitconfig は
それは「システム」であるから権威があるのです。これはその逆で、システム設定の優先順位が最も低く、
ユーザーごとの ~/.gitconfig オーバーリ

[切り捨てられた]

## Original Extract

When you hand a repo to an autonomous AI Agent, you need to bound what it can touch and capture what it did. Part 1 — the approach — keep the keys away from the agent, a git proxy as the hard boundary, and git hooks as the lighter alternative.

Founding Engineering Leader | Agentic AI DevTools & Experience
What Did the Agent Just Push? Auditing Git for AI Agents
Post
Post
You hand a repository to an autonomous AI Agent and let it work. It clones, branches, commits, pushes,
maybe opens a pull request. A few minutes later it tells you it is done. Two questions stay on the table,
and neither has a comfortable answer:
Which commits did it actually produce? Not the summary it wrote for you — the real objects that
landed on the server.
Are you sure it only touched the branches you allowed? One confused agent force-pushing main is
a bad afternoon.
I have been running fleets of agents on the same codebases for a while now (see
Orchestrating AI Fleets and the
Git Fork Pattern ), and the more agents you run, the
more these two questions turn from curiosity into operational risk. This post is about answering them at
the layer where the truth lives: git itself .
This post is the approach: why this is a git problem, the one rule that makes everything else possible, and
the two mechanisms — a git proxy and git hooks — at a level you can reason about and act on. The deepest
wire-protocol mechanics of the proxy are involved enough that they get their own follow-up; here we cover
the shape and the decisions, with enough pseudo-code that your AI Agent can start implementing it in your
stack.
Why This Is a Git Problem, Not a Policy Document
The naive answer is “just tell the agent which branches it may touch.” That is a policy. Policies are
suggestions to a process that has a shell, a network, and a strong urge to finish its task. Agents are
smart — give one a constraint phrased as a request and it may find a creative way around it, not out of
bad intent but because working around obstacles is literally what we trained it to do.
So we do not rely on asking nicely. We rely on two properties:
Blast radius — only the heads we selected can change on the real server, enforced somewhere the
agent cannot route around.
Provenance — every commit the agent created is captured and attributed, so we can recover, review,
and audit it after the fact. This will be a great source for iterative learning later.
There are two ways to get there. The first is a git proxy that sits outside the agent’s sandbox — the
hard boundary. The second is git hooks baked into the sandbox image — a lighter, best-effort layer for
when you own the image but not the network path. We use the proxy as the real wall and treat hooks as
extra telemetry. Here is why.
The One Rule Underneath Everything: Don’t Give Agents the Real Keys
Before any proxy or hook, there is a decision that makes the rest possible.
The agent never holds the credentials to the real git server.
Inside the sandbox, the agent gets a short-lived, scoped token that authenticates it to our endpoint —
not to GitHub or Space, not to your internal git host. We inject that token the boring way: a git credential helper
or a URL rewrite, never the real key on disk. The real push credential lives in the proxy, which the agent
cannot read. When a push is allowed, the proxy authenticates upstream with its own service credential.
This is the property that survives a clever agent. A sandbox is a permission boundary, and inside its own
sandbox an agent may have enough rights to do almost anything — read its environment, rewrite files,
disable the tooling we installed. If the real key were anywhere in that environment, a determined agent
could find it and talk to the Git server directly, and every control we built would be decoration. Because the
key is not there, the worst an escaped agent can reach is the proxy — and the proxy enforces policy and
holds the key. The blast radius is bounded by where the credential lives, not by how well-behaved the
agent is
I want to note – the pattern how we are going to limit the agent push/pull operations is not standard,
and potentially is hard-to-impossible to implement out of the box with GitHub, Space or any third-party Git host.
Agentic AI opens yet another direction for us – software is now very cheap to build, given you know how.
Below we explain enough ideas to implement it. Git supports multiple protocols, including HTTP, SSH, and Git over TCP .
We decided to use HTTP below, but the same technique can be applied to other protocols as well, mainly to SSH.
Solution 1: The Git Proxy — the Hard Boundary
A git push over HTTP is a surprisingly inspectable conversation. If you put a proxy on the path between the
agent and the server, you can read the whole exchange, decide whether to allow it, and capture everything
that flows through — without the agent being able to tell the difference, and without it being able to go
around you, because it has nowhere else to go
What a push actually looks like on the wire
Git’s “smart HTTP” protocol does a push in two HTTP requests
( gitprotocol-http ):
Ref advertisement — GET /info/refs?service=git-receive-pack . After a # service=git-receive-pack
header line and a flush, the server lists every ref and its current object id, and on the first ref
line it appends a NUL byte followed by the
capabilities it supports.
The push itself — POST /git-receive-pack . The body is a sequence of commands (“update
refs/heads/x from <old> to <new> ”) followed by a PACK — the binary blob of objects being
transferred.
The advertisement, the command list and the status report are all framed in pkt-lines : a 4-byte hex
length prefix in front of each chunk, with a special 0000 “flush” packet as a separator
( gitprotocol-common ). The PACK payload itself is raw
bytes that follow the command section. That framing is the whole reason a proxy can parse the stream
cheaply — you always know where the next record starts.
Enforcement: reject the heads you didn’t allow
This is the easy and most valuable half. The ref-update commands sit in plain pkt-lines at the front of the
POST body, before the pack. Parse them, compare each target ref against your allow-list, and refuse the
ones that don’t match — before a single byte reaches the real server.
on POST /git-receive-pack(body):
# read pkt-lines until the 0000 flush that ends the
# command list; the pack begins right after it.
# --signed wraps the commands in a push-cert ...
# push-cert-end block (then push-options): parse the
# commands from inside it and forward the cert
# byte-for-byte — the signature covers it
commands, pack = split_at_pack(body)
for (old, new, ref) in parse_commands(commands):
# e.g. only refs/heads/agent/* may change
if not allowed(ref):
return reject(ref, "ref not in allow-list")
if is_delete(new) and not allow_deletes(ref):
return reject(ref, "deletes not permitted")
# only now, with our real credential
response = forward_upstream(body)
# the audit trail — the capture step, in a follow-up
capture(response, pack)
Because the proxy holds the credential and the agent does not, this is a wall, not a request. An agent
that dislikes the wall cannot push around it — its token only opens the proxy.
That is the enforcement half — the part that bounds the blast radius, and the part you can ship first. The
other half is capture : getting the traffic to the proxy in the first place, reading the pushed pack,
logging every commit, and keeping the objects alive for recovery. It turns out to be a meaty topic on its
own — MITM and endpoint classification, thin packs and the no-thin trick, why a pack gives you whole
files and not diffs, backup marker refs, and the side-effects of sitting in the middle of a signed, atomic
protocol — so I’ll give the build its own follow-up post. If
you only ever do one thing, do enforcement; if you want the audit trail too, that follow-up is where it
gets built.
Solution 2: Git Hooks — the Lighter, In-Container Alternative
The proxy needs you to control the network path. Sometimes you don’t — you control the image the agent
runs in, but the egress is whatever it is. In that case git hooks get
you most of the provenance, with none of the wire-protocol work.
The idea: stamp every commit with a session id as it is created, and capture artifacts on push.
prepare-commit-msg is the primary tagging point. It fires for commit, amend, merge, cherry-pick and
the rebase reword/squash/fixup paths — and, crucially, it is not suppressed by git commit
--no-verify (unlike pre-commit and commit-msg , which are). That is exactly why it is the primary
tag: the one commit-time hook the agent can’t wave away. Have it append
a trailer like Agent-Session: <id> to the message. A
trailer is a structured Key: value line at the end of the message, the same machinery as
Signed-off-by , so it survives rebases and you can grep for it later. Re-collect existing values and
re-emit one block, so a squash preserves distinct sessions and dedups identical ones.
commit-msg is a thin validator — a second chance to confirm the trailer is present. Like the
others, it only observes and never blocks the commit.
post-rewrite records old→new commit mapping after amend/rebase, so provenance follows the rewrite.
(Mind the detail: its stdin is space -separated, not tab.)
pre-push is the enforcement-and-capture net. If it finds an untagged commit, it rewrites the whole
pushed suffix in topological order with a checkout-free
git commit-tree — no rebase, so no dirty-tree or
empty-commit hazards — re-parenting each descendant onto its rewritten ancestor (and keeping every parent
of a merge with repeated -p ), then moves the branch and aborts the push so the corrected commits go
out on the retry. That auto-rewrite only makes sense when the pushed <local-ref> is a real
refs/heads/* branch; for a delete, a detached-HEAD push or a raw-SHA push there is no branch to move, so
abort with instructions instead. On a clean push it captures — a
git bundle (validate it with git bundle verify ), per-commit
patches, and a small metadata.json . Be honest about what this is, though: pre-push fires before the
remote answers, so these are attempted-push artifacts — they record what the agent tried to send, not
what the server finally accepted. For “what actually landed”, you need the proxy or a server-side hook.
To make the hooks apply to every repository — /workspace and any clone the agent makes later — wire
them through core.hooksPath rather than copying files into each
.git/hooks . Bake the hooks into the image root-owned and read-only, and set core.hooksPath in the
root-owned /etc/gitconfig .
# Dockerfile (sketch)
COPY --chmod=555 --chown=root:root githooks/ /opt/githooks/
RUN git config --system core.hooksPath /opt/githooks
# resolve one canonical Agent-Session id in the entrypoint and write it to a fixed, well-known FILE
# (say /.session_id) that the hooks read — NOT an env var. A file is harder for the agent to unset or
# forge than the environment, and it survives sub-shells and `git` wrappers. Pre-create it owner-writable
# in the image (the entrypoint runs non-root); a missing or blank file just means "skip tagging" — the
# hooks still never block.
Why a file and not the obvious environment variable: the agent shares that environment. It can unset the
variable, override it, or spawn git with a scrubbed env — the same class of bypass we just saw with
GIT_CONFIG_* . A fixed path the entrypoint writes once, before the agent runs, is a smaller target. It is
still inside the sandbox, so still defeat-able by a determined agent — but it removes the easy footgun, and
that is the right bar for a best-effort layer.
Hooks Can Be Broken by the Agent or an IDE — and That’s Expected
Here is the part most write-ups skip, and the reason hooks are the second solution, not the first.
A hook runs inside the agent’s sandbox , which means it lives in a place where the agent often has enough
permissions to defeat it. The honest list of holes:
Config precedence runs the wrong way for you. It is tempting to think /etc/gitconfig is
authoritative because it is “system”. It is the opposite — system config is the lowest precedence, and
a per-user ~/.gitconfig overri

[truncated]
