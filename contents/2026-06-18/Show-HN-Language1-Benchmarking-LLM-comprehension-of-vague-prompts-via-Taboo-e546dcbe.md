---
source: "https://language1.app/"
hn_url: "https://news.ycombinator.com/item?id=48586263"
title: "Show HN: Language1 – Benchmarking LLM comprehension of vague prompts via Taboo"
article_title: ""
author: "kaandemirel"
captured_at: "2026-06-18T16:15:02Z"
capture_tool: "hn-digest"
hn_id: 48586263
score: 1
comments: 0
posted_at: "2026-06-18T14:47:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Language1 – Benchmarking LLM comprehension of vague prompts via Taboo

- HN: [48586263](https://news.ycombinator.com/item?id=48586263)
- Source: [language1.app](https://language1.app/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T14:47:25Z

## Translation

タイトル: HN の表示: Language1 – タブーによるあいまいなプロンプトの LLM 理解のベンチマーク
HN テキスト: こんにちは、HN。LLM に対して「リバース タブー」をプレイするワード ゲーム、Language1 ( https:// language1.app ) を作成しました。仕組み: ターゲット単語 (例: 「リンゴ」) と禁止された「タブー」単語のリスト (例: 「果物」、「赤」、「木」) が与えられます。目標は、禁止された単語を一切使用せずに、LLM が正確なターゲット単語を出力するようにガイドするプロンプトを作成することです。ベンチマークの目標: ゲームプレイ データを使用してベンチマーク データセットを構築する計画でこのプロジェクトを開発しています。目標は、意味論的な制約の下で、不明瞭なプロンプト、比喩、類推、および曖昧な説明を処理する際の LLM 機能をテストして評価することです。ゲームモード: シングルプレイヤー: 一連のチャレンジをプレイして、即時の精度をテストします。試行時間、解決時間、トークン消費量 (標準の cl100k_base エンコーディングで測定) において、他のプレイヤーと世界中で競い合います。登録せずにすぐにプレイしたり、サインイン (ワンクリックで Google ログイン) してスコアをリーダーボードに送信したりできます。
マルチプレイヤー レース: 最大 10 人のプレイヤーが 3 ラウンドをレースするリアルタイム ロビー。注: このゲームは新しいため、最初はパブリック ロビーが空である可能性がありますが、プライベート ロビーを作成して友達とプレイすることができます。
利用可能なモデル: 匿名ユーザーは、デフォルトの Gemma 3 Instruct モデルを使用してプレイします。無料の登録ユーザーは、Llama 3 8B、Liquid LFM 24B、Amazon Nova Micro、Ministral 8B などの複数のモデルから選択して推論スタイルをテストおよび比較できます。技術とガードレール: アプリは React フロントエンドと Node.js/AWS Lambda バックエンドで構築されています。公平性を保つために、入力された手がかりを解析して、文字間隔 (例: "A-P-P-L-E")、翻訳、暗号などの簡単なバイパスをブロックする検証ガードを構築しました。

、base64。モデルをガイドするには、純粋に意味論的推論に依存する必要があります。このゲームは完全に無料で、広告はなく、ブラウザですぐにプレイできます。ゲームプレイについてのあなたの意見を聞き、LLM を導くためにどのような創造的なセマンティック トリックを使用しているかを知りたいと思っています。

## Original Extract

Hi HN, I built Language1 ( https://language1.app ), a word game where you play "reverse Taboo" against an LLM. How it works: You are given a target word (e.g., "Apple") and a list of forbidden "taboo" words (e.g., "fruit", "red", "tree"). Your goal is to write a prompt that guides the LLM to output the exact target word, without using any of the forbidden words. The Benchmark Goal: I am developing this project with the plan of using the gameplay data to build a benchmark dataset. The goal is to test and evaluate LLM capabilities when processing unclear prompts, metaphors, analogies, and vague explanations under semantic constraints. Game Modes: Single Player: Play through a pool of challenges to test your prompt precision. You compete against other players globally across attempts, solve time, and token consumption (measured via standard cl100k_base encoding). You can play instantly without registering, or sign in (one-click Google login) to submit scores to the leaderboards.
Multiplayer Races: Real-time lobbies of up to 10 players racing through 3 rounds. Note: Since the game is new, public lobbies might be empty at first, but you can create private lobbies to play with friends.
Available Models: Anonymous users play with the default Gemma 3 Instruct model. Free registered users can choose between multiple models to test and compare reasoning styles, including Llama 3 8B, Liquid LFM 24B, Amazon Nova Micro, and Ministral 8B. The Tech & Guardrails: The app is built with a React frontend and a Node.js/AWS Lambda backend. To keep things fair, we built a validation guard that parses input clues to block easy bypasses like letter-spacing (e.g., "A-P-P-L-E"), translations, cyphers, and base64. You have to rely purely on semantic reasoning to guide the model. The game is completely free, has no ads, and is playable instantly in the browser. I'd love to hear your thoughts on the gameplay and see what creative semantic tricks you use to guide the LLM!

