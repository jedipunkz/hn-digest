---
source: "https://theaileverageweekly.com/posts/how-to-write-better-git-commit-messages-with-ai.html"
hn_url: "https://news.ycombinator.com/item?id=48538531"
title: "How to Write Better Git Commit Messages with AI"
article_title: "How to Write Better Git Commit Messages with AI — The AI Leverage Weekly"
author: "talvardi7"
captured_at: "2026-06-15T11:11:39Z"
capture_tool: "hn-digest"
hn_id: 48538531
score: 1
comments: 1
posted_at: "2026-06-15T09:01:07Z"
tags:
  - hacker-news
  - translated
---

# How to Write Better Git Commit Messages with AI

- HN: [48538531](https://news.ycombinator.com/item?id=48538531)
- Source: [theaileverageweekly.com](https://theaileverageweekly.com/posts/how-to-write-better-git-commit-messages-with-ai.html)
- Score: 1
- Comments: 1
- Posted: 2026-06-15T09:01:07Z

## Translation

タイトル: AI を使用してより適切な Git コミット メッセージを作成する方法
記事のタイトル: AI を使用してより適切な Git コミット メッセージを作成する方法 — The AI Leverage Weekly
説明: 不正なコミット メッセージは永久に支払う税金です。 6 か月後、git log にアクセスすると、「修正内容」、「WIP」、「asdf」、「最終版 v2」の壁を発見します。そして…

記事本文:
AI Leverage Weekly エンジニアのための実践的な AI ワークフロー
無料プロンプト 購読する
← すべての投稿
AI を使用してより適切な Git コミット メッセージを作成する方法
間違ったコミットメッセージは永久に支払う税金です。 6 か月後、git log にアクセスすると、「修正内容」、「WIP」、「asdf」、「最終版 v2」の壁に遭遇しました。そして今、あなたはエンジニアではなく考古学者になっています。良いニュース: AI はこの問題をほぼ完全に解決できます。このチュートリアルでは、今日から使用できるコピー & ペースト プロンプトを使用してこの問題をワークフローに組み込む方法を正確に説明します。
メッセージのコミットが失敗する理由 (そして AI が実際にここで役立つ理由)
根本的な原因は怠惰ではなく、タイミングにあります。コミット メッセージは、長いコーディング セッションの直後、頭が疲れていてとにかくプッシュしたいときに書きます。コンテキストはすべて頭の中にあり、画面上にはありません。
AIはこれをひっくり返します。あなたはそれに差分とあなたの大まかなメモを送ります。構造化されたメッセージを作成します。正確さを期すために編集します。合計時間: 30 秒。成果物は、ほとんどのエンジニアがプレッシャーを受けて書いたものよりも常に優れています。
ステップ 1: 変更をステージングして差分を生成する
プロンプトを表示する前に、コミットしようとしている内容のクリーンな diff を取得してください。
git diff --staged
全画面モードに入る
全画面モードを終了する
出力をコピーします。差分が大きい場合 (500 行以上)、最も意味のあるファイルに絞り込みます。
git diff --staged -- path/to/relevant/file.ts
全画面モードに入る
全画面モードを終了する
プロンプトにすべての行を貼り付ける必要はありません。代表的なスライスと意図の平易な英語の要約で十分です。
これを選択した AI アシスタント (ChatGPT、Claude、Copilot Chat など) に貼り付けます。
あなたは、Git コミット メッセージの作成を手伝ってくれるシニア エンジニアです。
段階的な差分は次のとおりです。
<差分をここに貼り付け>
私の意図: <この変更によって何が達成されるか、そしてその理由についての一文>
コミットを書いてください

従来のコミット形式を使用した ssage:
- 最初の行: type(scope): 短い命令の概要 (最大 72 文字)
- 空白行
- 本文: 「どのように」ではなく、「何が変わったのか」と「なぜ」を説明する 2 ～ 4 つの箇条書き
- 重大な変更がある場合は、重大な変更フッターを追加します
コミットメッセージのみを出力し、コメントは出力しません。
全画面モードに入る
全画面モードを終了する
トークン更新のバグ修正の実際の出力例は次のようになります。
fix(auth): 期限切れのセッションでのサイレント トークンの更新を防止します
- セッション TTL がすでに経過した場合の自動リフレッシュ呼び出しを削除
- refreshToken() を呼び出す前に明示的な有効期限チェックを追加します。
- 重複したリフレッシュ要求を引き起こす競合状態を防止します。
ネットワークが遅い状況下では
全画面モードに入る
全画面モードを終了する
これは、次のエンジニアが git log を読むときにすぐに役に立ちます。
ステップ 3: 一行のフォローアップで絞り込む
最初の草案は近いものの完全に正しくない場合は、最初から説明し直す必要はありません。特定の問題を修正するだけです。
スコープは「auth」ではなく「session」である必要があり、最初の行は
は長すぎます。60 文字未満にまとめてください。
全画面モードに入る
全画面モードを終了する
AI はこのような外科的編集をうまく処理します。あなたは査読者です。初稿の作家です。
ステップ 4: 速度向上のためにシェル エイリアスを構築する
手動でコピー＆ペーストする手間により、この習慣は消えてしまいます。それを自動化しましょう。これを .zshrc または .bashrc に追加します。
エイリアス gcm = 'git diff --staged | pbcopy && echo "差分がコピーされました。AI プロンプトに貼り付けてください。"'
全画面モードに入る
全画面モードを終了する
Linux では、 pbcopy を xclip -selection Clipboard に置き換えます。これで、ステージングされた差分が 1 つのコマンドでクリップボードに保存され、AI チャットに貼り付ける準備が整いました。
GitHub CLI を使用しているチームの場合は、さらに進んで、API を呼び出すスクリプトに直接パイプすることもできますが、手動のコピー＆ペーストの習慣だけで 80% の効果を得ることができます。

f 値。
ステップ 5: リンターを追加してフォーマットを強制する
良いメッセージを書くことは戦いの半分にすぎません。残りの半分はメッセージが退行しないようにすることです。 commitlint をリポジトリに追加します。
npm install --save-dev @commitlint/cli @commitlint/config-conventional
echo "module.exports = { extends: ['@commitlint/config-conventional'] };" > commitlint.config.js
npx ハスキー add .husky/commit-msg 'npx --no -- commitlint --edit "$1"'
全画面モードに入る
全画面モードを終了する
従来のコミット形式に従っていないコミットは、ブランチにアクセスする前に拒否されるようになりました。これを上記の AI プロンプトと組み合わせると、フォーマットの問題はほぼ完全に解消されます。
このプロンプト パターンは、私が『AI Leverage Playbook: 50 Prompts & Workflows for Engineers』にパッケージ化したものの 1 つですが、上記のバージョンだけでも実際の価値を得るのに十分です。
この習慣が確立されると、Git ログが実際のドキュメントになります。次のことができます。
git log --oneline を実行して、リリースの読み取り可能な変更ログを取得します
git log --grep="fix(auth)" を使用して、ソースを grep せずにすべての認証関連の修正を検索します
新しいエンジニアに Confluence ではなくコミット履歴を指示してオンボーディングします。
コミット メッセージは後付けではなく、第一級の成果物になります。
クイックリファレンス: プロンプト、もう一度
あなたは、Git コミット メッセージの作成を手伝ってくれるシニア エンジニアです。
差分:
<貼り付け>
意図: <一文>
形式: 従来のコミット。最初の行 ≤72 文字、命令形。
本文: 変更内容とその理由に関する 2 ～ 4 つの箇条書き。必要に応じて、重大な変更フッター。
コミットメッセージのみを出力します。
全画面モードに入る
全画面モードを終了する
それを保存してください。 「修正」と入力しようとするたびに、それに手を伸ばしてください。
The AI Leverage Weekly では、このようなワークフローを毎週 1 つずつ詳しく解説しています。実践的で、飾り気のない、無料のワークフローです。購読: https://theaileverageweekly。

beehiiv.com/subscribe?utm_source=devto&utm_medium=article&utm_campaign=long_w6
次のメールを受信箱で受け取る
エンジニア向けの実践的な AI ワークフロー。週に 1 冊、綿毛はありません。
AI を使用してオンボーディング時間を 40% 削減しました — これがまさに私たちがやったことです (第 5 週のまとめ)
AI を使用する若手開発者は不正行為をしていません - 彼らはより賢くトレーニングしています
エンジニアとしてブロックされないようにするために私が毎週達成する 10 のプロンプト

## Original Extract

Bad commit messages are a tax you pay forever. You hit git log six months later and find a wall of "fix stuff", "WIP", "asdf", and "final final v2" — and…

The AI Leverage Weekly Practical AI workflows for engineers
Free prompts Subscribe
← All posts
How to Write Better Git Commit Messages with AI
Bad commit messages are a tax you pay forever. You hit git log six months later and find a wall of "fix stuff", "WIP", "asdf", and "final final v2" — and now you're archaeologist instead of engineer. The good news: AI can eliminate this problem almost entirely, and in this walkthrough I'll show you exactly how to wire it into your workflow with copy-paste prompts you can use today.
Why Commit Messages Fail (And Why AI Actually Helps Here)
The root cause isn't laziness — it's timing. You write the commit message right after a long coding session, when your brain is fried and you just want to push. Context is all in your head, not on the screen.
AI flips this. You feed it the diff and your rough notes; it drafts a structured message. You edit for accuracy. Total time: 30 seconds. The output is consistently better than what most engineers write under pressure.
Step 1: Stage Your Changes and Generate a Diff
Before prompting, get a clean diff of what you're about to commit:
git diff --staged
Enter fullscreen mode
Exit fullscreen mode
Copy the output. If the diff is large (500+ lines), narrow it to the most meaningful files:
git diff --staged -- path/to/relevant/file.ts
Enter fullscreen mode
Exit fullscreen mode
You don't need to paste every line into the prompt — a representative slice plus a plain-English summary of your intent is enough.
Paste this into your AI assistant of choice (ChatGPT, Claude, Copilot Chat, etc.):
You are a senior engineer helping me write a Git commit message.
Here is the staged diff:
<paste diff here>
My intent: <one sentence about what this change accomplishes and why>
Write a commit message using the Conventional Commits format:
- First line: type(scope): short imperative summary (max 72 chars)
- Blank line
- Body: 2–4 bullet points explaining WHAT changed and WHY, not HOW
- If there's a breaking change, add a BREAKING CHANGE footer
Output only the commit message, no commentary.
Enter fullscreen mode
Exit fullscreen mode
A real example output for a token-refresh bug fix might look like this:
fix(auth): prevent silent token refresh on expired session
- Remove automatic refresh call when session TTL has already elapsed
- Add explicit expiry check before calling refreshToken()
- Prevents a race condition that caused duplicate refresh requests
under slow network conditions
Enter fullscreen mode
Exit fullscreen mode
That's immediately useful to the next engineer reading git log .
Step 3: Refine With a One-Line Follow-Up
If the first draft is close but not quite right, don't re-explain from scratch. Just correct the specific problem:
The scope should be "session" not "auth", and the first line
is too long — tighten it to under 60 characters.
Enter fullscreen mode
Exit fullscreen mode
AI handles surgical edits like this well. You're the reviewer; it's the first-draft writer.
Step 4: Build a Shell Alias for Speed
The friction of copy-pasting manually will kill this habit. Automate it. Add this to your .zshrc or .bashrc :
alias gcm = 'git diff --staged | pbcopy && echo "Diff copied. Paste into your AI prompt."'
Enter fullscreen mode
Exit fullscreen mode
On Linux, swap pbcopy for xclip -selection clipboard . Now your staged diff is on your clipboard in one command, ready to paste into any AI chat.
For teams using the GitHub CLI, you can go further and pipe directly into a script that calls an API — but the manual copy-paste habit alone will get you 80% of the value.
Step 5: Add a Linter to Enforce the Format
Writing good messages is only half the battle — the other half is making sure they don't regress. Add commitlint to your repo:
npm install --save-dev @commitlint/cli @commitlint/config-conventional
echo "module.exports = { extends: ['@commitlint/config-conventional'] };" > commitlint.config.js
npx husky add .husky/commit-msg 'npx --no -- commitlint --edit "$1"'
Enter fullscreen mode
Exit fullscreen mode
Now any commit that doesn't follow Conventional Commits format gets rejected before it ever touches your branch. Pair this with the AI prompt above and the format issues go away almost entirely.
This prompt pattern is one of the ones I've packaged into The AI Leverage Playbook: 50 Prompts & Workflows for Engineers — but the version above is enough to get real value on its own.
Once this habit is set, your git log becomes actual documentation. You can:
Run git log --oneline to get a readable changelog for a release
Use git log --grep="fix(auth)" to find every auth-related fix without grepping source
Onboard new engineers by pointing them at commit history, not Confluence
The commit message becomes a first-class artifact, not an afterthought.
Quick Reference: The Prompt, One More Time
You are a senior engineer helping me write a Git commit message.
Diff:
<paste>
Intent: <one sentence>
Format: Conventional Commits. First line ≤72 chars, imperative mood.
Body: 2–4 bullets on what changed and why. BREAKING CHANGE footer if needed.
Output the commit message only.
Enter fullscreen mode
Exit fullscreen mode
Save that. Reach for it every time you're about to type "fix stuff."
I break down one workflow like this every week in The AI Leverage Weekly — practical, no fluff, free. Subscribe: https://theaileverageweekly.beehiiv.com/subscribe?utm_source=devto&utm_medium=article&utm_campaign=long_w6
Get the next one in your inbox
Practical AI workflows for engineers. One issue a week, no fluff.
We Cut Onboarding Time by 40% Using AI — Here's Exactly What We Did (Week 5 Roundup)
Junior Devs Who Use AI Are Not Cheating — They're Training Smarter
10 Prompts I Reach for Every Week to Stay Unblocked as an Engineer
