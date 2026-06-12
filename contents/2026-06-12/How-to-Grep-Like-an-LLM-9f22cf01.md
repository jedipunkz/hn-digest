---
source: "https://marcusmichaels.com/notes/grep-like-an-llm/"
hn_url: "https://news.ycombinator.com/item?id=48502159"
title: "How to Grep Like an LLM"
article_title: "Grep like an LLM · notes to self"
author: "marcusmichaels"
captured_at: "2026-06-12T10:34:35Z"
capture_tool: "hn-digest"
hn_id: 48502159
score: 1
comments: 2
posted_at: "2026-06-12T10:16:10Z"
tags:
  - hacker-news
  - translated
---

# How to Grep Like an LLM

- HN: [48502159](https://news.ycombinator.com/item?id=48502159)
- Source: [marcusmichaels.com](https://marcusmichaels.com/notes/grep-like-an-llm/)
- Score: 1
- Comments: 2
- Posted: 2026-06-12T10:16:10Z

## Translation

タイトル: LLM のように Grep を実行する方法
記事タイトル: LLM のような Grep · 自分用メモ
説明: grep、git、および AI エージェントから借用したメソッドを使用してコードベースを回避する方法を見つける方法。

記事本文:
AI エージェントは、見たことのないコードベースに到達し、適切なコードベースを見つけることができます。
数秒でファイルに保存されます。マップもプロジェクトの記憶もありません。こう書かれています
ほとんど何もありません。
私にその方法を尋ねたところ、答えは grep 、 git 、そして質問が回答された瞬間に停止する方法でした。
grep は、すべての Mac および Linux に同梱されている検索プログラムです。
何十年も使ってきたマシン。単語を入力すると、すべての行が出力されます。
その単語を含むすべてのファイル。それがツール全体です。 git grep
git に組み込まれているのと同じアイデアで、プロジェクトのファイルのみを検索します
トラック。
これらのメモは、アイスクリーム バン ビジネスのコードベースである Sundae Service に基づいています。
動詞ではなく名詞を検索してください。何かを読む前に地図を作成してください。探す
エントリポイントを上から下に読んでください。ロジックの前に型を読んでください。トレース
書き込み場所から読み取り場所までの 1 つのフィールド。コードがうまくいかないとき
それ自体を説明し、歴史を尋ねてください。質問が答えられたら停止してください。
その各行は、シリーズの後半で独自のモジュールを取得しますが、
コマンドは初日から役立つので、ここに示します。ほとんどは git grep を使用します。
プロジェクトが追跡するファイルのみを検索する通常の grep です。
まさに、node_modules と結果からのビルド出力を維持するものです。
見慣れないものがあれば、それがモジュールの目的です。
7 つのフラグがほとんどの作業を行い、 grep では同じことを意味します。
git grep および ripgrep : -i 大文字と小文字を無視する、-l ファイル名のみ、-c
ファイルごとのカウント、-w 単語全体、-n 行番号、-F リテラル文字列
(正規表現オフ)、-C 3 各一致に関するコンテキストを 3 行表示します。あと2つ
git log から: --oneline は 1 行あたり 1 つのコミット、-S 'x' は
x を追加または削除するコミット。
これらを一度学習してください。これらは、これまでにインストールしたすべてのエディターよりも長持ちします。
各モジュールは通勤に適したサイズになっているため、電車の中でも読むことができます。
で

朝、その日のうちにデスクで試してみてください。全シリーズは、
以下。
プレイブック マップが読書に勝る理由と、モジュールを開始する前にその方法を完全に説明します。
ヒート マップとファイル リストを読むのではなく、マップを作成します。単一のファイルを開く前に機能を見つけます。
ドアの入口ポイントを見つけます。画面上のテキスト、ファイル名、一致する部分の周囲の行から始めます。
データを追跡する 最初に型を読み取り、部分文字列ではなく単語を照合し、値を書き込んだ人を見つけます。
歴史に尋ねる コード自体が説明できないときは、通常、コミットによって説明されます。
チートシート シリーズのすべてのコマンドが 1 か所にまとめられています。
© 2026 Marcus Michaels · ♥︎ とで構築
控えめなUI

## Original Extract

How to find your way around any codebase with grep, git, and a method borrowed from an AI agent.

An AI agent can land in a codebase it has never seen and find the right
file in seconds. It has no map and no memory of the project. It reads
almost nothing.
I asked mine how, and the answer was grep , git , and a method of stopping the moment your question is answered.
grep is a search program that has shipped with every Mac and Linux
machine for decades. You give it a word, and it prints every line of
every file that contains that word. That is the whole tool. git grep
is the same idea built into git, searching only the files your project
tracks.
These notes are based on Sundae Service, our codebase of a made-up ice cream van business.
Search the noun, not the verb. Build a map before you read anything. Find
the entry point and read top-down. Read the types before the logic. Trace
one field from where it's written to where it's read. When the code won't
explain itself, ask the history. Stop when your question is answered.
Each line of that gets its own module later in the series, but the
commands are useful from day one, so here they are. Most use git grep ,
ordinary grep that only searches the files your project tracks, which is
exactly what keeps node_modules and build output out of your results.
If any look unfamiliar, that's what the modules are for.
Seven flags do most of the work, and they mean the same thing in grep ,
git grep , and ripgrep : -i ignore case, -l filenames only, -c
count per file, -w whole word, -n line numbers, -F literal string
(regex off), -C 3 three lines of context around each match. Two more
from git log : --oneline for one commit per line, and -S 'x' for
commits that add or remove x .
Learn these once. They'll outlive every editor you ever install.
Each module is sized for a commute, meaning you can read one on the train
in the morning and try it at your desk the same day. The full series is
below.
Playbook Why maps beat reading, and the method in full before the modules begin.
Make a map, don't read The heat map and the file list. Locating a feature before opening a single file.
Find the door Entry points. Start from the text on screen, the names of files, and the lines around a match.
Follow the data Read the types first, match the word rather than the substring, and find who writes a value.
Ask the history When the code won't explain itself, the commits usually will.
The cheatsheet Every command from the series, in one place.
© 2026 Marcus Michaels · Built with ♥︎ and
modest-ui
