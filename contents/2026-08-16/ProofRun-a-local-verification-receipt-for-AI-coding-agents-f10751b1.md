---
source: "https://github.com/yebiguo/ProofRun"
hn_url: "https://news.ycombinator.com/item?id=49316605"
title: "ProofRun – a local verification receipt for AI coding agents"
article_title: "GitHub - yebiguo/ProofRun: A local verification receipt for AI coding agents · GitHub"
author: "yebiguo"
captured_at: "2026-08-16T03:40:17Z"
capture_tool: "hn-digest"
hn_id: 49316605
score: 1
comments: 0
posted_at: "2026-08-16T03:22:15Z"
tags:
  - hacker-news
  - translated
---

# ProofRun – a local verification receipt for AI coding agents

- HN: [49316605](https://news.ycombinator.com/item?id=49316605)
- Source: [github.com](https://github.com/yebiguo/ProofRun)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T03:22:15Z

## Translation

タイトル: ProofRun – AI コーディング エージェントのローカル検証レシート
記事タイトル: GitHub - yebiguo/ProofRun: AI コーディング エージェントのローカル検証レシート · GitHub
説明: AI コーディング エージェントのローカル検証レシート。 GitHub でアカウントを作成して、yebiguo/ProofRun の開発に貢献してください。

記事本文:
GitHub - yebiguo/ProofRun: AI コーディング エージェントのローカル検証レシート · GitHub
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
エビグオ
/
プルーフラン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
37 コミット 37 コミット .github/ workflows .github/ workflows cmd/proofrun cmd/proofrun docs docs 内部 内部 .gitignore .gitignore .goreleaser.yml .goreleaser.yml .proofrun.yml .proofrun.yml AGENTS.md AGENTS.md ライセンス LI

CENSE README.md README.md README.zh-CN.md README.zh-CN.md action.yml action.yml go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントのローカル検証レシート。
ProofRun はコードが正しいかどうかを判断しません。これは、上手に質問することではなく、暗号化によって、現在所有している正確なコードに対してどのチェックが実際に実行されたかを証明します。
AI コーディング エージェントは「すべてのテストに合格した」と言います。それは本当ですか？
多分。エージェントが最後に実際にテストを実行したときもそうでした。しかし、それは 3 回前の編集だったかもしれません。エージェントは、それらを実行したことさえ覚えていない可能性があります。単に「変更は正しいようで、おそらくテストはまだパスしているだろう」と推測しているだけかもしれません。言葉だけからは、「きっと合格するだろう」ということ以外に、「実行したら合格した」と判断する方法はありません。
ProofRun はそのギャップを埋めます。エージェントをより正直にすることではなく、主張自体をチェック可能にすることによってです。
$proofrun 実行テスト -- pytest
...
テスト: 合格 (終了 0、1841ms)
$ プルーフィング ステータス
テスト PASS (出口 0、1841ms)
# この時点以降のコード変更 — エージェントか人間かは関係ありません
$ プルーフィング ステータス
test STALE (最後の実行: パス、終了 0 — コードが変更されました)
すべてのチェック結果は、コードの正確な状態のフィンガープリント、つまり git コミットに加えて、ステージングされているかどうか、追跡されているかどうかなど、コミットされていないすべてのハッシュにバインドされます。 1 バイトを変更すると、結果は自動的に STALE に切り替わります。誰も「このパスはまだ有効ですか?」と尋ねるのを忘れる必要はありません。
カール -L https://github.com/yebiguo/proofrun/releases/download/v0.2.0/proofrun_linux_amd64.tar.gz |タールxz
# 他のプラットフォーム: https://github.com/yebiguo/proofrun/releases
またはソースからビルドします。
github.com/yebiguo/proofrun/cmd/proofrun@latest をインストールしてください
クイックスタート
proofrun init # .proofrun.yml を書き込みます
proofrun run test -- pytest # 実際に pytest を実行し、リソースをバインドします

ウルトラ
proofrun status --strict # ゼロ以外の値が PASS でない場合は終了します。
エージェントを信頼しているだけではなく、なぜそうなるのか
LLM コールはどこにもありません。 ProofRun は AI を検証するために AI を使用しません。実際のサブプロセスを開始し、実際の終了コードを読み取ります。これがメカニズム全体です。
4 つのステータス。決して推測ではありません。 PASS 、 FAIL 、 STALE 、 NOT RUN — それぞれは観察された実行、または文書化された実行の欠如に基づいています。 5 番目の「おそらく大丈夫」はありません。
完全にオフライン。ネットワーク通話、テレメトリ、アカウントはゼロです。
引数は正確ですが、文字列は一致しません。 pytest -k "foo bar" として宣言されたチェックは、テキストに平坦化されると単に似ているだけのコマンドでは満たされません。ProofRun は文字列ではなく実引数の配列を比較します。
ProofRun が意図的に行わないこと
テスト出力は解析されず、コードの品質も判断されず、何も自動修正されません。完全な境界については、AGENTS.md を参照してください。
AI エージェントによって構築され、AI エージェントが責任を負う
ProofRun は、人間の指示の下、AI コーディング エージェント (Claude Code) によって作成され、最初のリリース前に独立した読み取り専用の敵対的レビューが数回行われました。このレビューでは、ProofRun 自体のコマンド比較がだまされる可能性があることがわかりました。シェル引数の引用が間違っているため、チェックはサイレントにテストを実行せず、それでも PASS を報告します。完全な再現、正確な修正、および単純なパッチでは不十分な理由 → docs/case-study.md 。
すべての修正は、妥当性を確認するだけでなく、受け入れられる前に実際の複製に対して検証されました。 AI エージェントに責任を負わせるために構築されたツールは、それ自体に適用される同じ監視に耐えることができなければ、ビジネスとして成り立ちません。
proofrun init # .proofrun.yml を生成する
proofrun run < check-name > -- < cmd > # 実際に <cmd> を実行し、終了コード + 期間を現在の git 状態にバインドします
proofrun run-all [--only < name > ] # 12 月ごとに実行

完了したチェック、各結果の保存
proofrun ステータス [--strict] # チェックごとに PASS / FAIL / STALE / NOT RUN。 --strict は、必要なチェックが PASS でない場合、ゼロ以外で終了します。
proofrun report [--json] # 完全なレポート、人間または機械が判読可能
構成: .proofrun.yml
チェック:
テスト:
コマンド: [pytest]
必須 : true
ビルド:
コマンド: [npm、実行、ビルド]
必須 : true
糸くず:
コマンド: [ラフ、チェック、.]
必須 : false
コマンドはシェル文字列ではなく argv リストです。ProofRun はシェルを経由することはありません。実際に実行された内容と宣言された内容の比較は、要素ごとに正確である必要があります。 required: true は、チェック ブロックのステータスを --strict にするもので、これをコミット前のフックまたは CI ゲートに接続します。
すべての結果は、現在の git HEAD にバインドされ、追跡されていない無視されていないファイルの内容と結合された git diff HEAD の SHA-256 ハッシュにバインドされます。 proofrun ステータスは毎回そのフィンガープリントを再計算し、ローカルに保存されているものと比較します。変更された 1 つのスペースまたは 1 つの新しいファイルに至るまでの不一致は、 STALE を報告します。
上: pull_request
権限:
内容：読む
ジョブ:
確認します:
実行: ubuntu-最新
手順:
- 使用: yebiguo/proofrun@v1
これは、正確な PR ヘッド コミットを独自にチェックアウトします。呼び出し元のワークフローがすでにチェックアウトしているものは決して信頼しないため、pull_request トリガーは、代わりに GitHub の合成マージ プレビュー コミットを黙って渡すことはできません。次に、PR ブランチに入ってきたすべての Recipe.json をクリアし、チェックサム検証済みのproofrun バイナリをダウンロードし、proofrun status --strict でゲートする前に実際にproofrun run-allを実行します。 PR ブランチにチェックインされたレシートについては何も信頼されません。ゲートが確認するすべての結果は、この実行によって生成されたものです。
既知の制限: これは .proofrun.yml 自体を sam による弱体化から保護するものではありません

コードを変更する PR — PR はチェックのコマンドを緩めるか削除することができ、アクションはより弱いバージョンを忠実に再実行します。 .proofrun.yml が PR のベース ブランチと異なる場合は (ビルド アノテーションを介して) 警告しますが、それについてはブロックしません。変更の他の部分を確認するのと同じ方法で、その差分を確認してください。
v0.3 — 一般的なテスト ランナー (pytest、Jest、JUnit) の構造化出力サポート
署名済みの改ざん防止領収書が注目されていますが、まだデザインされていません
コードを変更する同じ PR 内で .proofrun.yml 自体が弱体化しないように保護します (現在は警告のみでブロックされていません。上記の「既知の制限事項」を参照)
問題やPRを歓迎します。これは、1.0 以前の若いプロジェクトで、対象範囲が狭く、意図的に行われています。STALE 検出や受信スキーマに関わるものを提案する前に、AGENTS.md を参照してください。これらは、このプロジェクトが最も間違えてはならない部分です。
AI コーディング エージェントのローカル検証レシート
Readme MIT ライセンス アクティビティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A local verification receipt for AI coding agents. Contribute to yebiguo/ProofRun development by creating an account on GitHub.

GitHub - yebiguo/ProofRun: A local verification receipt for AI coding agents · GitHub
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
yebiguo
/
ProofRun
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
37 Commits 37 Commits .github/ workflows .github/ workflows cmd/ proofrun cmd/ proofrun docs docs internal internal .gitignore .gitignore .goreleaser.yml .goreleaser.yml .proofrun.yml .proofrun.yml AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md action.yml action.yml go.mod go.mod go.sum go.sum View all files Repository files navigation
A local verification receipt for AI coding agents.
ProofRun doesn't judge whether your code is correct. It proves — cryptographically, not by asking nicely — which checks actually ran against the exact code you have right now.
An AI coding agent says "all tests pass." Is that true?
Maybe. It was true the last time the agent actually ran the tests. But that might have been three edits ago. The agent might not even remember running them — it might just be inferring "the change looks right, tests probably still pass." From the words alone, you have no way to tell "I ran it and it passed" apart from "I'm pretty sure it would pass."
ProofRun closes that gap. Not by making the agent more honest — by making the claim itself checkable.
$ proofrun run test -- pytest
...
test: pass (exit 0, 1841ms)
$ proofrun status
test PASS (exit 0, 1841ms)
# code changes after this point — agent or human, doesn't matter
$ proofrun status
test STALE (last run: pass, exit 0 — code changed since)
Every check result is bound to a fingerprint of your exact code state: the git commit, plus a hash of everything uncommitted — staged or not, tracked or not. Change a single byte, and the result flips to STALE automatically. Nobody has to remember to ask "does this PASS still count?"
curl -L https://github.com/yebiguo/proofrun/releases/download/v0.2.0/proofrun_linux_amd64.tar.gz | tar xz
# other platforms: https://github.com/yebiguo/proofrun/releases
Or build from source:
go install github.com/yebiguo/proofrun/cmd/proofrun@latest
Quick start
proofrun init # writes .proofrun.yml
proofrun run test -- pytest # runs pytest for real, binds the result
proofrun status --strict # non-zero exit if anything isn't PASS
Why this, not just trusting the agent
No LLM calls, anywhere. ProofRun doesn't use AI to verify AI. It starts a real subprocess and reads its real exit code — that's the entire mechanism.
Four statuses, never a guess. PASS , FAIL , STALE , NOT RUN — each one comes from an observed execution, or the documented absence of one. There's no fifth "probably fine."
Fully offline. Zero network calls, zero telemetry, zero accounts.
Argv-exact, not string-matched. A check declared as pytest -k "foo bar" can't be satisfied by a command that merely looks similar once flattened to text — ProofRun compares real argument arrays, not strings.
What ProofRun deliberately does not do
It does not parse test output, does not judge code quality, and does not auto-fix anything. See AGENTS.md for the complete boundary.
Built by an AI agent, held accountable by one
ProofRun was written by an AI coding agent (Claude Code) under human direction, then went through several rounds of independent, read-only adversarial review before the first release. That review found that ProofRun's own command comparison could be tricked: a misquoted shell argument made a check silently run zero tests and still report PASS . Full repro, the exact fix, and why a simple patch wasn't enough → docs/case-study.md .
Every fix was verified against a real reproduction before being accepted — not just reviewed for plausibility. A tool built to hold AI agents accountable has no business existing if it can't survive that same scrutiny applied to itself.
proofrun init # generate .proofrun.yml
proofrun run < check-name > -- < cmd > # run <cmd> for real, bind exit code + duration to current git state
proofrun run-all [--only < name > ] # run every declared check, saving a result after each one
proofrun status [--strict] # PASS / FAIL / STALE / NOT RUN per check; --strict exits non-zero if a required check isn't PASS
proofrun report [--json] # full report, human- or machine-readable
Config: .proofrun.yml
checks :
test :
command : [pytest]
required : true
build :
command : [npm, run, build]
required : true
lint :
command : [ruff, check, .]
required : false
command is an argv list, not a shell string — ProofRun never goes through a shell, and comparing what actually ran against what's declared has to be exact, element for element. required: true is what makes a check block status --strict , which is what you'd wire into a pre-commit hook or CI gate.
Every result is bound to your current git HEAD plus a SHA-256 hash of git diff HEAD combined with the contents of any untracked, non-ignored files. proofrun status recomputes that fingerprint every time and compares it against what's stored locally — any mismatch, down to a single changed space or one new file, reports STALE .
on : pull_request
permissions :
contents : read
jobs :
verify :
runs-on : ubuntu-latest
steps :
- uses : yebiguo/proofrun@v1
This does its own checkout of the exact PR head commit — it never trusts whatever the calling workflow already checked out, so a pull_request trigger can't silently hand it GitHub's synthetic merge-preview commit instead. It then clears out any receipt.json that came in on the PR branch, downloads a checksum-verified proofrun binary, and runs proofrun run-all for real before gating on proofrun status --strict . Nothing about a receipt checked into the PR branch is ever trusted — every result the gate sees was produced by this run.
Known limitation: this does not protect .proofrun.yml itself from being weakened by the same PR that changes the code — a PR could loosen or remove a check's command and the Action would faithfully re-run the weaker version. It warns (via a build annotation) when .proofrun.yml differs from the PR's base branch, but it does not block on that; review that diff the same way you'd review any other part of the change.
v0.3 — structured output support for common test runners (pytest, Jest, JUnit)
Signed, tamper-evident receipts are on the radar, not yet designed
Protecting .proofrun.yml itself from being weakened within the same PR that changes the code (currently only warned about, not blocked — see "Known limitation" above)
Issues and PRs welcome. This is a young, pre-1.0 project with a narrow, deliberate scope — see AGENTS.md before proposing anything that touches STALE detection or the receipt schema; those are the parts this project can least afford to get wrong.
A local verification receipt for AI coding agents
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
