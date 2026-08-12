---
source: "https://twitter.com/lifeofjer/status/2087608101293916466"
hn_url: "https://news.ycombinator.com/item?id=49277116"
title: "We started tracking which AI model wrote every line, and you should too"
article_title: "JER on X: \"https://t.co/KbwvXDAO6v\" / X"
author: "jeremyccrane"
captured_at: "2026-08-12T19:57:13Z"
capture_tool: "hn-digest"
hn_id: 49277116
score: 1
comments: 0
posted_at: "2026-08-12T19:05:17Z"
tags:
  - hacker-news
  - translated
---

# We started tracking which AI model wrote every line, and you should too

- HN: [49277116](https://news.ycombinator.com/item?id=49277116)
- Source: [twitter.com](https://twitter.com/lifeofjer/status/2087608101293916466)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T19:05:17Z

## Translation

タイトル: どの AI モデルが各行を記述したかの追跡を開始しました。あなたもそうすべきです
記事タイトル: X 上の JER: "https://t.co/KbwvXDAO6v" / X
説明: どの AI モデルが各行を記述したかの追跡を開始しました。あなたもそうすべきです

記事本文:
X の JER: "https://t.co/KbwvXDAO6v" / X ポスト
JER @lifeofjer 私たちはどの AI モデルが各行を書き込んだかを追跡し始めました。あなたもそうすべきです
私たちは AI エージェント @atlas_v_erified と 1 年以上同じアプリを構築してきました。
モデルは1つではありません。たくさん。テストモデル、QAモデル、構築モデル、設計モデル。フロントエンド、バックエンド、スキーマ、アーキテクチャ、そのすべて。 Cursor は新しいもの (彼らや他のラボ) を出荷し続けており、私たちはそれらを試し続けています。
そして、その年のほとんどの期間、私たちは非常に基本的な質問に答える本当の方法を持っていませんでした: どのモデルが実際にこのコードを書いたのか、それともどのモデルがこの計画を作成したのか?
たとえば作品4.7を見てみましょう。有望ではありましたが、頻繁に使用した後、作業が完了し、一部が完了しただけであると表示されることがわかりました。部分的な実装。私たちはそれをすぐに理解し、別のモデルに移行しました。しかし、それはまだ単なる感覚です。確かなデータではなく、逸話です。
そこで、モデル追跡をリポジトリに組み込みました。
エージェントがファイルに触れるたびに、署名欄を残す必要があります。機械可読。同じセッション、同じスプリント。
編集を行ったコーディングモデル
プランナー モデル (計画が別のモデルの場合)
それらの変更が生じた計画
それがバグ修正である場合は、どの変更が修正されるのか（実際にわかっている場合）
通常の TypeScript ファイルでは、上部は次のようになります。
実装する際に重要となるルールをいくつか紹介します。
ファイル内の最新の署名欄のみ。次の編集により置き換えられます。 Git には履歴が保存されます。
ハンクには来歴ブロック自体が含まれていないか、ヘッダーが移動した瞬間に行番号が役に立たなくなります。
Cursor が実際のモデル ID / 最大モード / コンテキスト ウィンドウを公開しない場合は、not_exused と書きます。エージェントがそれを発明することは許可されていません。
プランナーとコーダーは意図的に別の分野です。私たちの仕事の多くは、あるモデルによって計画され、別のモデルによって実装されます。
JSON、ロックファイル

s、バイナリはコメントを受け取ることができません。コメントは .cursor/agent-provenance.jsonl に 1 行として書き込まれます。このサイドカーは、「この .ts ファイルが大きかった」という抜け穴ではありません。
さて、ここからが興味深いところです。最初の試みは、AGENTS.md のスキーマと厳密なルールでした (はい、プロンプトがルールでもセキュリティでも要件でもないことはわかっていますが、これは迅速なテストです)。
都合が悪くなったとき、エージェントはそれを無視しました。左ハンク: 保留中。大規模なプランやスプリントの終わりにヘッダーをスキップしました。
そこで、彼らがスキップできないものにしました、そしてそれはうまくいきました。
常時オンカーソルルール - 毎ターンの短いリマインダー
カーソルフック — すべての編集を追跡し、書き込み後に小言を言い、何かが欠けている場合は停止時に強制的にクリーンアップターンを 1 回実行します
プリコミット — チェッカーがコミットをブロックする
CI — PR diff の同じチェッカー
実際に動作を変更したのは停止フックです。
エージェントはそれが完了したと考えます → そのセッションのすべてのファイルを再チェックします → 来歴が欠落している場合は、強制的に元に戻されます。
PROVENANCE GATE: 有効な atlas-agent-provenance/v1 属性なしでファイルを編集しました。新しい機能の作業を開始しないでください。以下のすべてのファイルの出自を修正してから停止します。
最大1ループ事前コミットと CI が残りをキャッチします。
ルール (または .md) はモデルのオプションです。フックと CI はそうではありません。
これがコードベース全体に渡されると、バグやアーキテクチャの変更が計画されていた箇所を実際にデバッグし、モデルごとに実装できるようになります。
どのコーディング モデルで最も多くの手戻りが発生するか
最初のパスでクリーンに出荷されるプランナー → コーダーのコンボはどれですか
バグを修正するとき、change_id とそれを導入したモデルまで遡ってバグを修正できますか?
特定のモデルに集中する「完了したと言って半分しか実行しなかった」故障モードです
このプロジェクトの開始時のデータがあればよかったのにと思います。それは私たちにとって信じられないほど貴重なものとなるでしょう。すべてのコードをレビューする計画を実行することを想像してください。

特定のモデルによって (ループ内で) 生成されるか、すべての問題をトレースバックして、どのモデルが最も問題を生成するかに関するデータを実行します (バックエンド、フロントエンド、DB スキーマなどによるセグメント化も)。
そこで、誰でも利用できるようにパブリック リポジトリをここにドロップし、誰でもワンクリックで追加できるようにこれを Cursor Marketplace に送信しました (承認されたら 🤞 @cursor_ai)。
Cursor (@cursor_ai) を使用して構築している内容を下にコメントしてください。何が構築されているか見てみたいと思います。
https://github.com/atlas-verified/agent-model-provenance
JER @lifeofjer 1h ここにカーソルを追加:
GitHub - atlas-verified/agent-model-provenance: カーソル プラグイン: 強制モデル署名 (コーダー + プランナー... github.com より 16
何が起こっているかを確認して会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名または電子メールでログイン 関係者

## Original Extract

We started tracking which AI model wrote every line, and you should too

JER on X: "https://t.co/KbwvXDAO6v" / X Post
JER @lifeofjer We started tracking which AI model wrote every line, and you should too
We've been building the same app for over a year with AI agents @atlas_v_erified.
Not one model. A bunch of them. Testing models, QA models, building models, design models. Frontend, backend, schema, architecture — all of it. Cursor keeps shipping new ones (theirs and other labs) and we keep trying them.
And for most of that year period we had no real way to answer a pretty basic question: which model actually wrote this code or which model created this plan?
Take for example Opus 4.7. Promising, but we learned after heavy use that it would say the work was done and only complete part of it. Partial implementations. We figured that out fast and moved to a different model. But that's still just a feeling. Anecdote, not hard data.
So I built model tracking into the repo.
Every time an agent touches a file, it has to leave a byline. Machine-readable. Same session, same sprint.
the coding model that made the edit
the planner model, if planning was a different model
the plan those changes came from
and if it's a bug fix, which earlier change it's fixing (when we actually know)
On a normal TypeScript file it looks like this at the top:
Here are a few rules that matter when implemented:
Newest byline only in the file. Next edit replaces it. Git keeps the history.
Hunks don't include the provenance block itself, or the line numbers get useless the second the header moves.
If Cursor doesn't expose the real model id / max mode / context window, we write not_exposed. Agents are not allowed to invent it.
Planner and coder are separate fields on purpose. A lot of our work is planned by one model and implemented by another.
JSON, lockfiles, binaries can't take comments — those go as one line into .cursor/agent-provenance.jsonl. That sidecar is not a loophole for "this .ts file was big."
Now, here is where it gets interesting. First attempt was the schema and a hard rule in our AGENTS.md (yes, I know that prompts are NOT rules or security or requirements, but it's a fast test).
The agents ignored it when it got inconvenient. Left hunks: pending. Skipped headers at the end of large plans and sprints.
So I made it into something they can't skip, and it's working great.
Always-on Cursor rule — short reminder every turn
Cursor hooks — track every edit, nag after writes, and on stop force one cleanup turn if anything is missing
Pre-commit — checker blocks the commit
CI — same checker on the PR diff
The stop hook is the one that actually changed behavior.
Agent thinks it's done → we re-check every file from that session → if provenance is missing, it gets forced back in:
PROVENANCE GATE: You edited files without valid atlas-agent-provenance/v1 attribution. Do not start new feature work. Fix provenance on every file below, then stop.
One loop max. Pre-commit and CI catch the rest.
Rules (or .md) are optional to a model. Hooks and CI are not.
Once this is all the way through our codebase, we can actually debug where bugs or architectural changes were planned and implement by model.
which coding models create the most rework
which planner → coder combos ship clean on the first pass
when we fix a bug, can we walk it back to the change_id and the model that introduced it
is the "said it was done, only did half" failure mode concentrated in specific models
I really wish I had this data back to the start of this project. It would be incredibly valuable to us. Imagine running a plan to review ALL code generated by a specific model (in loops), or tracing back every issue and running data on which models produce the least issues (even segmenting by back-end, front-end, db schema, etc).
So, I'm going to drop the public repo here so anyone can utilize this, and I have submitted this to Cursor Marketplace for anyone to one-click add (if they approve, 🤞 @cursor_ai).
Drop a comment below with what you are building using Cursor (@cursor_ai). Would love to see what is being built.
https://github.com/atlas-verified/agent-model-provenance
JER @lifeofjer 1h Add to Cursor here:
GitHub - atlas-verified/agent-model-provenance: Cursor plugin: force model bylines (coder + planner... From github.com 16
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
