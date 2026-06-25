---
source: "https://www.vibedrift.ai"
hn_url: "https://news.ycombinator.com/item?id=48675318"
title: "Show HN: VibeDrift – measuring AI coding drift across open source repos"
article_title: "VibeDrift: Stop your AI agent from getting worse as your project grows"
author: "samiahmadkhan"
captured_at: "2026-06-25T16:46:01Z"
capture_tool: "hn-digest"
hn_id: 48675318
score: 4
comments: 0
posted_at: "2026-06-25T15:56:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: VibeDrift – measuring AI coding drift across open source repos

- HN: [48675318](https://news.ycombinator.com/item?id=48675318)
- Source: [www.vibedrift.ai](https://www.vibedrift.ai)
- Score: 4
- Comments: 0
- Posted: 2026-06-25T15:56:17Z

## Translation

タイトル: Show HN: VibeDrift – オープンソース リポジトリ全体の AI コーディング ドリフトの測定
記事のタイトル: VibeDrift: プロジェクトの成長に伴う AI エージェントの劣化を防ぐ
説明: プロジェクトが成長するにつれて、AI エージェントは悪化します。 VibeDrift はドリフトを捕捉して修正し、エージェントにドリフトを停止するためのコンテキストを提供します。

記事本文:
VibeDrift: プロジェクトの成長に伴って AI エージェントが悪化するのを防ぎます。 ★ 新しい VibeDrift がオープンソースになりました。無料で完全に公開されており、その仕組みを正確にご覧ください。 GitHub でスターを付けてください → VIBEDRIFT 仕組み ドキュメント ブログ 価格 サインイン 始めましょう、無料 プロジェクトが成長するにつれて、AI エージェントは悪化します。
パターンを忘れ、コードを繰り返し、機能していたものを壊してしまいます。 VibeDrift はドリフトが広がる前にそれをキャッチします。
GitHub でのスター セッション全体のエージェント品質 VibeDrift あり ドリフトなし 92 · A 41 · D セッション 5 セッション 50 クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント クロード コード カーソル GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro 任意の MCP クライアント 同じプロジェクト。 2 つの異なるスタイル。
AI は 2 つのセッションで両方のファイルを書き込みました。どちらもリンターを通過します。彼らは一緒に、あなたのプロジェクトにはルールがないことを教えてくれます。
'../repositories/user' から { UserRepository } をインポートします。
import { requireAuth } から '../middleware/auth' ;
非同期関数 getUser をエクスポート (必須: リクエスト) {
requireAuth (要求);
const user = await UserRepository .fi

ndById(id);
if (!user) throw new NotFoundError ();
リターンユーザー。
✓ リポジトリ パターンを使用します ✓ ユーザーがログインしていることを確認します ✓ 適切な型指定エラー order_handler.ts import { db } from '../config/database' ;
非同期関数 getOrder (req: リクエスト) をエクスポート {
const id = req.params.id;
const rows = await db .query( 'SELECT * FROM 注文...' );
if (行の長さ === 0) {
return { ステータス: 404、エラー: '見つかりません' };
}
行[0]を返します。
} ✗ 生の SQL、リポジトリなし ✗ 誰でも呼び出すことができ、認証なし ✗ ランダム エラー形式 8 チェック ~2 秒でローカルで実行 // 堀: インループ MCP エージェントは書き込む前にチェックします。
着陸後の漂流についてはレポートでお知らせします。 VibeDrift MCP サーバーは、ランディングをまったく停止します。タスクの途中で、Claude Code と Cursor がコードベースにすでに何を行っているかを尋ね、1 秒以内に証拠に裏付けられた答えが得られます。これにより、ドリフト検出がドリフト防止に変わります。
見つけて、直して、きれいに保ちましょう。
1 回のスキャンで問題の程度を痛みの度合いでランク付けし、エージェントにコピー＆ペーストの修正を渡し、セッションごとに読み取った記憶を残します。
エージェントがセッションごとに読み取る .vibedrift/ フォルダー。
エージェントはタスクの途中でコードベースを尋ねます。1 秒以内に答えます。
保存するたびにメモリを更新します。ネットワーク通話ゼロ。
VibeDrift / VibeDrift TypeScript ★ 4 ⚖ MIT Star · 4 npm // Proof 動作確認しました。正直に。
独立したパスによってグレーディングされた VibeDrift の有無にかかわらず、モデルが間違っていた場合にのみ状況が変わり、何も起こらなかった部分が示されます。
いずれにせよ、モデルが規約に一致する場合、シグナルは何も行いませんが、それを隠しているわけではありません。
パターンは目に見えるドリフトを捉えます。ディープは彼らができないものをキャッチします。
無料のローカル スキャンにより、明らかなドリフトを数秒で検出します。ディープスキャンはコードが実際に何を意味するのかを読み取るため、より深い競合を迅速に検出します。

スキャンすることはできません。大きなコードベースに静かに積み重なっていくようなものです。
✓ 静的 + コード DNA 分析、パターン、構造、重複チェックをすべてマシン上で約 2 秒で実行
✓ ドリフト検出器、不一致パターン、認証ガードの欠落
✓ 足場の検出、中途半端なコピー＆ペーストされたコードのフラグ
✓ スキャンオーバースキャンデルタ、最後のスキャン以降に変更された内容
✓ CI 対応、あらゆる保存およびプルリクエストに十分な速度
コードの動作が明らかに異なっていることをキャッチします。
★ セマンティックな重複、同じロジックが別の名前で隠されている
★ 関数名がその機能と一致しない場合のインテント不一致検出
★ AI によって検証された結果、すべての結果はクロードによって二重チェックされるため、誤報を追う必要はありません
★ コヒーレンスレポート、コード自体との一貫性を最悪から順にランク付けしたスコアカード
★ 修正プロンプト、各検出結果に対するすぐに貼り付けられる修正
★ MCP 経由のエディタ内で、エージェントが問題を作成するときに問題をすぐにキャッチします。
目では見つけられない競合、つまりコードの意味に隠れている競合を捉えます。
ディープ スキャンは、AI エージェントが独自に規則を認識できない場合、つまりドリフトが忍び寄る場合に最も役立ちます。
無料で始められます。ディープスキャンに対してのみ料金を支払います。
ローカル スキャンとローカル MCP ツールは、どのプランでも永久に無料です。有料プランでは、予算の詳細なスキャンとエディタ内チェックが追加されます。
ディープスキャンが終了しましたか?どのプランでも、5 つまで $10 でチャージできます。クレジットに有効期限はありません。
npx @vibedrift/cli を実行します。どのプロジェクトでも、アカウントや設定は必要ありません。ローカル スキャンは無料で無制限です。ディープ スキャンまたはダッシュボードが必要な場合にのみサインインします。
両方。スキャンごとに、コピー可能な修正プロンプトが .vibedrift/fix-plan.md に書き込まれるため、ユーザーまたはエージェントはそれらを適用できます。ディープスキャンにより、パターンでは認識できない矛盾に対して AI が合成した修正が追加されます。
5 つのカテゴリにわたる 0 ～ 100 のスコア、再採点

一般的なルールブックではなく、独自のコードベースの主要なパターンを使用してください。プロジェクトの内部一貫性がどの程度あるのか、どのファイルがプロジェクトを下に引っ張っているのかがわかります。
はい。ローカル スキャナーは、github.com/VibeDrift/VibeDrift で MIT ライセンスを取得しています。
コードベースが AI に何を教えてきたかを確認します。
GitHub にスターを付ける スキャンする準備ができていませんか?更新の通知を受け取ります。
サポートとアップデートについては Discord に参加してください → VibeDrift AI が生成したコード内のアーキテクチャのドリフトを拡散する前にキャッチします。無料のオープンソースで、お使いのマシン上で実行できます。

## Original Extract

Your AI agent gets worse as your project grows. VibeDrift catches the drift, fixes it, and gives your agent the context to stop making it.

VibeDrift: Stop your AI agent from getting worse as your project grows ★ New , VibeDrift is now open source . See exactly how it works, free to use and fully public. Star us on GitHub → VIBEDRIFT How It Works Docs Blog Pricing Sign in Get started, free Your AI agent gets worse as your project grows.
It forgets your patterns, repeats code, and breaks what worked. VibeDrift catches the drift before it spreads.
Star on GitHub Agent quality over sessions With VibeDrift Without the drift 92 · A 41 · D session 5 session 50 Plays nice with Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Claude Code Cursor GitHub Copilot Windsurf Antigravity OpenAI Codex Kiro Any MCP client Same project. Two different styles.
Your AI wrote both files, in two sessions. Both pass the linter. Together they teach it your project has no rules.
import { UserRepository } from '../repositories/user' ;
import { requireAuth } from '../middleware/auth' ;
export async function getUser (req: Request) {
requireAuth (req);
const user = await UserRepository .findById(id);
if (!user) throw new NotFoundError ();
return user;
} ✓ Uses the repository pattern ✓ Checks the user is logged in ✓ Proper typed error order_handler.ts import { db } from '../config/database' ;
export async function getOrder (req: Request) {
const id = req.params.id;
const rows = await db .query( 'SELECT * FROM orders...' );
if (rows.length === 0) {
return { status: 404, error: 'not found' };
}
return rows[0];
} ✗ Raw SQL, no repository ✗ Anyone can call it, no auth ✗ Random error format 8 checks ~2s runs locally // The moat: in-loop MCP Your agent checks before it writes.
A report tells you about drift after it lands. The VibeDrift MCP server stops it from landing at all. Mid-task, Claude Code and Cursor ask your codebase what it already does, and get an evidence-backed answer in under a second. That turns drift detection into drift prevention .
Find it, fix it, keep it clean.
One scan ranks what's wrong by how much it hurts, hands your agent a copy-paste fix, and leaves a memory it reads every session.
A .vibedrift/ folder your agent reads every session.
Your agent asks your codebase mid-task, answer in under a second.
Refreshes the memory on every save. Zero network calls.
VibeDrift / VibeDrift TypeScript ★ 4 ⚖ MIT Star · 4 npm // Proof We checked if it works. Honestly.
With and without VibeDrift, graded by an independent pass, it only changes things when the model would've gotten it wrong, and we show where it did nothing.
When the model would match your conventions anyway, the signal does nothing, and we're not hiding that.
Patterns catch the visible drift. Deep catches what they can't.
The free local scan catches the obvious drift in seconds. The deep scan reads what your code actually means , so it finds the deeper conflicts a quick scan can't, the kind that quietly pile up in big codebases.
✓ Static + Code DNA analysis , pattern, structure, and duplicate checks, all on your machine in ~2s
✓ Drift detectors , mismatched patterns and missing auth guards
✓ Scaffolding detection , flags half-finished, copy-pasted code
✓ Scan-over-scan delta , what changed since your last scan
✓ CI-ready , fast enough for every save and pull request
Catches the stuff your code visibly does differently.
★ Semantic duplicates , the same logic hiding under different names
★ Intent mismatch detection , when a function's name doesn't match what it does
★ AI-validated findings , every result is double-checked by Claude, so you don't chase false alarms
★ Coherence report , a ranked scorecard of how consistent your code is with itself, worst first
★ Fix prompts , a ready-to-paste fix for each finding
★ In-editor via MCP , catches issues right as your agent writes them
Catches the conflicts you can't spot by eye, the ones hiding in what the code means.
The deep scan helps most exactly where your AI agent can't see your conventions on its own, which is right when drift creeps in.
Free to start. Pay only for deep scans.
Local scans and the local MCP tools are free on every plan, forever . Paid plans add a deep-scan budget and in-editor checks.
Out of deep scans? Top up 5 for $10 , on any plan. Credits never expire.
Run npx @vibedrift/cli . in any project, no account, no config. Local scans are free and unlimited; sign in only when you want a deep scan or the dashboard.
Both. Every scan writes copy-ready fix prompts to .vibedrift/fix-plan.md so you, or your agent, can apply them. Deep scans add AI-synthesized fixes for the contradictions patterns can't see.
A 0–100 score across five categories, graded against your own codebase's dominant patterns, not a generic rulebook. It tells you how internally consistent your project is and which files drag it down.
Yes. The local scanner is MIT-licensed at github.com/VibeDrift/VibeDrift.
See what your codebase has been teaching your AI.
Star on GitHub Not ready to scan? Get notified of updates.
Join the Discord for support and updates → VibeDrift Catch architectural drift in AI-generated code, before it spreads. Runs on your machine, free and open source.
