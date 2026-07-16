---
source: "https://heimwall.ai/individual"
hn_url: "https://news.ycombinator.com/item?id=48939783"
title: "Show HN: HeimWall – Catch secrets before you paste them into Claude or Cursor"
article_title: "The free Mac app | HeimWall"
author: "ata2gxxx"
captured_at: "2026-07-16T20:53:02Z"
capture_tool: "hn-digest"
hn_id: 48939783
score: 2
comments: 0
posted_at: "2026-07-16T20:27:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HeimWall – Catch secrets before you paste them into Claude or Cursor

- HN: [48939783](https://news.ycombinator.com/item?id=48939783)
- Source: [heimwall.ai](https://heimwall.ai/individual)
- Score: 2
- Comments: 0
- Posted: 2026-07-16T20:27:20Z

## Translation

タイトル: HN を表示: HeimWall – クロードまたはカーソルに貼り付ける前に秘密をキャッチ
記事のタイトル: 無料の Mac アプリ |ハイムウォール
説明: Mac から秘密が流出する前に秘密をキャッチします。 AI プロンプト内のシークレットと PII にフラグを付ける、個人開発者向けの無料のメニュー バー アプリです。100% オンデバイス、ローカルのみ、アカウントは必要ありません。入力した内容がマシンから離れることはありません。

記事本文:
無料の Mac アプリ | HeimWall コンテンツへスキップ Heim Wall ホーム 仕組み 無料アプリ 価格 ブログ よくある質問 エンジニア向け チーム向け 無料 Mac アプリ · 100% オンデバイス Mac から流出する前にその秘密をキャッチしましょう。
個人開発者向けの無料のメニューバー アプリ。 AI プロンプト (カーソル、クロード コード、およびコパイロット) 内のシークレットと PII にフラグをすべて Mac 上で付けます。アカウントも組織も、どこにも送信されません。
このアプリはどこにも何も送信しません。作成するアカウントも、電話をかけるサーバーもありません。キャプチャ、検出、アラートはすべて Mac 上で行われます。ネットワーク ケーブルを引き抜くと、まったく同じように動作します。
貼り付けるつもりはなかったもの。
.env ファイルまたはログから誤って 1 回貼り付けるだけで十分です。アプリは、ユーザーがプロンプトを作成するときにプロンプ​​トを監視し、出力すべきではないものにフラグを立てます。
AWS アクセス キー、GitHub トークン、秘密キー、JWT、データベース URL — 明確な形状を持つ認証情報に関する 25 以上の手書きのルール。シークレットがプロンプトに入力される最も一般的な方法は、.env ファイルから貼り付けることです。これはソースでそれをキャッチします。
デバッグのために Cursor にドロップしたログまたはスタック トレースに埋め込まれた SSN、クレジット カード番号、電子メール。プロンプトが送信される前に、即座にフラグが立てられます。
カーソル、クロード コード、コパイロット — 1 つのメニュー バー エージェントがコンポーザー自体を読み取るため、ツールが独自のクラウドと直接通信している場合でもプロンプトが表示されます。ツールごとのセットアップはありません。
面倒な作業は決定論的な正規表現であり、即時、説明可能、オフラインで実行されます。小さな事前トレーニング済み分類器が、ファジー リークのフォールバックとしてデバイス上で実行されます。いずれにせよ、プロンプトは収集されないため、プロンプトでモデルがトレーニングされることはありません。
プロンプト、編集されたスニペット、カウントではありません。ループ全体があなたのマシン上に存在します。これは私たちが監査した主張であり、スローガンではありません。
プロンプトを作成したときにプロンプ​​トを読み取ります — 開発上

ネットワークの前の氷。
完全にオフラインで、50 ミリ秒未満でシークレットと PII をスキャンします。
ソースでふんわりトースト。何も保存されず、何も送信されません。
プロンプト スキャン トーストがメモリから削除されました
サインインする必要はありません。アプリには送信するものがないため、データを受信するバックエンドはありません。
~15 MB、Rust + Tauri、アイドル状態の RAM は 150 MB 未満。完全にオフラインで実行されます。ネットワークには一切触れません。
同じエンジンです。こちらは何も送信しません。
チーム製品はまったく同じオンデバイス エージェントであり、組織に登録されているため、マネージャーは編集されたシグナル (カウントとマスクされたスニペット) を確認できますが、プロンプトは表示されません。無料アプリはそのすべてをスキップします。登録されていないため、アップリンクはまったくありません。
アカウントも組織もネットワークもありません。あなたは入力します。それは捕まえます。生のバイトはメモリから削除されます。何も記録されたり送信されたりすることはありません。永久無料。
同じエージェントが 1 つのコードで組織に参加しています。編集されたメタデータをアップリンクするので、管理者は安全スコアと傾向を確認できます。生のプロンプトがデバイスから出ることはまだありません。チーム向け→
キャプチャ、検出、秘匿化パイプラインが内部でどのように機能するか興味がありますか?完全なアーキテクチャを参照 →
永久無料、アカウントなし。 Apple によって署名および公証されているため、通常のダブルクリックで開きます。クリックスルーするためのセキュリティ警告はありません。
Mac 用ダウンロード v 0.0.5 · 3.9 MB · Apple Silicon · macOS 13+ 署名および公証済み · 100% オンデバイス · アカウントなし。
HeimWall をアプリケーション フォルダーにドラッグします。
一度起動してアクセシビリティを許可すると、AI ツールのプロンプト ボックスを読み取ることができます。 Mac から何も離れることはありません。
ウィンドウには「保護されました」と表示されます。 Cursor に偽のキーを入力し、デバイス上でフラグが設定され、編集されるのを確認します。
Intel Mac または Windows を使用していますか?ビルドの準備ができたら通知を受け取ります。
エージェントの従業員に対するマネージャー優先の可観測性。エンジニアが Cursor、Claude Code、Copi をどのように使用しているかをご覧ください。

ロット - プロンプトを読まずに、漏洩した秘密と PII を捕捉します。
HeimWall はパフォーマンスの評価を目的としたものではありません。契約により、安全スコアとフラグ履歴は運用上の信号であり、パフォーマンスの信号ではありません。これらは、昇進、報酬、または解雇の決定に使用することはできません。

## Original Extract

Catch the secret before it leaves your Mac. A free menu-bar app for individual developers that flags secrets and PII in your AI prompts — 100% on-device, local-only, no account. Nothing you type ever leaves the machine.

The free Mac app | HeimWall Skip to content Heim Wall Home How it works Free app Pricing Blog FAQ For engineers For teams Free Mac app · 100% on-device Catch the secret before it leaves your Mac.
The free menu-bar app for individual developers. It flags secrets and PII in your AI prompts — in Cursor, Claude Code, and Copilot — entirely on your Mac. No account, no org, nothing sent anywhere.
This app sends nothing, anywhere. There is no account to create and no server to phone home to. Capture, detection, and the alert all happen on your Mac. Pull the network cable and it works exactly the same.
The stuff you never meant to paste.
One accidental paste from a .env file or a log is all it takes. The app watches the prompt as you write it and flags the things that shouldn’t go out.
AWS access keys, GitHub tokens, private keys, JWTs, database URLs — 25+ hand-written rules for the credentials with hard shapes. The most common way a secret enters a prompt is a paste from a .env file; this catches it at the source.
SSNs, credit-card numbers, and emails buried in a log or a stack trace you dropped into Cursor to debug. Flagged instantly, before the prompt is sent.
Cursor, Claude Code, Copilot — one menu-bar agent reads the composer itself, so it sees the prompt even when the tool talks straight to its own cloud. No per-tool setup.
The heavy lifting is deterministic regex — instant, explainable, and offline. A small pretrained classifier runs on-device as a fallback for fuzzier leaks. Either way, no model is ever trained on your prompts, because they are never collected.
Not the prompt, not a redacted snippet, not a count. The entire loop lives on your machine — this is a claim we’ve audited, not a slogan.
Reads the prompt as you write it — on-device, before the network.
Scans for secrets and PII in under 50ms, entirely offline.
A soft toast at the source. Nothing is stored, nothing is sent.
prompt scan toast dropped from memory
You never sign in. There is no backend to receive your data because the app has nothing to send.
~15 MB, Rust + Tauri, under 150 MB RAM idle. Runs fully offline — the network is never touched.
Same engine. This one sends nothing.
The team product is the exact same on-device agent, enrolled in an organization so a manager can see redacted signal — counts and masked snippets, never prompts. The free app skips all of that: it is not enrolled, so there is no uplink at all.
No account, no org, no network. You type; it catches; the raw bytes are dropped from memory. Nothing is ever recorded or sent. Free forever.
The same agent, joined to an org with one code. It uplinks redacted metadata so a manager sees a Safety Score and trends. Raw prompts still never leave the device. For teams →
Curious how the capture, detection, and redaction pipeline works under the hood? See the full architecture →
Free forever, no account. Signed and notarized by Apple, so it opens with a normal double-click — no security warnings to click through.
Download for Mac v 0.0.5 · 3.9 MB · Apple Silicon · macOS 13+ Signed & notarized · 100% on-device · no account, ever.
Drag HeimWall into your Applications folder.
Launch it once and allow Accessibility, so it can read the prompt box of your AI tools. Nothing ever leaves your Mac.
The window shows Protected. Type a fake key into Cursor and watch it get flagged, redacted, on-device.
On an Intel Mac or Windows? Get notified when your build is ready.
Manager-first observability for the agentic workforce. See how your engineers use Cursor, Claude Code, and Copilot — catch leaked secrets and PII, without reading their prompts.
HeimWall is not for performance reviews. By contract, safety scores and flag history are operational signal, not performance signal. They may not be used in promotion, compensation, or termination decisions.
