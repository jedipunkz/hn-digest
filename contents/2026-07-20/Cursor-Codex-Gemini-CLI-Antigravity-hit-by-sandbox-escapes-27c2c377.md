---
source: "https://www.bleepingcomputer.com/news/security/cursor-codex-gemini-cli-antigravity-hit-by-sandbox-escapes/"
hn_url: "https://news.ycombinator.com/item?id=48986015"
title: "Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes"
article_title: "Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes"
author: "sbulaev"
captured_at: "2026-07-20T23:47:56Z"
capture_tool: "hn-digest"
hn_id: 48986015
score: 4
comments: 0
posted_at: "2026-07-20T23:07:06Z"
tags:
  - hacker-news
  - translated
---

# Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes

- HN: [48986015](https://news.ycombinator.com/item?id=48986015)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/cursor-codex-gemini-cli-antigravity-hit-by-sandbox-escapes/)
- Score: 4
- Comments: 0
- Posted: 2026-07-20T23:07:06Z

## Translation

タイトル: カーソル、コーデックス、ジェミニ CLI、サンドボックスによる反重力の脱出
説明: 研究者らは、AI エージェントに信頼できるホスト ツールが後で実行されるファイルを作成させることで、Cursor、Codex、Gemini CLI、Antigravity のサンドボックスから脱出しました。複数の CVE、パッチ、Google による反重力に関する 2 つの調査結果のダウングレード。

記事本文:
カーソル、コーデックス、ジェミニ CLI、サンドボックスに襲われた反重力がエスケープ
新しい ClickLock macOS マルウェアがユーザーを罠にかけてログイン パスワードを漏らす
アボット、恐喝の申し立ての中で2件のサイバー事件を調査
WordPress コア「wp2shell」RCE の欠陥が公開エクスプロイト、今すぐパッチ適用
新しい Windows LegacyHive ゼロデイによりハッカーに管理者権限が与えられる
エスティ ローダー、Oracle E-Business の欠陥によるデータ侵害を明らかに
SonicWall SMA1000 の欠陥をゼロデイとして悪用してカスタム マルウェアをプッシュ
ハッカーがオフチェーン攻撃でオスティアムから2,370万ドルの暗号通貨を盗む
カーソル、コーデックス、ジェミニ CLI、サンドボックスに襲われた反重力がエスケープ
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
カーソル、コーデックス、ジェミニ CLI、サンドボックスに襲われた反重力がエスケープ
カーソル、コーデックス、ジェミニ CLI、サンドボックスに襲われた反重力がエスケープ
セキュリティ研究者らは、Cursor、OpenAI の Codex、Google の Gemini CLI、Antigravity を含む 4 つの広く使用されている AI コーディング エージェントで、サンドボックスを正面から攻撃することなくサンドボックスを突破しました。
エージェントはボックス内に留まり、あらゆるルールに従います。ファイルを書き込むだけで、その後、ボックス外の信頼できるツールが実行、ロード、またはスキャンし、エスケープが自動的に行われます。
Pillar Security の研究チーム、Eilon Cohen、Dan Lisichkin、Ariel Fogel は、数カ月にわたってバイパスを再現し、今日、彼らが「Week of Sandbox Escapes」と呼ぶシリーズとして、1 日 1 回の記事として公開しました。
これらのサンドボックスは単純な線を描きます。つまり、プロジェクト ワークスペース内ではエージェントが信頼され、外部のホストは保護されます。

問題は、ワークスペース内のファイルは不活性ではないということです。サンドボックスの外で実行されているツールは、それらを読み取って動作するため、エージェントが書き込みを許可されたファイルは、後でホストが実行するコマンドに変わる可能性があります。
IDE と CLI エージェントは、サンドボックスの外側で常に独自のツールを実行します。つまり、インタープリターを解決する Python 拡張機能、リポジトリをスキャンする Git 統合、タスク ファイルを実行する VS Code、コマンドを実行するフック エンジン、ローカル ソケットを公開する Docker デスクトップなどです。
サンドボックス エージェントは、与えられたすべてのルールに従い、コンポーネントが読み取るファイルを形成できます。
即時注入がきっかけです。 README、問題、依存関係、または差分に埋め込まれた悪意のある命令は、開発者のマシン上でローカル アクションになります。
Pillar は 7 つの調査結果を 4 つの故障モードに分類しています。
オペレーティング システムに対応できないサンドボックスを拒否リストに登録する
実際には実行可能なコードであるワークスペース構成
コマンドの引数ではなく名前を信頼する「安全な」コマンド許可リスト、および
完全にサンドボックスの外側にある特権付きローカル デーモン。
ほとんどの問題にはパッチが適用されており、ベンダーも認識しています。
Cursor では、ワークスペースで制御される .claude フック構成が、サンドボックス化されていないコマンド実行に変わりました。現在、これは CVE-2026-48124 として追跡されており、バージョン 3.0.0 で修正されました。
2 番目の Cursor バグにより、エージェントは virtualenv インタープリターを編集でき、エディターの Python 拡張機能が検出中に独自に実行されました。
3 つ目は、Git メタデータが .git というフォルダーに存在する必要がないという事実を悪用し、fsmonitor を通じて実行を開始し、Cursor のパスベースのルールをすり抜けました。 3.0.0 でパッチが適用され、CVE は保留中です。
Codex CLI では、「安全な」コマンドは信頼された git show を名前で許可リストに登録しますが、実際の呼び出しは読み取り専用ではありませんでした。 OpenAI は v0.95.0 でパッチを適用し、CVE 保留中として重大度の高い報奨金を支払いました。
1つ

Docker ソケット検索は Codex、Cursor、Gemini CLI を同時にヒットしました。エージェントがアクセスできる特権付きローカル デーモンは、コードを実行するためのサンドボックス化されていない場所になりました。この問題は現在修正されています。
Antigravity の 2 つの調査結果、macOS のシートベルト拒否リストのバイパスと、セキュア モードの .vscode タスク設定のバイパスに対する Google の反応は、よりクールなものでした。
Pillar氏によると、Googleは両方を「その他の有効なセキュリティ脆弱性」として分類し、ソーシャルエンジニアリングが必要であるか、間接的なプロンプトインジェクションを行うリポジトリをユーザーが信頼しているため、悪用するのは難しいと判断し、ダウングレードを適用したという。
しかし、ピラー氏は、Googleのチームは依然としてその成果を高く評価しており、あるレポートは「非常に質が高い」というフィードバックを引用したと述べた。
基礎となるクラスは新しいものではありません。 Cymulate は 4 月に、Claude Code、Gemini CLI、Codex CLI にわたって、「Configuration-Based Sandbox Escape」と名付けた同じパターンを文書化しました。このパターンでは、サンドボックス内に書き込まれたファイルが次回の起動時にホスト上で実行されます。
新しいのは幅の広さです。同じ障害が 3 つのベンダーの 4 つのツールで発生します。
Pillar の修正は、禁止されたファイル名の別のリストではありません。信頼できるローカル ツールがエージェントが作成した内容を実行する瞬間を監視します。
エージェント コーディング ツールを比較検討している人にとって、それはより有益なシグナルです。エージェントにサンドボックスがあるかどうかではなく、エージェントが残したファイルに何が起こるかということです。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは目に見えずに環境内を移動します。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
今すぐ更新: 7-Zip は悪意のあるアーカイブで悪用可能な RCE の欠陥を修正します
WordPress コア「wp2shell」RCE の欠陥が公開エクスプロイト、今すぐパッチ適用
HollowByteのDDoS欠陥がOpenSSを肥大化させる

11 バイトのペイロードを備えた L サーバー メモリ
Claude Chrome 拡張機能の欠陥により、悪意のある拡張機能が AI アクションをトリガーする
Zoom、重大なアカウント乗っ取りの脆弱性を警告
まだメンバーではありませんか?今すぐ登録
アーンスト・アンド・ヤング社、サポートシステムハッキング後のデータ侵害を明らかに
新しい Windows LegacyHive ゼロデイによりハッカーに管理者権限が与えられる
今すぐ更新: 7-Zip は悪意のあるアーカイブで悪用可能な RCE の欠陥を修正します
AI は脆弱性管理を破壊しました。 CISO が予算を BAS に移行するために使用するガイドを入手してください。
DARKROOMに挑戦するのに必要なものはありますか?限定の DEFCON CTF にサインアップしてください!
ポリシーによるプライバシー、それともアーキテクチャによるプライバシー?顔がデバイスから離れない場合に年齢チェックがどのように機能するかを確認してください。
Pixellot が管理されていない数百もの AI エージェントの ID を、数か月ではなく数週間で発見し、保護した方法をご覧ください。ケーススタディを読んでください。
Web アプリをオンデマンドで侵入テストします。人間が見逃しているものを見つけてください。数分で侵入テストの範囲を設定し、開始します。
MDR を交換することで節約できる金額を計算してください。
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

Researchers escaped the sandboxes in Cursor, Codex, Gemini CLI and Antigravity by having the AI agent write files that trusted host tools later run. Multiple CVEs, patches, and Google downgrading two Antigravity findings.

Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes
New ClickLock macOS malware traps users into revealing login password
Abbott probes two cyber incidents amid extortion claims
WordPress Core "wp2shell" RCE flaws get public exploits, patch now
New Windows LegacyHive zero-day gives hackers admin privileges
Estée Lauder discloses data breach via Oracle E-Business flaw
SonicWall SMA1000 flaws exploited as zero-days to push custom malware
Hackers steal $23.7 million in crypto from Ostium in off-chain attack
Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes
Cursor, Codex, Gemini CLI, Antigravity hit by sandbox escapes
Security researchers broke out of the sandboxes in four widely used AI coding agents, including Cursor, OpenAI's Codex, Google's Gemini CLI and Antigravity, without attacking the sandbox head-on.
The agent stays inside the box and follows every rule. It just writes a file that a trusted tool outside the box later runs, loads, or scans, and the escape happens on its own.
Pillar Security's research team, Eilon Cohen, Dan Lisichkin and Ariel Fogel, reproduced the bypasses over several months and published them today as a series they call the Week of Sandbox Escapes , one write-up a day.
These sandboxes draw a simple line: the agent is trusted inside the project workspace, the host outside is protected.
The catch is that files inside the workspace are not inert. Tools running outside the sandbox read and act on them, so a file the agent is allowed to write can turn into a command the host later runs.
IDEs and CLI agents constantly run their own tools outside the sandbox: Python extensions resolving interpreters, Git integrations scanning repos, VS Code running task files, hook engines firing commands, Docker Desktop exposing a local socket.
A sandboxed agent can obey every rule it is given and still shape the files those components read.
Prompt injection is the trigger. A malicious instruction planted in a README, an issue, a dependency or a diff becomes a local action on the developer's machine.
Pillar sorts the seven findings into four failure modes:
denylist sandboxes that cannot keep pace with the operating system
workspace config that is really executable code
"safe" command allowlists that trust a command's name rather than its arguments, and
privileged local daemons that sit outside the sandbox entirely.
Most of the issues are patched and vendor-acknowledged.
In Cursor, a workspace-controlled .claude hook config turned into unsandboxed command execution. It is now tracked as CVE-2026-48124 and was fixed in version 3.0.0.
A second Cursor bug let the agent edit a virtualenv interpreter that the editor's Python extension then ran on its own during discovery.
A third abused the fact that Git metadata does not have to live in a folder called .git, firing execution through fsmonitor and slipping past Cursor's path-based rules. It was patched in 3.0.0, with a CVE pending.
In Codex CLI, a "safe" command allowlist trusted git show by name while the actual invocation was not read-only. OpenAI patched it in v0.95.0 and paid a high-severity bounty, with a CVE pending.
One Docker socket finding hit Codex, Cursor and Gemini CLI at once. A privileged local daemon the agents could reach became an unsandboxed place to run code. That issue is now fixed.
Google's response to the two Antigravity findings, a macOS Seatbelt denylist bypass and a .vscode task-config bypass of its Secure Mode, was cooler.
According to Pillar, Google classified both as "Other valid security vulnerabilities" and applied a downgrade, judging them difficult to exploit because they need social engineering or a user trusting a repository that carries an indirect prompt injection.
Pillar says, however, Google's team still rated the work highly, quoting feedback that one report was "of exceptional quality."
The underlying class is not new. In April, Cymulate documented the same pattern, which it named "Configuration-Based Sandbox Escape," across Claude Code, Gemini CLI and Codex CLI, where a file written inside the sandbox runs on the host at the next launch.
What is new is the breadth. The same failure shows up across four tools from three vendors.
Pillar's fix is not another list of banned filenames. It watches the moment a trusted local tool runs something the agent wrote.
For anyone weighing agentic coding tools, that is the more useful signal: not whether an agent has a sandbox, but what happens to the files it leaves behind.
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
Update now: 7-Zip fixes RCE flaw exploitable with malicious archives
WordPress Core "wp2shell" RCE flaws get public exploits, patch now
HollowByte DDoS flaw bloats OpenSSL server memory with 11-byte payload
Claude Chrome extension flaw lets malicious extensions trigger AI actions
Zoom warns of critical account takeover vulnerability
Not a member yet? Register Now
Ernst & Young discloses data breach after support system hack
New Windows LegacyHive zero-day gives hackers admin privileges
Update now: 7-Zip fixes RCE flaw exploitable with malicious archives
AI broke vulnerability management. Get the guide CISOs use to shift budget to BAS.
Do you have what it takes to challenge DARKROOM? Signup for an exclusive DEFCON CTF!
Privacy by policy or privacy by architecture? See how age checks work when the face never leaves the device.
See how Pixellot discovered and secured hundreds of unmanaged AI agent identities in weeks, not months. Read the case study.
Pentest your web apps on-demand. Find what humans miss. Scope and launch a pentest in minutes.
Calculate what you’d save by replacing your MDR.
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.
