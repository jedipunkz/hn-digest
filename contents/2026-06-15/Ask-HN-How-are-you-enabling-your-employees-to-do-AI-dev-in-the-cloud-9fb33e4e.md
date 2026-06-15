---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48543969"
title: "Ask HN: How are you enabling your employees to do AI dev in the cloud?"
article_title: ""
author: "nate"
captured_at: "2026-06-15T19:27:10Z"
capture_tool: "hn-digest"
hn_id: 48543969
score: 2
comments: 2
posted_at: "2026-06-15T16:51:44Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How are you enabling your employees to do AI dev in the cloud?

- HN: [48543969](https://news.ycombinator.com/item?id=48543969)
- Score: 2
- Comments: 2
- Posted: 2026-06-15T16:51:44Z

## Translation

タイトル: HN に聞く: 従業員がクラウドで AI 開発を行えるようにするにはどうすればよいですか?
HN テキスト: 確かに、私たちエンジニアは最近、ラップトップ上でローカルにクロード コードを作成して嵐を起こすことができます。しかし今では、誰もがあらゆるものをバイブコーディングしようとしているため、同様の開発を行うための「適切な」ローカル開発環境を持っていない人がかなりいます。テストスイートを実行してみましょう。私たちの開発者は、それを実行するにはかなり充実した環境を必要とします。したがって、これらの環境はクラウド内にあるのが理想的です。しかし、Claude Code Web は非常に「環境に優しい」ため、実際には代替品ではありません。それとも、Claude Code に大量の依存関係や API キーを Web インストールさせて、そこで作業を行っていますか?現在、私は VM のシステムをデプロイしており、これを人々に使用するよう奨励しています。クロードはこれらのクラウド VM を直接使用して、完全に遅い CI ワークフローを使用せずに、コードの作業やテストの実行などを行うことができるからです。しかし、それはとても壊れやすいようです。皆さんは、従業員が実際にローカルの開発環境を持っていなくても、クロードのような開発環境で従業員をパワーアップするために何を使用していますか?

## Original Extract

Sure, us engineers can Claude Code up a storm locally on our laptops these days. But now with everyone trying to vibe code everything, there's quite a few people that don't have a "proper" local dev environment to do that same kind of development. Let's just take running a test suite. Our devs need a pretty beefy environment to run that. So ideally, these environments are just in the cloud. But Claude Code web, is so "environment lite" that it really isn't a substitute. Or are you all having Claude Code web install a bunch of dependancies and even API keys to do things there? I've got a system of VMs that get deployed now that I'm encouraging folks use, because Claude can just use those cloud VMs directly to work on code, run tests and things without a full slow CI workflow. But it seems so fragile. What are you all using to power up employees at your company with dev environments something like Claude can use without those employees actually having local dev environments?

