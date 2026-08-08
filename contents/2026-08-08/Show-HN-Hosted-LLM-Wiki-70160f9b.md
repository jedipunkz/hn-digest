---
source: "https://getmana.md/"
hn_url: "https://news.ycombinator.com/item?id=49219987"
title: "Show HN: Hosted LLM Wiki"
article_title: "Think. — Mana"
author: "zintus"
captured_at: "2026-08-08T09:27:31Z"
capture_tool: "hn-digest"
hn_id: 49219987
score: 2
comments: 0
posted_at: "2026-08-08T08:48:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hosted LLM Wiki

- HN: [49219987](https://news.ycombinator.com/item?id=49219987)
- Source: [getmana.md](https://getmana.md/)
- Score: 2
- Comments: 0
- Posted: 2026-08-08T08:48:57Z

## Translation

タイトル: HN を表示: ホストされた LLM Wiki
記事タイトル：考える。 — マナ
説明: マナは、あなたが寝ている間にナレッジ ベースを統合します。マージ、リンク、更新が行われるため、Wiki は腐らずに複合化します。
HN テキスト: Karpathy 先生の要点が私に直撃し、アドホック バージョンを自分で数か月間使用した後、エンドツーエンドのホスト型バージョンを構築しました。重要なのは、エージェントに同じファイルを反復処理させ、過去の結果を再利用させることです。かなりうまくいきます。その下には、ホストされた Pi (:love:) エージェント ランタイム、いくつかのテンプレート、Astro レンダラー、いくつかのウィジェット、および GitHub リポジトリに支えられた Markdown Wiki のホスティングがあります。私の使用方法の良い例は、移動が多いため、旅行の計画です。組み込みの Web 検索のほか、チェックリストやライブ天気などのいくつかのウィジェットがあり、便利です。もう 1 つ非常に優れていたのは、Karpathy が提案した論文の取り込みと相互リンクのワークフローです。書かれた内容はすべて自分の github リポジトリにコミットされるため、いつでも好きなときに取得でき、編集履歴も取得できます。私自身は Obsidian ユーザーではありませんが、Getmana を既存の Obsidian 保管庫に向けることは、ほとんど問題なく機能するようです。インポートを試したい場合は、いくつかの鋭いエッジを修正するようにエージェントに依頼することも簡単です。ここに GitHub ユーザー名をドロップするか、Web サイトの「早期アクセスに参加」を使用してください。手動の手順は嫌いですが、今のところは自分でユーザーを追加する必要があります。トークンはまだ安くはなく、いかなる種類の請求もありません。

記事本文:
考えてみましょう。 — マナ マナ ドキュメント サインアップ それ自体を書く wiki。
マナは AI によって管理される Wiki です。論文、トランスクリプト、および記事を /raw にドロップします。エージェントはそれらを読み取り、引用し、相互参照し、あなたが所有するプレーンなマークダウン ファイル内の構造化 Wiki に織り込みます。
✓ 価格設定の主張に引用を追加
✓ 重複した「AI メモリ」ページを統合
• 3 つの矛盾を検討する準備ができています
Wiki はすでに整理されているため、すべての答えが簡単になります。
紙を /raw にドロップした後のエージェントの動作
Obsidian およびマークダウン ユーザー向け
Obsidian スタイルの AI Wiki は、あらゆるマークダウン ナレッジ ベースに威力を発揮します。
プレーンなマークダウンで、Mana を第 2 の脳のホスト型 AI アシスタントとして使用します。メモ、PDF、トランスクリプト、記事をフィードします。それらを引用され、リンクされた Wiki ページに変換し、すべての変更をレビュー可能な git diff として返すため、ナレッジ ベースは移植可能であり、あなたのものになります。 Obsidian や Notion との比較をご覧ください。
あらゆる重要なトピックについて、時間をかけて知識を蓄積している人向け。
数か月にわたる記事、ポッドキャスト、PDF — Wiki には、学んだすべての生きた地図が保存されています。
各章をファイルし、読みながら登場人物、テーマ、つながりのページを作成します。
旅行のチェックリスト、プロジェクトのプレイブック、日常のメモなど、実行するたびに次の実行が改善されます。
論文、トランスクリプト、記事、または PDF を /raw にドロップします。エージェントはそれを読み、永続的なアイデアを抽出し、関連ページを更新し、概念をリンクし、矛盾にフラグを立てます。 1 つのソースで 10 個の音を改善できます。
質問してください。マナは、回答を総合ページに変換します。引用され、適切な概念にリンクされ、知識ベースの拡大に合わせて再利用できます。
あなたの知識が生み出す軌跡をたどってください。マナはつながり、ギャップ、矛盾、そして次の疑問を明らかにし、有益な発見をリンクされたページ、インデックス、ログに変換します。
建てられた

ロックインではなく、所有権のために。
プレーン ファイル、Git 履歴、ローカル レビュー、オフライン アクセス - したがって、ナレッジ ベースは移植可能で検査可能であり、あなたのものになります。
生のソースと Wiki ページはプレーンなマークダウン ファイルです。データベースやロックインはなく、どのエディターでも開きます。
すべてのエージェント操作はコミットです。いつでもナレッジ ベース全体の差分、元に戻す、分岐、またはクローンを作成できます。
マナは、ソースを読み取り、ページを更新し、概念をリンクし、矛盾にフラグを立て、引用された要約を自動的に書き込みます。
共有シートからキャプチャしたり、音声メモを口述したり、機内でナレッジベースを読んだり、差分を確認したり、オンラインに戻ったら同期したりできます。
論文を収集し、フィールドマップを作成し、引用されたレビューを書きます。
定期的な旅行レッスンを再利用可能なチェックリストに変え、旅行のたびに改善します。
論文を、数学、引用、ソースの概要を備えたクリーンなサイトネイティブの Markdown に変換します。
別のツールから来ていますか?正直な比較 - 他のツールが正しい選択であるかどうかを含みます。
古典的なトレードオフ (所有するファイルとオールインワン ワークスペース)、そして 3 番目のオプション。 →
同じ単純な値下げですが、誰が組織するのでしょうか? →
オールインワンのワークスペースですか、それとも自分で管理する自分が所有するファイルですか? →
退職理由別に後任者を選ぶための正直なガイド。 →
山を知識ベースに変えます。
書類、保存した記事、会議メモ、通話記録、クリエイターのアーカイブ、ボイスメモ、または増え始めたマークダウン フォルダーなど、散らかったソースを持ち込んでください。
Mana はそれらを、それ自体を維持する Wiki に変えます。プレーンなマークダウン、git-backed、引用、リンクがあり、最初のコミットからレビュー可能です。
✓ あなたが完全に所有するマークダウン Wiki — ロックインや独自の形式はありません
✓ レビュー可能な git diff としてのすべての変更
✓ ソースを追加すると複合する合成ページ
✓ モバイルキャプチャ、デスクトップレビュー

## Original Extract

Mana integrates your knowledge base while you sleep — merging, linking, and refreshing so your wiki compounds instead of rotting.

Karpathy-sensei gist hit me directly, and i built end to end hosted version of it, after using ad-hoc version myself for a few months. Whole point is to have agent iterate on same files, and reuse past results. It works quite well. Underneath, it’s a hosted Pi (:love:) agent runtime, some templates, an Astro renderer, a few widgets, and hosting for a Markdown wiki backed by your GitHub repo. Good example of how i use it is travel planning, since i move quite a bit. It has builtin web search, and a few widgets like checklists and live weather, which makes it useful. Another thing which was quite good is Karpathy proposed workflow of ingesting and cross linking papers. Everything written is commited to your own github repo, so you can take it whenever you want, and you get edit history. I’m not an Obsidian user myself, but pointing Getmana at an existing Obsidian vault seems to mostly just work. It's also trivial to ask agent to fix a few sharp edges, if you want to try import. Drop your GitHub username here, or use “Join early access” on the website. I hate the manual step, but I have to add people myself for now. Tokens aren’t cheap yet, and there’s no billing of any kind.

Think. — Mana Mana Docs Sign up A wiki that writes itself.
Mana is an AI-maintained wiki. Drop papers, transcripts, and articles into /raw — the agent reads, cites, cross-references, and weaves them into a structured wiki in plain markdown files you own.
✓ citations added to pricing claims
✓ duplicate “AI memory” pages merged
• 3 contradictions ready for review
Every answer gets easier because the wiki is already organized.
What the agent does after you drop a paper into /raw
For Obsidian and markdown users
Obsidian-style AI wiki powers for any markdown knowledge base.
Use Mana as a hosted AI assistant for your second brain in plain markdown. Feed it notes, PDFs, transcripts, and articles; it turns them into cited, linked wiki pages and returns every change as a reviewable git diff, so the knowledge base stays portable and yours. See how it compares to Obsidian and Notion .
For anyone accumulating knowledge over time — on any topic that matters.
Articles, podcasts, PDFs over months — the wiki keeps a living map of everything you've learned.
File each chapter, build pages for characters, themes, and connections as you read.
Travel checklists, project playbooks, routine notes — each run makes the next one better.
Drop a paper, transcript, article, or PDF into /raw . The agent reads it, extracts the durable ideas, updates related pages, links concepts, and flags contradictions. One source can improve ten notes.
Ask questions. Mana turns the answers into synthesis pages — cited, linked to the right concepts, and reusable as your knowledge base grows.
Follow the trails your knowledge creates. Mana surfaces connections, gaps, contradictions, and next questions — then turns the useful discoveries into linked pages, indexes, and logs.
Built for ownership, not lock-in.
Plain files, git history, local review, and offline access — so your knowledge base stays portable, inspectable, and yours.
Raw sources and wiki pages are plain markdown files. No database, no lock-in, open in any editor.
Every agent operation is a commit. Diff, revert, branch, or clone your entire knowledge base at any moment.
Mana reads sources, updates pages, links concepts, flags contradictions, and writes cited summaries — automatically.
Capture from the share sheet, dictate voice notes, read your knowledge base on a plane, review diffs, sync when back online.
Ingest papers, build field maps, write cited reviews.
Turn recurring travel lessons into reusable checklists that improve after each trip.
Turn papers into clean, site-native Markdown with math, citations, and source overviews.
Coming from another tool? Honest comparisons — including when the other tool is the right choice.
The classic trade-off — files you own vs an all-in-one workspace — and a third option. →
Same plain markdown — but who does the organizing? →
All-in-one workspace, or files you own that maintain themselves? →
An honest guide to picking a replacement — by why you're leaving. →
Turn the pile into a knowledge base.
Bring your messy sources — papers, saved articles, meeting notes, call transcripts, creator archives, voice memos, or markdown folders that have started to sprawl.
Mana turns them into a wiki that maintains itself: plain markdown, git-backed, cited, linked, and reviewable from the first commit.
✓ A markdown wiki you fully own — no lock-in, no proprietary format
✓ Every change as a reviewable git diff
✓ Synthesis pages that compound as you add sources
✓ Mobile capture, desktop review
