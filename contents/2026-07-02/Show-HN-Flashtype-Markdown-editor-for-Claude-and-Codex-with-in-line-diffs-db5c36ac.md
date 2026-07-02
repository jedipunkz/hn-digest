---
source: "https://flashtype.com/"
hn_url: "https://news.ycombinator.com/item?id=48764289"
title: "Show HN: Flashtype – Markdown editor for Claude and Codex with in-line diffs"
article_title: "Flashtype | The markdown editor for Claude & Codex"
author: "samuelstros"
captured_at: "2026-07-02T17:52:56Z"
capture_tool: "hn-digest"
hn_id: 48764289
score: 5
comments: 0
posted_at: "2026-07-02T17:00:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Flashtype – Markdown editor for Claude and Codex with in-line diffs

- HN: [48764289](https://news.ycombinator.com/item?id=48764289)
- Source: [flashtype.com](https://flashtype.com/)
- Score: 5
- Comments: 0
- Posted: 2026-07-02T17:00:54Z

## Translation

タイトル: Show HN: Flashtype – インライン diff を備えた Claude および Codex 用のマークダウン エディター
記事タイトル: Flashtype | Claude & Codex のマークダウン エディター
説明: クロード コードとコーデックスが組み込まれた WYSIWYG マークダウン エディター。エージェントが編集し、差分をレビューします。すべての変更を受け入れるか拒否します。無料でオープンソース。
HN テキスト: Claude/Codex とコラボレーションするためのマークダウン エディターを改善したかったので、Flashtype を構築しました ( https://github.com/opral/flashtype ): - ローカルのマークダウン ファイルを開きます - Claude/Codex はネイティブに統合されました (既存のサブスクリプションと!) - 変更をすばやく承認/拒否するためのインライン diff - WYSIWYG (!) インライン diff は、他のアプリの静的プレビューに比べて大きな進歩であり、理論的には、最初の共同作業にすぎません。原始的な。しかし、「編集」は過去のものになるかもしれないというフィードバックを受けています。モデルが修正の必要がないほど良くなっているのに、なぜ Codex/Claude との差分を見る必要があるのでしょうか?そこで気になるのですが、AI を使用した執筆環境は過去数か月間でどのように変化しましたか?まだ投稿を繰り返しているのでしょうか、それともエージェントが一発で投稿しているのでしょうか?

記事本文:
フラッシュタイプ | Claude & Codex Flashtype のマークダウン エディター GitHub ↗ ダウンロード マークダウン エディター
クロードとコーデックスのための
ローカル マークダウン ファイル用の無料のオープンソース エディター - Claude Code と Codex が組み込まれています。ドキュメントのように記述し、エージェントの編集を差分としてレビューします。
★ Star on GitHub 無料 & オープン ソース · macOS Flashtype - san-francisco-blog-post.md ファイル AGENTS.md san-francisco-blog-post.md twitter-script.md writing-style.md サンフランシスコ: ベイ・バイ・ザ・ベイの街 サンフランシスコスタンドは、太平洋の海岸線と合流するなだらかな丘陵が特徴のアメリカで最も象徴的な都市、大都市圏の 1 つとして区別されます。有名なゴールデン ゲート ブリッジは、錆色の建築の驚異として湾に架かっています。
午後の霧がゴールデンゲートを通過し、スカイラインを水彩画のように柔らかくしている間、ケーブルカーは依然としてノブヒルに向かって登ります。
> イントロを引き締めて差分を確認する
⏺ san-francisco-blog-post.md を編集 - 8 件追加、12 件削除
⏺ 微気候都市のセクションを追加
Flashtype をディスク上の任意のフォルダー (メモ、ドキュメント、リポジトリ) に指定します。すべてのドキュメントはプレーンな .md ファイルです。同期サービス、独自のフォーマット、ロックインはありません。
ブログ ドラフト san-francisco.md 公開 AGENTS.md writing-style.md サンフランシスコ: 湾沿いの街 アメリカで最も象徴的な都市の 1 つで、太平洋の海岸線に接するなだらかな丘陵が特徴です。
見出し、リスト、リンクは入力と同時にレンダリングされます。分割プレビューやマークダウン構文は表示されません。書いている間、完成したページのように読み取れます。
san-francisco.md H1 H2 B I U リンク • リスト サンフランシスコ: 湾沿いの街 アメリカで最も象徴的な都市の 1 つで、太平洋の海岸線に接するなだらかな丘陵が特徴です。ゴールデン ゲート ブリッジは、錆色の驚異として湾に架かっています。
横のペインで Claude Code または Codex を実行します。

あなたのドラフト。彼らは同じファイルを読み取って編集します。コピーアンドペーストやコンテキストのジャグリングはありません。
san-francisco.md Claude Code クロード コード ターミナル エージェント セッション · ローカル プロジェクトを編集する準備ができました
~/Documents/flashtype > イントロを締めてセクションを追加
なだらかな丘陵が太平洋の海岸線に接し、ゴールデン ゲート ブリッジが湾にかかっています。
エージェントは、ドキュメント内のインライン差分としてランドを変更します。良いものは受け入れ、残りは拒否します。あなたが見なければ何も変わりません。
san-francisco.md サンフランシスコは、太平洋の海岸線につながるなだらかな丘陵が特徴的な、アメリカで最も象徴的な都市、大都市圏の 1 つとして知られています。
すべての変更は自動的にチェックポイントに記録されます。ドキュメントの完全な履歴を参照し、ワンクリックで以前のバージョン (自分またはエージェントのバージョン) を復元します。
san-francisco.md バージョン履歴 現在のドラフト · You Claude · 強化されたイントロ 2 分前 · +8 -12 このバージョンを復元 微気候を追加 14 分前 · +21 昨日作成 湾沿いの街 なだらかな丘陵が太平洋の海岸線に接し、ゴールデン ゲート ブリッジが湾にかかっています。
Flashtype のバージョン履歴と差分は Lix 上で実行されます。すべての編集にはチェックポイントが設定されるため、変更内容を確認して、任意のバージョンに戻ることができます。
Flashtype は、Claude Code と Codex が組み込まれた、無料のオープンソース macOS マークダウン エディタです。
Flashtypeはローカルファイルでも動作しますか?
はい。ディスク上の任意のフォルダーを開いて、独自の形式を使用せずにプレーンな .md ファイルに書き込み続けます。
Claude Code と Codex は同じファイルを編集でき、Flashtype は変更を承認または拒否する前にインライン diff として表示します。
野外に建てられています。問題、プルリクエスト、スターは大歓迎です。

## Original Extract

A WYSIWYG markdown editor with Claude Code and Codex built in. Agents edit, you review diffs. Accept or reject every change. Free and open source.

I wanted to better markdown editor for collaborating with Claude/Codex and built Flashtype ( https://github.com/opral/flashtype ): - opens local markdown files - Claude/Codex natively integrated (with my existing subscription!) - in-line diffing to quickly accept/reject changes - WYSIWYG (!) The in-line diffs are a big step up to other apps static previews, and, in theory, just the first collaborative primitive. But, I am getting feedback that "editing" might become a thing of the past. Why do I need to see diffs from Codex/Claude if the models are getting so good to the point that they need no correction? Which makes me wonder, how did your writing setup with AI change over the past months? Are you still iterating on posts or are agents one shotting them?

Flashtype | The markdown editor for Claude & Codex Flashtype GitHub ↗ Download The markdown editor
for Claude & Codex
A free, open-source editor for your local markdown files - with Claude Code and Codex built in. Write like a doc, review agent edits as diffs.
★ Star on GitHub Free & open source · macOS Flashtype - san-francisco-blog-post.md Files AGENTS.md san-francisco-blog-post.md twitter-script.md writing-style.md San Francisco: City by the Bay San Francisco stands is distinguished as one of America's most iconic cities, metropolitan areas, characterized by rolling hills that converge with the Pacific coastline. The celebrated Golden Gate Bridge spans the bay as a rust-colored architectural marvel.
Cable cars still climb toward Nob Hill while afternoon fog rolls through the Golden Gate, softening the skyline into watercolor.
> tighten the intro and review diffs
⏺ Edited san-francisco-blog-post.md - 8 additions, 12 deletions
⏺ Added section A city of microclimates
Point Flashtype at any folder on your disk - your notes, docs, or a repo. Every document is a plain .md file. No sync service, no proprietary format, no lock-in.
blog drafts san-francisco.md published AGENTS.md writing-style.md San Francisco: City by the Bay One of America's most iconic cities, characterized by rolling hills that meet the Pacific coastline.
Headings, lists and links render live as you type - no split preview, no markdown syntax in your face. It reads like the finished page while you write.
san-francisco.md H1 H2 B I U Link • List San Francisco: City by the Bay One of America's most iconic cities, characterized by rolling hills that meet the Pacific coastline. The Golden Gate Bridge spans the bay as a rust-colored marvel.
Run Claude Code or Codex in a pane next to your draft. They read and edit the same files - no copy-paste, no context juggling.
san-francisco.md Claude Code Claude Code terminal Ready to edit Agent session · local project
~/Documents/flashtype > tighten the intro and add a section
Rolling hills meet the Pacific coastline, and the Golden Gate Bridge spans the bay.
Agent changes land as inline diffs right in your document. Accept the good ones, reject the rest - nothing changes without you seeing it.
san-francisco.md San Francisco stands is distinguished as one of America's most iconic cities, metropolitan areas, characterized by rolling hills that converge with the Pacific coastline.
Every change is checkpointed automatically. Browse a document's full history and restore any earlier version in one click - yours or the agent's.
san-francisco.md Version History Current draft now · You Claude · tightened intro 2m ago · +8 -12 Restore this version Added microclimates 14m ago · +21 Created yesterday City by the Bay Rolling hills meet the Pacific coastline, and the Golden Gate Bridge spans the bay.
Flashtype's version history and diffs run on Lix. Every edit is checkpointed, so you can see what changed and go back to any version.
Flashtype is a free, open-source macOS markdown editor with Claude Code and Codex built in.
Does Flashtype work with local files?
Yes. Open any folder on disk and keep writing in plain .md files without a proprietary format.
Claude Code and Codex can edit the same files, and Flashtype shows their changes as inline diffs before you accept or reject them.
Built in the open. Issues, pull requests and stars welcome.
