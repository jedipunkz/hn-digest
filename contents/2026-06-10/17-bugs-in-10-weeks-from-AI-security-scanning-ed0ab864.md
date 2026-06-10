---
source: "https://lalitm.com/post/perfetto-security-bugs-ai/"
hn_url: "https://news.ycombinator.com/item?id=48479239"
title: "17 bugs in 10 weeks from AI security scanning"
article_title: "17 bugs in 10 weeks from AI security scanning - Lalit Maganti"
author: "lalitmaganti"
captured_at: "2026-06-10T19:01:57Z"
capture_tool: "hn-digest"
hn_id: 48479239
score: 2
comments: 0
posted_at: "2026-06-10T16:58:29Z"
tags:
  - hacker-news
  - translated
---

# 17 bugs in 10 weeks from AI security scanning

- HN: [48479239](https://news.ycombinator.com/item?id=48479239)
- Source: [lalitm.com](https://lalitm.com/post/perfetto-security-bugs-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T16:58:29Z

## Translation

タイトル: AI セキュリティ スキャンから 10 週間で 17 件のバグが発見
記事のタイトル: AI セキュリティ スキャンから 10 週間で 17 件のバグ - Lalit Maganti
説明: ソフトウェアのロングテールは、これまではできなかったセキュリティ上の注目をついに集めています。

記事本文:
ラリット・マガンティ
× RSSニュースレターについて
AI セキュリティ スキャンで 10 週間で 17 件のバグが発見
過去数週間にわたり、次のセキュリティ バグ レポートをさらに受け取りました。
Perfetto のトレース プロセッサは、これまでにないもので、すべて AI によって検出されました。そして
とても嬉しいです！これらは、ほぼ確実に存在しないであろうバグです。
1年前に発見されたので、抜け穴を塞ぐのは良いことだとは思いますが、
トレースプロセッサは決してセキュリティ上重要なものではありません。
セキュリティ研究者は長年にわたり、最も危険な問題に時間を集中してきました。
ターゲット: カーネル、暗号化ライブラリ、パスワード マネージャー。でもたくさんあるよ
セキュリティに関連しているものの、真にセキュリティクリティカルではないコードが大量に存在します。で
私の経験では、この種のプロジェクトはあまり注目を集めませんでした。今のシステム
ロングテールの人々は、以前は得られなかった注目を集めることができます。
Trace プロセッサは、まさにそのロングテールに位置するプロジェクトです。 C++ です
ライブラリ (はい、今日では Rust が当然の選択肢でしょうが、現実的ではありません)
書き換え、さまざまな記録されたトレースの処理については脚注 1 を参照してください。
フォーマット。これらは通常、自分で収集したトレース、またはテスト インフラストラクチャで収集したトレースです。
オフラインで処理するため、「信頼できない入力」はあまり問題になりません。
ただし、自分で収集したわけではないトレースを処理する人もいます (例:
ユーザーバグレポート、ドッグフードユーザーからの自動収集）。そういった場合に備えて
サンドボックス トレース プロセッサを強く推奨してきました (例:
ジバイザー、
サンドボックス2 、または
minijail ) または、さらに機密性の高い使用のために
場合、VM。
サンドボックスを超えて、問題を積極的に検出するために、私たちは主にファジングに依存しました。
Google の内部で実行されています。これらのファザーは時々現実に現れますが、
対処可能なバグ: 任意のトレース バイトを渡すように設定しました (これは
主要な「攻撃対象領域」

）しかし時間が経つにつれて、彼らが発見したように、これらは非常にまれになりました
簡単に解決できる問題の多くはすぐに修正されました。残ったバグ
内部の奥深くに生息する傾向があり、非常に精密に作られたものでのみ到達可能です。
ファザーが突然変異だけでヒットする可能性が低いバイトのシーケンス。
これを除けば、人間のための帯域幅やリソースはほとんどありません。
セキュリティの専門家か私のチームの誰かに、時間をかけて見つけてもらいましょう
トレースプロセッサのセキュリティ問題。 Perfetto には常に他の部分がありました
セキュリティに時間を費やす価値がある (例: トレース サービス、オンデバイス)
プロファイラー) は実稼働システムでアクティブに実行されているためです。
数か月前からすべてが変わりました。バグを受け取り始めました
AI ベースのセキュリティを実行していると思われる中央チームによって提出されたレポート
Google 全体のさまざまなプロジェクトに対してスキャンします。残念ながら、私はそうでなければなりません
彼らの仕事は実際に何をしているように見えないので、手を振りながら
公共の。
4月初旬から、週に1匹のペースでゆっくり点滴していましたが、終わりから
4 月以降、これは週に数回の割合に増加し、日によっては 3 回または 4 回の日もありました。
4 が立て続けにオープンされます。この状況は5月中旬まで続き、その時点で
週に1～2回に減り始め、何もない週もあった。
バグのクオリティが高いとも言いたいです。それらは詳しく説明されており、
多くの場合、関連する攻撃者モデルがすでに確立されており、最小限の修正さえ行われます。
提案: 基本的にバグレポートから求められるすべてのもの。これは一致します
カールと
Linuxカーネル
メンテナは受け取ったセキュリティ バグ、特にそのバグについてメモしています。
ここ数か月で品質が大幅に向上しました。
私に対して報告されたバグのみが表示されるため、生の出力は表示されません。
AIスキャナー、いいですね

上流でどれだけのトリアージが行われているか正確にはわかりません。私の推測は
報告が届く前に明らかなノイズを除去するために人間が軽いパスを行っている
クライアント チームですが、バグが公開される割合と方法から判断すると、
それらはファイルされていますが、それぞれを徹底的にトリアージしている人はいないでしょう。
合計で 21 件のバグを受け取りました (実際の問題が 17 件、対処不可能なものが 4 件)。
これは次のカテゴリに分類できます。
10 個の境界チェック: 固定サイズのバッファに流入する任意のトレース データ
または配列インデックスがチェックされていないため、範囲外の読み取りまたは書き込みが発生します。
5 use-after-free : バックポインター、ポインタースナップショット、またはハッシュマップキー
彼らが参照するオブジェクトよりも存続します。
1 スタック オーバーフロー: 入力が深くネストされている場合の無制限の再帰。
1 アクセス制御 : 一部のまれなコードパスではホワイトリストを強制しません。
4 対処不可能としてクローズ: 純粋にエクスプロイトの可能性があった場合
仮説、または修正するには根本的な設計変更が必要になる場合
わずかなセキュリティリスクを冒す価値はありませんでした。
17 件の実際の問題はすべて修正されており、ほぼすべてが Perfetto v56.0 で出荷されます。
２．
実際にこれらのレポートを受け取るとどう感じますか?まあ、あなたが思っているほど悪くはありません
考えてください。 OpenSSL やcurl のようなセキュリティ クリティカルなアプリケーションとは異なり、トレースでは
プロセッサ、セキュリティ上の問題が P0 である可能性は非常に低いため、削除する必要があります
修正するすべて。誤解しないでください、それは依然として優先事項ですが、
正しい答えを見つけるのに数日かかる余裕があり、
急いで修正をリリースしようとするのではなく、通常のスケジュールに従って修正をリリースします。
CVE を実行し、全員にすぐにパッチを適用させます。
また、ありがたいことに、問題の大部分は機械的なものであるため、修正されます。
一般に非常に簡単です。
たとえば、次の PR を考えてみましょう。
いくつかのキー文字列を解析しながら、固定サイズのスタック バッファーにキー文字列を構築します。
メタ

データ。
境界チェックはデバッグ ビルドでのみ実行され、メタデータ名が追加されます。
跡から真っすぐ。十分に長い名前を付けると、
バッファ。
修正は、スタック バッファーを std::string に交換するだけです。の
コード パスは非常にコールドであるため (トレース内で 1 回か 2 回のみ)、追加のヒープ
割り当ては関係ありません。
実際、この種の問題は非常に機械的なので、私はコーディング エージェントを信頼しています。
最小限のガイダンスでそれらを修正するだけです。よく書かれたレポートを取得して、それをフィードします。
エージェントに問い合わせると、約 10 分以内に 10 ～ 20 行の PR で問題が解決されます。私は
すべての行を徹底的に見直して、理解していることを確認しますが、これらのタスクは
難しいことではなく、AI ができることの「ギザギザのフロンティア」の内側にしっかりと入っています。
ただし、すべての問題が機械的なものではない、または AI に任せられるわけではないことを強調したいと思います。
いくつかのレポートは、実際には機能の誤りよりも設計の問題を指摘しています。
実装。
無料後のこの使い方は良いです
例:
問題は、状態オブジェクトがバックポインタを保持しており、それが最終的にポイントする可能性があることでした。
トレースに特定のデータが表示されると、メモリが解放されます。
即時の未解決のケースは、コールバックを使用するだけで簡単に修正できました。
フリー時のバックポインタを無効にしました。しかし、これは恐ろしいハックです。
関係するオブジェクトの寿命を推論することは不可能です。
ここでの本当の問題は、親が移動できる子オブジェクトがあることです。
このコードが適切であれば、実際には起こらないはずです
建築された。
それを適切に修正するということは、所有権モデルを再構築することを意味するため、寿命は長くなります。
構造上正しくなります。
興味深いのは、これが私が認識していた問題であり、私が抱えていた問題だったということです。
1年近く大掃除をしようと思っていましたが、結局実行できませんでした。
セキュリティのバグは、私にそれを実行する後押しと正当化を与えただけです。このアプリ

に重ねた
他にもいくつかのバグがあり、セキュリティ上の問題が起こり得ることを認識させられました。
場合によっては、より深い設計上の欠陥やハッキーなコードと相関関係があるため、
「セキュリティ スキャン」には、直接的なバグを発見するだけでなく、より幅広いメリットがあります。
私が警戒しているのは、このバグの流れがいつまで続くかということです。感じています
ほんの数か月しか続いていないことを考えると、それは良いことですが、私はできる
これがさらに数カ月続けば、次のような事態になることは容易に想像できます。
精神的に疲れる。
しかし、私の疑念は、これがゼロになるのではないかということです。なぜ？それはパターンに関係する
これらのバグがどのように報告されているかを説明します。コードベースの各部分は次のようになります。
次の作業に移る前に、1 ～ 2 日注意を払って (および関連するバグに) 対処してください。
違う部分。再発はまれで、特に次の地域ではバグのペースが遅くなりました。
過去数週間: 5 月の初めにはもっとたくさんありました (週に数回)
でも今は週に1～2件に減りました。ファイルの数は有限なので、
結局、私の直感はそれらがなくなるだろうと言います。
重要な考慮事項は、スキャナーよりも早く新しいバグを追加するかどうかです。
古いものを見つけることができます。私の疑いはノーです。これまでのところ 17 件の実際の問題は次のとおりです。
9 年間の開発期間をスキャンしました。たとえその数が以前の3倍になったとしても
状況が落ち着いても、スキャナは何年にもわたって蓄積されたコードを通じて動作し続けます。
そして、プロジェクトの初期には、より多くのコードをより速く書きました。
そのため、現在、新しいコードが追加される割合は以前よりも低くなりました。
もう一つの問題は、新しいモデルのリリースではより複雑なデザインが採用されるかどうかです。
今日私たちが見つけている単純な問題ではなく、問題です。それらはかかります
修正するのにかなり多くの時間と労力がかかるため、次のような場合はさらに苦痛になります。
私たちはそれらの多くを手に入れることになっていました。これについては非常に自信がないので、待つしかありません

そして見てください！
ダニエル・ステンバーグ (カール管理者) がこの言葉をうまく表現していると思います。
ポスト:
AI を活用したツールでソース コードをスキャンしていないプロジェクト
膨大な数の欠陥、バグ、潜在的な脆弱性が見つかる可能性があります。
この新世代のツール。
これは私にとって非常に真実に思えます。もっと大まかに言うと、人々は次の 3 つのうちの 1 つを持っていると思います。
経験:
信頼できない入力 + セキュリティクリティカル (curl、カーネル、OpenSSL など): 多数
他のカテゴリよりも誤検知率が高い、複雑なレポート。
なぜなら、このプロジェクトには多くの注目が集まっており、簡単に解決できる問題もたくさんあるからです。
重要なコードパスではすでに成果が得られているでしょう。でも
あまり使用されていない機能 (レガシー ドライバーなど) のコードパスは、
代わりにカテゴリ 2。
信頼できない入力 + 以前に監査されていない (例: トレース プロセッサ): の波
機械的なバグは管理可能なペースで解決され、個人のストレスは低くなります。
プロジェクトはセキュリティ上重要なコード パス上にありません。ここはダニエル両方がいる場所です
そして私は、AI セキュリティ スキャナーが最も大きな影響を与えると予想しています。
信頼できない入力はありません (内部ツール、数学ライブラリ、上でのみ動作するもの)
信頼できるデータ）: この変化にはおそらくまったく気付かないでしょう。
私自身のケースは、その 2 番目のバケツにきちんと収まっています。でもしたくない
私の経験から過度に一般化しますが、理由は 3 つあります。
これは私にとっては対処できることですが、すべての人にとっては当てはまらないでしょう: a) 私は報酬をもらっています
フルタイムの仕事の一環としてトレースプロセッサを保守します。 b) 他の人が服用している
AI スキャンを実行して最初にバグを発見する取り組み。 c)
レポートは上流の人間のレビュー担当者によって軽くフィルタリングされているようです。
明らかなノイズは取り除きますが、おそらく詳細なトリアージは行いません。
私にとって、これはエコシステムのギャップを示しています。ほとんどのオープンソース プロジェクトは、
専用のお茶を飲む余裕がある

私は彼らのためにセキュリティスキャンを行っており、
メンテナは、セキュリティ リスクがわずかな場合に独自のパイプラインを立ち上げることができます。
これは、最も意欲的なプロジェクトのみに制限されます。私たちはそうだと思います
この分野では、巨大なAIを含め、さらに多くのイノベーションが起こるだろう
研究室。
全体として、私は自分自身の経験について、慎重ながらも肯定的です。つまり、ほとんどのバグは
機械的なものもありますが、長年懸案だった設計のクリーンアップを少しずつ進めているものもあり、そのペースは
管理可能。これがどのように進展するのか、私には分からないことがたくさんあります。
将来のモデルがより困難な設計上の問題を発見し始めるかどうかは、当てはまります。したがって、これは
予測ではなく、スナップショットとして扱う必要があります。
この投稿を気に入っていただけた場合は、ニュースレターを購読するか、 RSS 経由でフォローすることをご検討ください。 Hacker News や Lobsters で共有することもできます。
私が予想する一般的な応答は次のとおりです。「次のような任意のバイナリ データを解析している場合」
トレースは Rust で書かれている必要があります。」真空状態では私も同意します、そしてもし私がそうであれば
今日、トレースプロセッサをゼロから作成するなら、私は間違いなくRustを使用するでしょう。しかし、
残念ながら、Rust への切り替えはまったく現実的ではありません。図書館は
かなりの量のコードがあり、何百ものダウンストリーム ツールに埋め込まれています。
多くはRustツールチェーンを持たない環境にあります。私たち全員に尋ねると、
エンベッダーが Rust を使い始めるのはかなりの負担になるでしょうが、私はそうではありません

[切り捨てられた]

## Original Extract

The long tail of software is finally getting the security attention it never could before.

Lalit Maganti
× About RSS Newsletter
17 bugs in 10 weeks from AI security scanning
Over the last several weeks, I’ve been receiving more security bug reports for
Perfetto’s trace processor than I ever have before, all of them found by AI. And
I’m very happy about it! These are bugs that would almost certainly not have
been found a year ago and it feels good to close these loopholes even though
trace processor is by no means security critical.
For years, security researchers concentrated their time on the highest-stakes
targets: kernels, cryptography libraries, password managers. But there’s a lot
of code out there which is security-relevant but not truly security-critical. In
my experience, these sorts of projects didn’t draw much attention. Now systems
in the long tail can get that attention which they wouldn’t have before.
Trace processor is a project which sits squarely in that long tail. It’s a C++
library (yes, Rust would be the obvious choice today but it’s not practical to
rewrite, see footnote 1 ) for processing recorded traces of various
formats. These are typically traces you collected yourself or in your test infra
and process offline so “untrusted input” isn’t much of a concern.
However, some people do process traces they didn’t collect themselves (e.g.
user bug reports, automated collection from dogfood users). For those cases
we’ve strongly recommended sandboxing trace processor (e.g.
gvisor ,
sandbox2 , or
minijail ) or, for even more sensitive use
cases, a VM.
Beyond sandboxing, for catching issues proactively, we mainly relied on fuzzing
running internally in Google. These fuzzers occasionally surfaced real,
actionable bugs: we set them up to pass in arbitrary trace bytes (as this is the
main “attack surface”) but over time these became quite rare as they discovered
much of the low hanging fruit, which we quickly fixed. The bugs that remained
tend to live deep in the internals, reachable only with a very precisely crafted
sequence of bytes that a fuzzer is unlikely to hit by mutation alone.
Apart from this, there has rarely been any bandwidth or resources for a human,
either a security expert or someone from my team, to spend lots of time finding
security issues in trace processor. There were always other parts of Perfetto
more worth spending security time on (e.g. the tracing service, on-device
profilers) as they’re actively running in production systems.
All of this changed as of a couple of months ago. We started receiving bug
reports filed by some central team which appears to be running AI-based security
scanning against various projects throughout Google. Unfortunately, I have to be
hand wavy about what exactly they’re doing as their work doesn’t appear to be
public.
Starting in early April, we had a slow drip of 1 bug a week, but since the end
of April this increased to a rate of several a week, with some days having 3 or
4 being opened in quick succession. This lasted until mid-May, at which point it
started tapering back to 1-2 a week with some weeks having none.
I also want to say that the quality of the bugs is high. They’re well-described,
often with the relevant attacker model already worked out and even minimal fixes
proposed: basically everything I could ask for from a bug report. This matches
what both curl and
Linux kernel
maintainers have noted about security bugs they’ve received, especially how
sharply quality has improved in the last few months.
As I can only see the bugs that get filed against me, not the raw output of the
AI scanner, I don’t know exactly how much triage happens upstream. My guess is
there’s a human doing a light pass to drop obvious noise before reports reach
client teams, but judging from the rate at which bugs are opened and the way
they’re filed, I doubt anyone is deeply triaging each one.
In total, we’ve received 21 bugs (17 real issues and 4 not actionable),
which can be broken down into the following categories:
10 bounds checking : arbitrary trace data flowing into fixed-size buffers
or unchecked array indices, leading to out-of-bounds reads or writes.
5 use-after-free : back-pointers, pointer snapshots, or hashmap keys
outliving the object they refer to.
1 stack overflow : unbounded recursion when input is deeply nested.
1 access control : not enforcing allowlists on some rare codepaths.
4 closed as not actionable : either where the chance of exploit was purely
hypothetical or where fixing would have required fundamental design changes
which were not worth the tiny security risk.
All 17 real issues have been fixed, almost all shipping in Perfetto v56.0
2 .
How does receiving one of these reports actually feel? Well not as bad as you’d
think. Unlike a security-critical application like OpenSSL or curl, in trace
processor, a security issue is very unlikely to be a P0 I have to drop
everything to fix. Don’t get me wrong, it’s still a priority but one where I
have the luxury of taking a few days to figure out the right answer and can
release fixes according to our normal schedule, instead of trying to rush out a
CVE and get everyone to patch immediately.
Also thankfully, because the majority of the issues are mechanical, the fixes
are generally quite straightforward.
Take this PR , for example:
We build a key string into a fixed-size stack buffer while parsing some
metadata.
The bounds check only runs in debug builds, and the metadata name comes
straight from the trace. Putting a long enough name means you would escape the
buffer.
The fix is a simple matter of swapping the stack buffer for a std::string. The
code path is very cold (only once or twice in a trace) so the extra heap
allocation doesn’t matter.
In fact, these sorts of issues are so mechanical that I trust a coding agent
to just fix them with minimal guidance: take the well-written report, feed it to
the agent, and within ~10 minutes there’s a 10-20 line PR which fixes it. I
review every line thoroughly and make sure I understand it, but these tasks are
not difficult and firmly inside the “jagged frontier” of what AI can do.
I want to stress though that not every issue is mechanical or can be left to AI;
a few reports actually point more to design problems than incorrect function
implementations.
This use after free is a good
example:
The problem was a state object held a back-pointer that could end up pointing
to freed memory given certain data appearing in the trace.
The immediate dangling case was easy to patch by just having a callback which
invalidated the back-pointer on free. But this is a horrible hack which makes
the lifetimes of the objects involved impossible to reason about.
The real problem here is that you had a child object whose parent could go
away before it, which really shouldn’t happen if this code is properly
architected.
Fixing it properly meant restructuring the ownership model so the lifetime was
correct by construction.
The interesting thing was that this was a problem I was aware of and that I had
been meaning to clean up for close to a year but never got round to: the
security bug just gave me the push and justification to do it. This applied in a
couple of other bugs as well and made me internalize that security issues can
sometimes be correlated with deeper design flaws or hacky code so there are
wider benefits to “security scanning” than just the direct bugs they find.
One thing I am wary of is how long this stream of bugs will keep up; I’m feeling
good about it given it’s only been going on for a couple of months, but I can
easily imagine that if this goes on for several more months, it might become
mentally exhausting.
But my suspicion is that this will go to zero. Why? It’s to do with the pattern
of how these bugs are being filed. Each part of the codebase seems like it’s
getting a day or two of attention (and associated bugs) before moving on to a
different part. Repeats are rare, and the pace of bugs has slowed especially in
the last couple of weeks: we had a lot more in the start of May (several a week)
but now we’re down to 1-2 a week. There are a finite number of files, so
eventually my gut tells me they will run out.
An important consideration is whether we’ll add new bugs faster than the scanner
can find old ones. My suspicion is no; the 17 real issues so far are from
scanning across 9 years of development. Even if that number triples before
things settle, the scanner is still working through years of accumulated code.
And we wrote a lot more code, a lot faster, in the earlier years of the project,
so the rate of new code being added now is lower than it once was.
The other question is whether new model releases will find more complex design
issues rather than the simple issues we’re finding today. Those take
significantly more time and effort to fix and so would be a lot more painful if
we were to get many of those. I’m very unsure on this so we’ll just have to wait
and see!
I feel Daniel Stenberg (curl maintainer) phrased it well in this
post :
Any project that has not scanned their source code with AI powered tooling
will likely find huge number of flaws, bugs and possible vulnerabilities with
this new generation of tools.
This rings very true to me. More broadly, I think folks will have one of three
experiences:
Untrusted input + security critical (e.g. curl, kernel, OpenSSL): many
complex reports, with a higher false positive rate than the other categories,
because there’s a lot of attention on the project and much of the low hanging
fruit would already have been picked in the critical codepaths. Though
codepaths for lesser-used functionality (e.g. legacy drivers) could end up in
category 2 instead.
Untrusted input + not previously audited (e.g. trace processor): a wave of
mechanical bugs at a manageable pace and low individual stress because the
project is not on a security critical code path. This is where both Daniel
and I expect AI security scanners to have the most impact.
No untrusted input (internal tools, math libs, anything operating only on
trusted data): you probably won’t notice this shift at all.
My own case sits squarely in that second bucket. But I don’t want to
over-generalize from my experience, because there are three things that make
this manageable for me that wouldn’t be true for everyone: a) I’m paid to
maintain trace processor as part of my full time job; b) someone else is taking
the effort to run the AI scans and discover the bugs in the first place; c) the
reports appear to be lightly filtered by an upstream human reviewer, enough to
strip obvious noise but probably not a deep triage.
To me, this points to a gap in the ecosystem: most open-source projects cannot
afford to have a dedicated team doing security scanning for them, and telling a
maintainer to stand up their own pipeline when their security risk is marginal
will restrict this to only the most motivated projects. I would guess we’re
going to see a lot more innovation in this space, including from the big AI
labs.
All in all, I’m cautiously positive about my own experience: most of the bugs
are mechanical, a few have nudged long-overdue design cleanups, and the pace is
manageable. There’s plenty I don’t know about how this evolves: whether the pace
holds, whether future models start finding harder design issues. So this should
very much be treated as a snapshot, not a forecast!
If you enjoyed this post, consider subscribing to my newsletter or following via RSS . You can also share it on Hacker News or Lobsters .
A common response I expect is: “if it’s parsing arbitrary binary data like
traces, it should be written in Rust.” In a vacuum I agree and if I was
writing trace processor from scratch today, I would definitely use Rust. But
switching to Rust is unfortunately quite impractical; the library is a
significant amount of code and is embedded in hundreds of downstream tools,
many in environments that don’t have a Rust toolchain. Asking all our
embedders to start using Rust would be a significant burden and one I don’t

[truncated]
