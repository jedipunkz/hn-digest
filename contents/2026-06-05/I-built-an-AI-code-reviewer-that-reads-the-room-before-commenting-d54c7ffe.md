---
source: "https://codemouse.ai"
hn_url: "https://news.ycombinator.com/item?id=48410581"
title: "I built an AI code reviewer that reads the room before commenting"
article_title: "CodeMouse — AI Code Reviews on Every Pull Request"
author: "pro_methe5"
captured_at: "2026-06-05T13:10:26Z"
capture_tool: "hn-digest"
hn_id: 48410581
score: 2
comments: 1
posted_at: "2026-06-05T10:42:39Z"
tags:
  - hacker-news
  - translated
---

# I built an AI code reviewer that reads the room before commenting

- HN: [48410581](https://news.ycombinator.com/item?id=48410581)
- Source: [codemouse.ai](https://codemouse.ai)
- Score: 2
- Comments: 1
- Posted: 2026-06-05T10:42:39Z

## Translation

タイトル: コメントする前に部屋を読む AI コードレビューアーを構築しました
記事のタイトル: CodeMouse — すべてのプル リクエストの AI コード レビュー
説明: すべての GitHub PR の自動コード レビュー — Claude + GPT によるコンセンサス AI レビュー。人間または他のボットがすでに言ったことをスキップします。月額 10 ドル、14 日間のトライアル、自分のキーを持参してください。

記事本文:
CodeMouse — すべてのプル リクエストの AI コード レビュー
すべてのプルリクエストに対する自動化された AI コードレビュー
GitHub アプリをインストールします。 Anthropic キー、OpenAI キー、またはその両方をご持参ください。 Claude や GPT から、コンテキストを意識した詳細なコード レビューを数分で自動的に取得できます。
14 日間の無料トライアル。 $10/月。クレジットカードは必要ありません。
GitHub アプリをインストールする — CodeMouse を GitHub 組織または個人アカウントに追加します。接続するリポジトリを選択します。
AI キーを追加する — Anthropic キー、OpenAI キー、またはその両方を接続します。両方を接続すると、各 PR はクロードと GPT によって並行してレビューされます。キーは保存時に暗号化されます (AES-256-GCM)。
プル リクエストを開く — CodeMouse が自動的にレビューし、詳細なコメントを GitHub に直接投稿します。
完全なリポジトリ コンテキスト — レビューでは、diff だけでなくリポジトリ全体が使用されます。
クロード、GPT、または両方 - いずれかまたは両方を接続します。両方が接続されている場合は、結果がマージされ、重複が削除されるため、単一の統合レビューが得られます。
独自のキーの持ち込み — AI コストの値上げはありません。 Anthropic および/または OpenAI は直接支払います (プロバイダーごとにレビューごとに ~0.05 ～ 0.15 ドル)。
部屋を読む — 既存の PR コメント (人間、他のボット、以前の CodeMouse の実行) をすべてレビューし、すでに述べたことを繰り返しません。修正が完了すると、独自のスレッドをフォローアップします。
クリーンな PR を承認する — 差分が良好で、以前の問題が解決され、CI が緑色の場合、CodeMouse は実際に PR を承認します。
スロットリングなし — すべてのレビューは即座に開始されます。キューやレート制限はありません。
TypeScript、Python、Go、Rust、Java など、あらゆる言語で動作します。コードであれば、CodeMouse がレビューします。
インライン コメント — レビュー コメントは PR の特定の行に表示されます。
構成ゼロ — CI パイプライン、YAML ファイル、ビルド手順はありません。
14 日間の無料トライアル付きで月額 10 ドル。無制限のリポジトリ、

プルリクエストとチームメンバー。スロットルなし - レビューは即座に開始されます。さらに、独自の Anthropic および/または OpenAI API の使用 (プロバイダーごとのレビューごとに ~0.05 ～ 0.15 ドル)。
他の AI コード レビュー ツールと比較する
直接比較を参照してください。
すべてのツールを並べて、3 者間で比較します。
CodeMouse 対 CodeRabbit — 月額一律 10 ドル対シートごと。
CodeMouse 対 Sourcery — すべての言語対 Python 中心。
CodeMouse と GitHub Copilot — 専用の PR レビュー担当者と IDE ファーストのオートコンプリート。
関連項目: AI コード レビュー ツールの比較 · vs CodeRabbit · vs Sourcery · vs GitHub Copilot · ドキュメント · サポート · 利用規約 · プライバシー

## Original Extract

Automated code review for every GitHub PR — consensus AI reviews from Claude + GPT. Skips what humans or other bots already said. $10/mo, 14-day trial, bring your own keys.

CodeMouse — AI Code Reviews on Every Pull Request
Automated AI code review on every pull request
Install a GitHub App. Bring your Anthropic key, your OpenAI key, or both. Get detailed, context-aware code reviews from Claude and/or GPT — automatically, in minutes.
14-day free trial. $10/mo. No credit card required.
Install the GitHub App — Add CodeMouse to your GitHub organization or personal account. Choose which repos to connect.
Add your AI key(s) — Connect your Anthropic key, your OpenAI key, or both. Connect both and each PR gets reviewed by Claude and GPT in parallel. Keys are encrypted at rest (AES-256-GCM).
Open a pull request — CodeMouse reviews it automatically and posts detailed comments directly on GitHub.
Full repo context — Reviews use the entire repository, not just the diff.
Claude, GPT, or both — Connect one or both. When both are connected we merge the findings and drop duplicates, so you get a single consolidated review.
Bring your own key — No AI cost markup. You pay Anthropic and/or OpenAI directly (~$0.05–0.15 per review per provider).
Reads the room — Reviews every existing PR comment (humans, other bots, earlier CodeMouse runs) and won't repeat what's already been said. Follows up on its own threads when fixes land.
Approves clean PRs — When the diff looks good, prior concerns are addressed, and CI is green, CodeMouse actually approves the PR.
No throttling — Every review starts instantly. No queues, no rate limits.
Works with any language — TypeScript, Python, Go, Rust, Java — if it's code, CodeMouse reviews it.
Inline comments — Review comments appear on specific lines in your PR.
Zero config — No CI pipelines, no YAML files, no build steps.
$10/month with a 14-day free trial. Unlimited repositories, pull requests, and team members. No throttling — reviews start instantly. Plus your own Anthropic and/or OpenAI API usage (~$0.05–0.15 per review per provider).
Compare with other AI code review tools
See the head-to-head comparisons:
All tools side by side — three-way comparison.
CodeMouse vs CodeRabbit — flat $10/mo vs per-seat.
CodeMouse vs Sourcery — all languages vs Python-focused.
CodeMouse vs GitHub Copilot — dedicated PR reviewer vs IDE-first autocomplete.
See also: Compare AI code review tools · vs CodeRabbit · vs Sourcery · vs GitHub Copilot · Docs · Support · Terms · Privacy
