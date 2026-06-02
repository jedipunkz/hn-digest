---
source: "https://brentfitzgerald.com/posts/lima-vms-from-a-sandboxed-agent/"
hn_url: "https://news.ycombinator.com/item?id=48360921"
title: "Starting a VM from a macOS sandbox for AI agents"
article_title: "Starting a VM from a MacOS sandbox | Brent Fitzgerald"
author: "_doctor_love"
captured_at: "2026-06-02T00:34:18Z"
capture_tool: "hn-digest"
hn_id: 48360921
score: 2
comments: 1
posted_at: "2026-06-01T18:43:04Z"
tags:
  - hacker-news
  - translated
---

# Starting a VM from a macOS sandbox for AI agents

- HN: [48360921](https://news.ycombinator.com/item?id=48360921)
- Source: [brentfitzgerald.com](https://brentfitzgerald.com/posts/lima-vms-from-a-sandboxed-agent/)
- Score: 2
- Comments: 1
- Posted: 2026-06-01T18:43:04Z

## Translation

タイトル: AI エージェント用の macOS サンドボックスからの VM の起動
記事のタイトル: MacOS サンドボックスから VM を起動する |ブレント・フィッツジェラルド
説明: シートベルトにより、仮想化を使用するプロセスのルールが許可されます

記事本文:
MacOS サンドボックスから VM を起動する |ブレント・フィッツジェラルド
ブレント・フィッツジェラルド
執筆
お読みください
今
私と一緒に働きましょう
メニュー
ブレント・フィッツジェラルド
閉じる
書く
今
仕事
お読みください
*]:pointer-events-auto" data-astro-cid-bbe6dxrz> 2026 年 5 月 27 日のエッセイ Brent Fitzgerald MacOS サンドボックスからの VM の起動
シートベルトにより、仮想化を使用するプロセスのルールが許可されます
これは、Lima を使用して MacOS サンドボックス プロセスから VM を起動することに関する簡単な技術メモです。
複数のワークツリーが進行中で、それぞれのワークツリーに特定の機能を実行するエージェントが含まれることが多くなりました。すべてのサービスの依存関係とポートを分離するために、私は Lima で使い捨て Linux VM を使用し始め、VM 内で docker compose やその他のアプリを実行します。
私は、コーディング エージェントにもこれを実行させることが非常に役立つと判断しました。そうすることで、Docker Compose ポートの衝突や共有データベースを発生させることなく、完全な開発とテストのライフサイクルを完全に分離して実行できます。
シートベルトとノノのサンドボックス
各コーディング エージェントをサンドボックスで実行します。私はエージェントが提供するサンドボックスには少し疑問を感じています。サンドボックスは時間の経過とともに変更される可能性があり、自動モードなどのオプションが少し賢すぎるように思えます。そのため、別のシステムを使用しています。
Mac では、サンドボックス化には通常、 Sandbox-exec によって使用されるカーネルレベルのエンジンである Apple の Seatbelt が使用されます。シートベルト プロファイルは、何を拒否し、何を許可するかについての非常に明確なルールの束です。
私のサンドボックスは、最初に Agent Safehouse で作成したシートベルト プロファイルを使用します。
ただし、私は最近、シートベルトを内部で使用する機能ベースのサンドボックスである nono も試しています。理論的には、nono を使用すると、(グループ経由で) 機能を簡単に構成したり、既製のプロファイルから継承したり上書きしたりすることが容易になります。
リマと仮想化の権利
Apple Silicon 上の Lima のドライバーは vz であり、Ap を使用します

ple の仮想化フレームワーク。フレームワークには権限が必要です。 limactl はすでに同梱されています。確認方法は次のとおりです。
$ codesign -d --entitlements - /opt/homebrew/bin/limactl
...
[キー] com.apple.security.virtualization
[値]
[ブール] true
したがって、バイナリは VM を作成できます。ここまでは順調ですね！
残念ながら、サンドボックス プロファイルを使用して liactl start を実行すると、エラーが発生して失敗しました。
[hostagent] VZの起動
[hostagent] Rosetta共有の設定
致命的: エラー ドメイン=VZErrorDomain コード=1 説明="内部仮想化エラー。仮想マシンを起動できませんでした。"
さらに、その他の問題もいくつかあります。
システム DNS の検出に失敗しました。[8.8.8.8 1.1.1.1] にフォールバックします。 error="/etc/resolv.conf を開く: 操作は許可されていません"
DNS 警告は /etc/resolv.conf への読み取りアクセスの問題であり、機密性の高いものは何もありません。本当の問題は VZErrorDomain Code=1 です。 「内部仮想化エラー」と表示されますが、何が拒否されたのかは正確には示されていません。
「Rosetta 共有のセットアップ」の直後にエラーが発生したため、最初は Rosetta の問題かと思いました (実際にはそうではありませんでした)。
シートベルト着用拒否はリマの記録には残されていなかった。 Lima には一般的な VZ エラーのみが表示されます。 Seatbelt ログを見つけるには、カーネルによって書き込まれた MacOS 統合ログを確認する必要がありました。
最初のログ表示クエリでは何も見つかりませんでした。拒否はカーネルによって特定の形式で出力されるため、使用する適切な述語を理解するのは困難でした。最終的には、リテラルの Sandbox: プレフィックスでフィルタリングすることでそれを見つけました。
log stream --style Compact --predicate 'eventMessage CONTAINS "サンドボックス: "'
私が理解した拒否をトリガーしてこれをテストしました。つまり、サンドボックスの下でブロックされたファイルを読み取るということです。ログには次のことが示されていました。
サンドボックス: wc(89150)deny(1) file-read-data /private/etc/sudoers
いいですね。ロギングは機能します。形式はサンです

dbox: <プロセス>(<pid>)deny(1) <操作> <詳細> 。
ストリームが実行されている状態で、VM を再度起動しました。私が見た否定文は次のとおりです。
サンドボックス: limactl(42303)deny(1) mach-lookup com.apple.Virtualization.VirtualMachine (xpc)
そのため、VM を起動するために、仮想化フレームワークは com.apple.Virtualization.VirtualMachine という名前の XPC サービスを検索しようとします。そのサービスは、実際にゲストをホストするプロセスです。サンドボックスでは検索が許可されず、検索は失敗し、フレームワークは詳細なしで Code=1 を返しました。
エージェントのサポートを受けて調査した結果、特定のマッハ ルックアップを許可するルールを作成する方法があることがわかりました。
(マッハルックアップを許可します
(グローバル名 " com.apple.Virtualization.VirtualMachine " ))
それからもう一度試してみました。私はその否定を乗り越えましたが、さらに進むにつれて、さらに多くの否定が次々と現れました。
limactldeny(1) mach-task-name その他 [com.apple.Virtualization.Virtual(44215)]
limactl 拒否 (1) 汎用問題拡張機能拡張クラス:com.apple.virtualization.extension.fuse
limactldeny(1) generic-issue-extension extension-class:com.apple.virtualization.extension.rosetta-directory-share
それぞれが VM を起動する手順です。
mach-task-name には少し時間がかかりました。これは、生成したばかりのホスト プロセスへのハンドルを取得するフレームワークです。このプロセスは、launchd によって開始されるため、サンドボックスの外で実行されます。親を監視するには、親のタスク名が必要です。
Fuse 拡張子は virtio-fs ディレクトリ共有であり、ホスト ファイルシステムを VM にマウントするために使用されます。
Rosetta-directory-share 拡張機能は、ゲストが x86-64 バイナリを実行できるようにする Rosetta マウントです。 VM はそれなしで起動しますが、許可しない限り、Rosetta はゲストで動作しません。
Lima VM の起動を可能にするシートベルト ルール
シートベルトに関して、私が変更しなければならなかったのは次の 4 つです。 3 つは仮想化固有のものです。
(マッハルックアップを許可します
(グロ

bal-name " com.apple.Virtualization.VirtualMachine " ))
(マッハタスク名を許可)
(一般的な問題の拡張子を許可します
(拡張クラス " com.apple.virtualization.extension.fuse " )
(拡張クラス " com.apple.virtualization.extension.rosetta-directory-share " ))
4 番目は単なるファイルアクセスでした。 Lima は、インスタンスごとの状態とソケットを ~/.lima に保持し、ゲスト イメージを ~/Library/Caches/lima にダウンロードします。
(ファイル読み取り* ファイル書き込み* を許可します
(サブパス「 /Users/you/.lima 」)
(サブパス " /Users/you/Library/Caches/lima " ))
DNS 警告も修正する価値がありました。 /etc/resolv.conf は、 /var/run/resolv.conf へのシンボリックリンクです。私のプロファイルではシンボリックリンク パスは許可されましたが、実際のターゲットは許可されなかったため、読み取りに失敗し、Lima は他の DNS にフォールバックしました。この修正により、ファイルの読み取りが可能になります。
(ファイル読み取りを許可*
(リテラル " /private/var/run/resolv.conf " ))
そこで、.sb ファイル内の Lima ファイル許可の隣にある小さなブロックに仮想化ルールを追加しました。
これで、VM がサンドボックス内で READY 状態で起動します。わーい！
(allow mach-task-name) に関する注意点の 1 つは、範囲が広いということです。 Seatbelt を使用すると、タスクポート ルールのスコープを Same-sandbox に設定できますが、VM ホスト プロセスがサンドボックス内にないため、そのスコープは一致しません。単一の実行可能ファイルにスコープを設定する明確な方法がわかりません。リスクは小さいと思います。タスク名はイントロスペクションに使用されますが、プロセスは依然としてターゲットのメモリを読み取ったり、その実行を制御したりすることはできません。
nono のモデルはより高いレベルなので、s 式の構文に触れる必要はありません。それはファイル、ネットワーク、資格情報です。ディレクトリとドメインを付与しますが、mach-lookup またはサンドボックス拡張機能のスキーマ サポートが見つかりませんでした。また、Lima または仮想化用の組み込みグループも見つかりませんでした。
ただし、nono は Mac の内部でシートベルトを使用するため、カスタムを提供するために unsafe_macos_seatbelt_rules を公開します。

ああ、型なしルール。これは、シートベルトの S 式をそのまま適用したリストです。このファイルは、通常の nono 機能へのマップを許可します。そこには 3 つの仮想化ルールが含まれます。
{
「ファイルシステム」: {
"allow" : [ " ~/.lima " , " ~/Library/Caches/lima " ] ,
"read_file" : [ " /private/var/run/resolv.conf " ]
} 、
"unsafe_macos_seatbelt_rules" : [
" (allow mach-lookup (グローバル名 \" com.apple.Virtualization.VirtualMachine \" )) " ,
" (allow mach-task-name) " ,
" (allow generic-issue-extension (extension-class \" com.apple.virtualization.extension.fuse \" ) (extension-class \" com.apple.virtualization.extension.rosetta-directory-share \" )) "
]
}
したがって、Lima ファイルの許可は簡単ですが、子プロセスに XPC ホスト サービスを使用させてタスク名を取得させるのは、生のシートベルトが必要な部分です。
これを nono 「パック」として提供することを検討しましたが、現在 nono は複数のプロファイルの作成をサポートしていないため、これを使用している人は「クロード + リマのサポートが欲しい」とは言えません。 nono の構成メカニズムはグループであるようですが、グループは含めるか除外するように設計されているため、プロファイルで使用できる一般的な unsafe_macos_seatbelt_rules エントリはサポートされていません。
他の旅行者向けの持ち帰りメモ
VZErrorDomain Code=1 はあまり意味がありませんが、仮想化フレームワークがサンドボックス内で起動できない場合は、おそらく mach-lookup が拒否されたことが原因です。
シートベルトの拒否は、カーネルによって書き込まれた統合ログの Sandbox: プレフィックスの下に記録されます。したがって、それを探しているストリームをログに記録できます。
Lima vz が limactl 経由で VM を起動するための最小セット: com.apple.Virtualization.VirtualMachine mach-lookup、ホスト プロセスの mach-task-name、および 2 つの仮想化問題拡張機能。さらに、~/.lima とそのキャッシュへのファイル アクセス。
についての nono ディスカッションを開始しました

これらのルールを文書化し、パックにまとめ、場合によってはそれらを一流の機能に昇格させる必要があります。

## Original Extract

Seatbelt allow rules for processes to use virtualization

Starting a VM from a MacOS sandbox | Brent Fitzgerald
Brent Fitzgerald
Writing
README
Now
Work with me
Menu
Brent Fitzgerald
Close
WRITING
NOW
WORK
README
*]:pointer-events-auto" data-astro-cid-bbe6dxrz> May 27, 2026 essay Brent Fitzgerald Starting a VM from a MacOS sandbox
Seatbelt allow rules for processes to use virtualization
This is a quick tech note about using Lima to start VMs from MacOS sandboxed processes.
Increasingly often I have multiple worktrees in progress, each with an agent working on a specific feature. To isolate all service dependencies and ports, I’ve been starting to use throwaway Linux VMs with Lima, and then I’ll run docker compose and other app stuff inside of the VM.
I decided it would be really helpful to make my coding agents also do this. That way they can do a full development and test lifecycle in full isolation without any docker compose port collisions or shared databases.
Sandboxes with Seatbelt and nono
I run each coding agent in a sandbox. I am a little suspicious of the agent-provided sandboxes, which may change over time and seem a bit too clever with options like auto mode. So I use a separate system.
On a Mac, sandboxing typically uses Apple’s Seatbelt, the kernel-level engine used by sandbox-exec . A seatbelt profile is a bunch of very explicit rules about what to deny and allow.
My sandbox uses a Seatbelt profile that I initially constructed with Agent Safehouse .
However, I’ve also been recently experimenting with nono , a capability-based sandbox that also uses Seatbelt under the hood. Theoretically nono can make it easier to compose capabilities (via groups) and to inherit from and override readymade profiles.
Lima and the Virtualization entitlement
Lima’s driver on Apple Silicon is vz , which uses Apple’s Virtualization framework. The framework requires an entitlement. limactl already ships with it. Here’s how you can see it:
$ codesign -d --entitlements - /opt/homebrew/bin/limactl
...
[Key] com.apple.security.virtualization
[Value]
[Bool] true
So the binary is allowed to make VMs. So far so good!
Unfortunately running limactl start with my sandbox profile failed with an error:
[hostagent] Starting VZ
[hostagent] Setting up Rosetta share
fatal: Error Domain=VZErrorDomain Code=1 Description="Internal Virtualization error. The virtual machine failed to start."
Plus some other issues:
failed to detect system DNS, falling back to [8.8.8.8 1.1.1.1] error="open /etc/resolv.conf: operation not permitted"
The DNS warning is just an issue of read access on /etc/resolv.conf , and there’s nothing sensitive in there. The real problem is VZErrorDomain Code=1 . It says “Internal Virtualization error” but does not say what was denied exactly.
Since the error occurred right after “Setting up Rosetta share” I thought at first it looked like a Rosetta problem (which turned out to not be the case).
The Seatbelt denial was not in Lima’s logs. Lima only sees the generic VZ error. To find the Seatbelt logs I had to look at the MacOS unified log, written by the kernel.
My initial log show queries found nothing. The denials are emitted by the kernel in a specific format and it was hard to figure out the right predicate to use. I eventually found it by filtering on the literal Sandbox: prefix:
log stream --style compact --predicate 'eventMessage CONTAINS "Sandbox: "'
I tested this by triggering a denial I understood: read a blocked file under the sandbox. The log showed it:
Sandbox: wc(89150) deny(1) file-read-data /private/etc/sudoers
Good. Logging works. The format is Sandbox: <process>(<pid>) deny(1) <operation> <detail> .
With the stream running, I started the VM again. Here’s the deny I saw:
Sandbox: limactl(42303) deny(1) mach-lookup com.apple.Virtualization.VirtualMachine (xpc)
So to start a VM, the Virtualization framework tries to look up an XPC service named com.apple.Virtualization.VirtualMachine . That service is the process that actually hosts the guest. The sandbox did not allow the lookup, the lookup failed, and the framework returned Code=1 with no detail.
After some research and with agent assistance, I learned there’s a way to write a rule to allow a specific mach-lookup:
(allow mach-lookup
(global-name " com.apple.Virtualization.VirtualMachine " ))
Then I tried again. I got past that denial, and more denials appeared, in sequence, as we got further:
limactl deny(1) mach-task-name others [com.apple.Virtualization.Virtual(44215)]
limactl deny(1) generic-issue-extension extension-class:com.apple.virtualization.extension.fuse
limactl deny(1) generic-issue-extension extension-class:com.apple.virtualization.extension.rosetta-directory-share
Each one is a step in starting the VM:
mach-task-name took me a bit. It is the framework getting a handle to the host process it just spawned. That process runs outside the sandbox since launchd starts it. The parent needs its task-name to monitor it.
The fuse extension is the virtio-fs directory share, used to mount the host filesystem into the VM.
The rosetta-directory-share extension is the Rosetta mount that lets the guest run x86-64 binaries. The VM boots without it, but Rosetta will not work in the guest unless you allow it.
Seatbelt rules to enable Lima VM start
Here are the four things I had to change, in Seatbelt form. Three are virtualization-specific:
(allow mach-lookup
(global-name " com.apple.Virtualization.VirtualMachine " ))
(allow mach-task-name)
(allow generic-issue-extension
(extension-class " com.apple.virtualization.extension.fuse " )
(extension-class " com.apple.virtualization.extension.rosetta-directory-share " ))
The fourth was just file access. Lima keeps per-instance state and sockets under ~/.lima , and downloads guest images into ~/Library/Caches/lima :
(allow file-read* file-write*
(subpath " /Users/you/.lima " )
(subpath " /Users/you/Library/Caches/lima " ))
The DNS warning was worth fixing too. /etc/resolv.conf is a symlink to /var/run/resolv.conf . My profile allowed the symlink path but not the real target, so the read failed and Lima fell back to other DNS. The fix is allowing a file-read:
(allow file-read*
(literal " /private/var/run/resolv.conf " ))
So I added the virtualization rules in a small block in my .sb file next to Lima file grants.
Now the VM boots to READY inside the sandbox. Yay!
One note on (allow mach-task-name) is that it’s broad, and I would rather it weren’t. Seatbelt lets you scope task-port rules to same-sandbox , but the VM host process is not in the sandbox, so that scope does not match it. I don’t see a clean way to scope to a single executable. I think the risk is small. The task name is used for introspection, but the process still cannot read the target’s memory or control its execution.
nono’s model is higher level so you don’t have to touch the s-expression syntax. It’s files, network, and credentials. You grant directories and domains, but I couldn’t find schema support for mach-lookup or a sandbox extension, and no built-in group for Lima or virtualization.
However, since it uses Seatbelt under the hood on Macs, nono does expose unsafe_macos_seatbelt_rules to provide custom, untyped rules. It’s a list of raw Seatbelt S-expressions applied verbatim. The file grants map to normal nono capabilities. The three virtualization rules go in there:
{
"filesystem" : {
"allow" : [ " ~/.lima " , " ~/Library/Caches/lima " ] ,
"read_file" : [ " /private/var/run/resolv.conf " ]
} ,
"unsafe_macos_seatbelt_rules" : [
" (allow mach-lookup (global-name \" com.apple.Virtualization.VirtualMachine \" )) " ,
" (allow mach-task-name) " ,
" (allow generic-issue-extension (extension-class \" com.apple.virtualization.extension.fuse \" ) (extension-class \" com.apple.virtualization.extension.rosetta-directory-share \" )) "
]
}
So the Lima file grants are straightforward, but letting the child process use an XPC host service and get task names is the part where we need raw Seatbelt.
I’ve looked into contributing this as a nono “pack” , but currently nono doesn’t support composing multiple profiles, so someone using this couldn’t say “I want claude + lima support.” nono’s composition mechanism seems to be groups, but because groups are designed to be either included or excluded, they don’t support the generic unsafe_macos_seatbelt_rules entry available to profiles.
Takeaway notes for other travelers
VZErrorDomain Code=1 doesn’t say much, but if the Virtualization framework fails to start inside a sandbox, it’s probably the denied mach-lookup .
Seatbelt denials are in the unified log, written by the kernel, under the Sandbox: prefix. So you can log stream looking for that.
The minimal set for Lima vz to start a VM via limactl : the com.apple.Virtualization.VirtualMachine mach-lookup, mach-task-name for the host process, and the two virtualization issue-extensions. Plus file access to ~/.lima and its cache.
I’ve started a nono discussion about documenting these rules, composing them into a pack, and maybe promoting them to a first-class capability.
