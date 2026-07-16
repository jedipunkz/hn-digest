---
source: "https://twitter.com/esrtweet/status/2074889702381953222"
hn_url: "https://news.ycombinator.com/item?id=48935114"
title: "Eric S. Raymond take on AI"
article_title: "Eric S. Raymond on X: \"This is a certain kind of talk around LLMs that I find increasingly puzzling. That is all of the people bitching that LLMs constantly generate crap code and hallucinate solutions, and are worthless for programming.\nThis has almost never happened to me, and never during the las\n[truncated]"
author: "lobo_tuerto"
captured_at: "2026-07-16T15:27:35Z"
capture_tool: "hn-digest"
hn_id: 48935114
score: 3
comments: 1
posted_at: "2026-07-16T14:30:11Z"
tags:
  - hacker-news
  - translated
---

# Eric S. Raymond take on AI

- HN: [48935114](https://news.ycombinator.com/item?id=48935114)
- Source: [twitter.com](https://twitter.com/esrtweet/status/2074889702381953222)
- Score: 3
- Comments: 1
- Posted: 2026-07-16T14:30:11Z

## Translation

タイトル: エリック・S・レイモンドが AI に挑む
記事のタイトル: Eric S. Raymond on X: 「これは、LLM に関するある種の話で、私はますます不可解に感じています。LLM は常にくだらないコードを生成し、幻覚を起こすソリューションであり、プログラミングには価値がないと愚痴っている人はこれだけです。」
このようなことは私にはほとんど起こったことがなく、この間も一度もありませんでした
[切り捨てられた]
説明: これは LLM に関するある種の話で、私はますます不可解に感じています。 LLM は常にくだらないコードを生成し、幻覚を起こして解決策を提示し、プログラミングには価値がないと愚痴っている人はこれだけです。
このようなことは私にはほとんど起こったことがなく、過去にも一度も起こりませんでした

記事本文:
@esrtweet これは LLM に関するある種の話で、私はますます不可解に感じています。 LLM は常にくだらないコードを生成し、幻覚を起こして解決策を提示し、プログラミングには価値がないと愚痴っている人はこれだけです。
このようなことは私にはほとんど起こったことがなく、私が使用した過去 2 世代のモデルでも一度も起こりませんでした (GPT 5.4 と 5.5 について話しましょう)。コンテキストの制限を超えるとモデルが少しおかしくなることがありましたが、コーデックスではそのようなことはもう起こりません。代わりに、制限を超えたためセッションをクリアする必要がある場合、赤で強調表示された警告が表示されました。
私は、C、Go、Rust、Python、シェルで書かれた 63 種類以上のプロジェクトの機能変更、リファクタリング、デバッグに AI を適用してきました。私はそれを使ってドキュメントを書きました。 DOS バイナリを読み取り可能なソース コードに逆コンパイルしました。自分のプロジェクトのいずれかに触れる必要があるときは、まず回帰テストを実行し、次に codex を起動して、コードのバグを監査し、改善を提案するように依頼するのが、今では日常的になっています。
私の経験では、LLM は優れており、非常に強力なツールであると感じています。彼らの最悪の制限は、一種の建築上のトンネルビジョンです。仕様に合わせてコードを生成することには非常に優れていますが、より高いレベルのパターンが見えないことがあります。それは大丈夫です、それをうまくやるのが私の肉脳の仕事です。
LLM について私が発見した最も価値のある点は、詳細や特殊なケースを「台無しにしない」ことです。私は人間の基準からすると非常に優れたプログラマーです (50 年の経験があるので、もっと優れていると思います!) が、LLM は私よりも優れています。なぜなら、コード変更でコード内の (たとえば) 5 か所を変更する必要がある場合、4 か所を修正し、見逃した 5 か所があることに気づくまで何時間もデバッグするという人間の作業を行うことなく、5 か所すべてを確実に見つけることができるからです。
ダウですか

私とは違う宇宙に住んでいるシャウター？古くて性能の悪いモデルを使用しているのでしょうか?それとも、私には適切な精神的習慣とコミュニケーションスキルがあるので、彼らは私には分からない何らかのスキルの問題を抱えているのでしょうか？
[切り捨てられた]
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。

## Original Extract

This is a certain kind of talk around LLMs that I find increasingly puzzling. That is all of the people bitching that LLMs constantly generate crap code and hallucinate solutions, and are worthless for programming.
This has almost never happened to me, and never during the last

@esrtweet This is a certain kind of talk around LLMs that I find increasingly puzzling. That is all of the people bitching that LLMs constantly generate crap code and hallucinate solutions, and are worthless for programming.
This has almost never happened to me, and never during the last two model generations I have used (chat GPT 5.4 and 5.5). Occasionally a model used to get a little deranged when I pushed its context limit, but under codex that doesn't happen anymore; instead I got a red-highlighted warning when the limit has been exceeded and I need to clear my session.
I've applied AI to feature changes, refactoring, and debugging over 63 different projects written in C, Go, Rust, Python, and shell. I've written documentation with it. I've decompiled a DOS binary into readable source code. It's now routine that whenever I have to touch one of my projects I start by running the regression tests, then fire up codex and asking it to audit the code for bugs and suggest improvements.
My experience is that LLMs are excellent and tremendously empowering tools. Their worst limitation is a kind of architectural tunnel vision - they're extremely good at generating code to specification but sometimes blind to higher-level patterns. Which is okay, it's my meatbrain job to be good at that.
The most valuable thing I find about LLMs is exactly that they *don't* screw up details and edge cases. I'm a very, very good coder by human standards (I'd better be, with 50 years of experience!) but the LLMs are better than me. Because if a code change needs to touch (say) five places in the code, they reliably find all five rather than doing the human thing of fixing four and then having to debug for hours before you figure out that there's a fifth one you missed.
Are the downshouters living in a different universe than me? Are they using old, weak models? Or do they have some kind of skill issue that I can't see because I have mental habits and communication skills that are a good fit
[truncated]
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
