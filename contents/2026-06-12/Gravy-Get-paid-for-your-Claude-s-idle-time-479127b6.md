---
source: "https://gravycli.xyz"
hn_url: "https://news.ycombinator.com/item?id=48499212"
title: "Gravy: Get paid for your Claude's idle time"
article_title: "gravy — earn from your idle terminal time"
author: "dvptp"
captured_at: "2026-06-12T04:35:03Z"
capture_tool: "hn-digest"
hn_id: 48499212
score: 2
comments: 0
posted_at: "2026-06-12T02:38:54Z"
tags:
  - hacker-news
  - translated
---

# Gravy: Get paid for your Claude's idle time

- HN: [48499212](https://news.ycombinator.com/item?id=48499212)
- Source: [gravycli.xyz](https://gravycli.xyz)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T02:38:54Z

## Translation

タイトル: Gravy: クロードのアイドル時間に対して報酬を受け取ります
記事のタイトル: グレイビー — ターミナルのアイドル時間を利用して稼ぐ
説明: Claude Code の CLI ファースト広告マーケットプレイス。ステータス行にスポンサー付きの行を表示し、70% を維持します。端末から直接アドバタイズします。

記事本文:
グレービーソース
ベータ・テストモード
端末のアイドル時間を利用して収入を獲得しましょう。
Claude Code の CLI ファースト広告マーケットプレイス。目立たないスポンサー付きのレンダリング
ステータス行に行を追加して 70% を維持するか、次のように宣伝します。
開発者はターミナルから直接アクセスできます。ウェブアプリはありません。殻だけ。
ベータ版: 支払いは Stripe で実行されます
テストモード — 広告主はテストカードで支払います
( 4242… ) と収益は、関心を測定しながらシミュレーションされます。
# npm — 任意の OS、ノード 18 以降
npm install -g gravy-cli
# apt — Debian/Ubuntu、自己完結型 (ノードなし)
カール -fsSL https://gravycli.xyz/apt/gravy.gpg | sudo tee /usr/share/keyrings/gravy.gpg >/dev/null
echo "deb [signed-by=/usr/share/keyrings/gravy.gpg] https://gravycli.xyz/apt 安定したメイン" \
| sudo tee /etc/apt/sources.list.d/gravy.list
sudo apt update && sudo apt install グレービー
$ gravy # インタラクティブ メニュー — すべてのコマンドに対して gravy -h を実行します
広告を見ることで報酬を得る
グレイビーはクロードにワイヤーのスポンサー付き回線をインストールします
コードのステータス行。収益の 70% を保持します。 Stripe 経由で現金を引き出します。
グレービー広告 — ウォレットに資金を提供し、入札価格で運営します
キャンペーンを実施し、開発者が働いている場所にリーチします。グレービーマーケット
進行率を示します。
すべてのコマンドは生きています
グレービー -h で。これは CLI です。ターミナルはドキュメントです。

## Original Extract

A CLI-first ad marketplace for Claude Code. Render sponsored lines in your status line and keep 70%; advertise straight from the terminal.

gravy
BETA · TEST MODE
Earn from your idle terminal time.
A CLI-first ad marketplace for Claude Code. Render unobtrusive sponsored
lines in your status line and keep 70% — or advertise to
developers straight from the terminal. No web app. Just the shell.
Beta: payments run in Stripe
test mode — advertisers pay with test cards
( 4242… ) and earnings are simulated while we gauge interest.
# npm — any OS, Node 18+
npm install -g gravy-cli
# apt — Debian/Ubuntu, self-contained (no Node)
curl -fsSL https://gravycli.xyz/apt/gravy.gpg | sudo tee /usr/share/keyrings/gravy.gpg >/dev/null
echo "deb [signed-by=/usr/share/keyrings/gravy.gpg] https://gravycli.xyz/apt stable main" \
| sudo tee /etc/apt/sources.list.d/gravy.list
sudo apt update && sudo apt install gravy
$ gravy # interactive menu — run gravy -h for all commands
Get paid to see ads
gravy install wires sponsored lines into your Claude
Code status line. You keep 70% of the revenue; cash out via Stripe.
gravy advertise — fund a wallet, run bid-priced
campaigns, and reach developers where they work. gravy market
shows the going rate.
Every command lives
in gravy -h . It's a CLI — the terminal is the docs.
