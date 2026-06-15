---
source: "https://arxiv.org/abs/2606.13044"
hn_url: "https://news.ycombinator.com/item?id=48536341"
title: "You Can Game AI Peer Review with Presentation-Only Revisions"
article_title: "[2606.13044] No Hidden Prompts Needed! You Can Game AI Peer Review with Presentation-Only Revisions"
author: "ilreb"
captured_at: "2026-06-15T04:37:01Z"
capture_tool: "hn-digest"
hn_id: 48536341
score: 1
comments: 0
posted_at: "2026-06-15T03:40:58Z"
tags:
  - hacker-news
  - translated
---

# You Can Game AI Peer Review with Presentation-Only Revisions

- HN: [48536341](https://news.ycombinator.com/item?id=48536341)
- Source: [arxiv.org](https://arxiv.org/abs/2606.13044)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T03:40:58Z

## Translation

タイトル: プレゼンテーションのみの改訂で AI ピアレビューをゲーム化できます
記事のタイトル: [2606.13044] 隠れたプロンプトは必要ありません。プレゼンテーションのみのリビジョンで AI ピアレビューをゲーム化できる
説明: arXiv 論文 2606.13044 の要約ページ: 隠れたプロンプトは必要ありません。プレゼンテーションのみのリビジョンで AI ピアレビューをゲーム化できる

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.13044
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 6 月 11 日に提出]
タイトル: 隠れたプロンプトは必要ありません!プレゼンテーションのみのリビジョンで AI ピアレビューをゲーム化できる
要約: AI によって生成されたレビューが実験ツールからピアレビュー インフラストラクチャに移行するにつれて、堅牢性に関する懸念のほとんどは、隠された命令やプロンプト インジェクションなどの明示的な攻撃に焦点が当てられています。私たちは、より困難でよりポリシーに関連した障害モードを研究します。つまり、隠しテキスト、プロンプト挿入はなく、方法、実験、数値、方程式、証明、または数値結果の変更はありません。攻撃者は、要約、寄稿の枠組み、関連作業、ディスカッション、物語の構造など、プレゼンテーション レベルのコンテンツのみを変更します。敵対的再パッケージングを導入します。これは、科学的証拠を固定したまま、AI レビュー担当者のフィードバックを使用してプレゼンテーション レベルの改訂を検索する閉ループ攻撃です。 3 人の主流 AI レビュー担当者全体で、敵対的再パッケージ化は 75.1% の攻撃成功率と +1.21/10 の平均スコア増加を達成しました。この効果は、通常の散文の磨きでは説明できません。また、関連作業の再配置や分析的議論の拡張など、査読者が論文を解釈する方法を変える戦略は、局所的な研磨、表の書式設定、アルゴリズムボックスなどの表面的な編集よりも大幅に優れていることも明らかにしました。
私たちの分析により、2 つのより深い構造破壊モードが明らかになりました。まず、AI レビュー担当者は、説得するよりも印象づけるほうが簡単です。長所を強調すると、認識されるメリットが確実に増加しますが、短所を解消しようとする試みは頻繁に行われます。

裏目に出る。第 2 に、AI の査読者は、制限に対処しているように見えることと、実際に制限を解決しているように見えることを混同する可能性があり、変更されていない証拠がより強力な科学的貢献として再解釈される可能性があります。これらの結果は、展開リスクが悪意のある隠された命令だけではなく、紙のプレゼンテーション自体が最適化の表面として出現することであることを示しています。 AI レビュー担当者がプレゼンテーションのみの編集下で科学コンテンツに固定されたままであるかどうかをテストするための、汚染のないローリング ベンチマークと攻撃フレームワークをリリースします。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.13044: No Hidden Prompts Needed! You Can Game AI Peer Review with Presentation-Only Revisions

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.13044
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 11 Jun 2026]
Title: No Hidden Prompts Needed! You Can Game AI Peer Review with Presentation-Only Revisions
Abstract: As AI-generated reviews move from experimental tools into peer-review infrastructure, most robustness concerns have focused on explicit attacks such as hidden instructions and prompt injection. We study a harder and more policy-relevant failure mode: no hidden text, no prompt injection, and no changes to methods, experiments, figures, equations, proofs, or numerical results. The attacker modifies only presentation-level content, such as the abstract, contribution framing, related work, discussion, and narrative structure. We introduce adversarial repackaging: a closed-loop attack that uses AI-reviewer feedback to search for presentation-level revisions while keeping the scientific evidence fixed. Across three mainstream AI reviewers, adversarial repackaging achieves a 75.1% attack success rate and a mean score gain of +1.21/10. The effect is not explained by ordinary prose polishing. We also reveal that strategies that change how the reviewer interprets the paper, such as related-work repositioning and analytical discussion expansion, substantially outperform surface edits such as local polishing, table formatting, and algorithm boxes.
Our analysis reveals two deeper structural failure modes. First, AI reviewers are easier to impress than to convince: highlighting strengths reliably increases perceived merit, while attempts to dissolve weaknesses frequently backfire. Second, AI reviewers can confuse the appearance of addressing a limitation with actually resolving it, allowing unchanged evidence to be reinterpreted as stronger scientific contribution. These results show that the deployment risk is not only malicious hidden instructions, but the emergence of paper presentation itself as an optimization surface. We release a contamination-free rolling benchmark and attack framework for testing whether AI reviewers remain anchored to scientific content under presentation-only edits.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe
