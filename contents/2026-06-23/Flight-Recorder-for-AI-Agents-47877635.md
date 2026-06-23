---
source: "https://stordai.com/"
hn_url: "https://news.ycombinator.com/item?id=48650789"
title: "Flight Recorder for AI Agents"
article_title: "Stord — The Undo Button for AI Agents"
author: "ishwantsingh"
captured_at: "2026-06-23T20:42:01Z"
capture_tool: "hn-digest"
hn_id: 48650789
score: 1
comments: 1
posted_at: "2026-06-23T20:21:13Z"
tags:
  - hacker-news
  - translated
---

# Flight Recorder for AI Agents

- HN: [48650789](https://news.ycombinator.com/item?id=48650789)
- Source: [stordai.com](https://stordai.com/)
- Score: 1
- Comments: 1
- Posted: 2026-06-23T20:21:13Z

## Translation

タイトル: AI エージェント用のフライト レコーダー
記事のタイトル: Stord — AI エージェントの「元に戻す」ボタン
説明: AI エージェントが読み書きするすべてのファイル (エージェントに起因) を記録し、すべての変更を元に戻せるオープンソース CLI。クロード コード、コーデックスなどで動作します。 npm install -g ストレージ。

記事本文:
Stord — AI エージェントの元に戻すボタン Stord ★ GitHub 参加待機リストの書き込み · rev 41 保存 ✗ ブロック · スコープ外 ✓ ロールバック · 2 ファイルの読み取り · ログに記録 オープン ソース · MIT · クロード コード、コーデックスなどと連携 AI エージェントの元に戻すボタン
権限をスキップしてエージェントを実行できるようにします。 Stord は、読み取りおよび書き込みを行ったエージェントに起因するすべてのファイルを記録し、すべての変更を元に戻せるようにします。
すべての操作を記録に残す
エージェントは .env を読み取り、ビルドを rm -rf します。 Stord は、誰が何をしたかを正確に表示します。その後、Stord Undo ですべてを元に戻します。
エージェントは本番環境にアクセスし、シートベルトを着用せずに行動します。
コーディング エージェントはスキップ権限フラグを使用して実行されます。運用エージェントは、どの従業員よりも幅広い資格情報を保持しています。構成を上書きしたり、ディレクトリを消去したり、ファイルを漏洩したりすると、何が起こったのか記録が残らず、元に戻すこともできません。
ID ツールはエージェントが誰であるかを認識します。ゲートウェイはそれが何と呼ばれているかを知っています。データに何が起こったのかは誰も知りませんし、元に戻すこともできません。
エージェントとファイル間のガバナンス層。
今すぐオープンソース CLI を始めましょう。依存関係を持たずに、自分のマシンで監査および元に戻すことができます。クラウド層は、強制的なアクセス許可と改ざんを明示するフリート全体の監査を追加します。それが待機リストの目的です。
正確なエージェント、セッション、およびユーザーに起因する、すべての読み取り、書き込み、および削除の追加専用ジャーナル。 stord ログにはタイムラインが表示されます。 stord エージェントは、誰が何を触ったかを示します。
すべての書き込みは、実行前にスナップショットが作成されます。 rm -rf の後であっても、 stord undo を使用して 1 つの操作を元に戻すか、 stord list を使用してツリー全体をロールバックします。実際の git 履歴には決して触れません。
最小特権キーはチームごとではなくエージェントごとに。 write: /build/** などのパススコープのアクセス許可が強制され、さらに改ざん防止のオフホスト監査とフリート ダッシュボードが追加されます。

。ウェイティングリストに参加する →
1 つのノード ファイル、依存関係なし、MIT ライセンス。そのままの状態で Claude Code と連携し、フックを介して Codex、Hermes、OpenClaw、または任意のカスタム エージェントと連携します。
# エージェントが作業する任意のプロジェクト内
$ ストア有効
$ claude --dangerously-skip-permissions
$ stord log # 触れたすべてのもの
$ stord 最後に元に戻す # 元に戻す クラウド — 早期アクセス
艦隊を運営しましょう。領収書は保管しておいてください。
CLI は今日あなたのものになります。現在、クラウド層 (エージェントごとの強制的なアクセス許可、改ざん防止オフホスト監査、すべてのエージェントとマシンにわたるダッシュボード) が初期チームに導入されています。早期アクセスと創設ユーザー価格で参加してください。
スパムなし - アップデートのみを起動します。

## Original Extract

Open-source CLI that records every file your AI agents read and write — attributed to the agent — and makes every change reversible. Works with Claude Code, Codex, and more. npm install -g stordai.

Stord — The Undo Button for AI Agents Stord ★ GitHub Join waitlist write · rev 41 saved ✗ blocked · outside scope ✓ rolled back · 2 files read · logged Open source · MIT · works with Claude Code, Codex & more The Undo Button for AI Agents
Let your agents run with permissions skipped. Stord records every file they read and write — attributed to the agent that did it — and makes every change reversible.
Every Operation, On the Record
An agent reads your .env and rm -rf s your build. Stord shows you exactly who did what — then stord undo brings it all back.
Agents act with production access and no seatbelt.
Coding agents run with skip-permissions flags. Ops agents hold credentials broader than any employee’s. When one overwrites a config, wipes a directory, or leaks a file, there’s no record of what happened and no way back.
Identity tools know who the agent is. Gateways know what it called . Nobody knows what happened to your data — or can undo it.
A governance layer between agents and your files.
Start with the open-source CLI today — audit and undo on your own machine, zero dependencies. The cloud layer adds enforced permissions and tamper-evident, fleet-wide audit. That’s what the waitlist is for.
An append-only journal of every read, write, and delete — attributed to the exact agent, session, and user. stord log shows the timeline; stord agents shows who touched what.
Every write is snapshotted before it runs. Reverse a single operation with stord undo or roll the whole tree back with stord restore — even after rm -rf . Your real git history is never touched.
Least-privilege keys per agent, not per team. Path-scoped permissions like write: /build/** , enforced — plus tamper-evident, off-host audit and a fleet dashboard. Join the waitlist →
One Node file, zero dependencies, MIT licensed. Works with Claude Code out of the box, and with Codex, Hermes, OpenClaw, or any custom agent via hooks.
# in any project your agents work in
$ stord enable
$ claude --dangerously-skip-permissions
$ stord log # everything it touched
$ stord undo last # take it back Cloud — early access
Run your fleet. Keep the receipts.
The CLI is yours today. The cloud layer — enforced per-agent permissions, tamper-evident off-host audit, and a dashboard across every agent and machine — is onboarding early teams now. Join for early access and founding-user pricing.
No spam — launch updates only.
