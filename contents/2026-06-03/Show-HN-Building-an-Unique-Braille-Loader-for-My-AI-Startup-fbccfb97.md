---
source: "https://dominikkoch.dev/blog/building-an-unique-braille-loader-for-my-ai-startup"
hn_url: "https://news.ycombinator.com/item?id=48374890"
title: "Show HN: Building an Unique Braille Loader for My AI Startup"
article_title: "Building an Unique Braille Loader for My AI Startup by Dominik Koch"
author: "DominikKoch"
captured_at: "2026-06-03T00:37:05Z"
capture_tool: "hn-digest"
hn_id: 48374890
score: 3
comments: 0
posted_at: "2026-06-02T19:18:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Building an Unique Braille Loader for My AI Startup

- HN: [48374890](https://news.ycombinator.com/item?id=48374890)
- Source: [dominikkoch.dev](https://dominikkoch.dev/blog/building-an-unique-braille-loader-for-my-ai-startup)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T19:18:36Z

## Translation

タイトル: HN を表示: AI スタートアップ向けに独自の点字ローダーを構築する
記事のタイトル: 私の AI スタートアップ向けに独自の点字ローダーを構築する by Dominik Koch
説明: 通常のスピナーでは少し退屈すぎると感じたので、Notra での AI チャット用に小さな点字ローダーを作成しました。
HN テキスト: こんにちは、HN、私は Notra ( https://www.usenotra.com ) の創設者のドミニクです。私たちは最近 AI チャットをアプリに組み込み、そのためにいくつかのカスタム コンポーネントを構築しました。そのうちの 1 つは私たちのローダーで、私たちの会社名をそれに統合するというアイデアがとても気に入ったので、それからコンポーネントを作成し、小さなブログ投稿を作成しました ^^

記事本文:
私の AI スタートアップ向けに独自の点字ローダーを構築する Dominik Koch によるすべての投稿 2026 年 5 月 27 日に公開
AI スタートアップ向けに独自の点字ローダーを構築する
TLDR: コンポーネントのみが必要で、それをアプリで使用する場合は、コンポーネントを追加するコマンドを次に示します。
pnpmyarn npm bun shadcn Copy $ pnpm dlx shadcn@latest add @dominik-ui/braille-loader このページには簡単なデモもあります。
私の会社 Notra では、最近新しい AI チャットを構築しています (これは実際にはかなり優れています)。より面白くするために、クールなローディングアニメーションが必要でした。完璧なローダーを見つけるための探索に私と一緒に参加してください!
点字は、目の見えない人がテキストを読めるようにするために現実世界で使用されています。たとえば、薬には通常、パッケージに点字が印刷されているため、視覚障害者でも自分が手にしている薬が何であるかを実際に理解できます。
これら 4 つの点は点字の小文字 p を表します。
そうは言っても、それは素晴らしいローダーアニメーションにもなります。何かが読み込まれていることを示すために、ある種のカタツムリが円を描くようにすることができるので。これはすでに CLI 内で非常に一般的に使用されています。
でも、普通のスピナーだとかなり退屈ですよね？
点字はすでに実際の文字を表しているため、会社名である notra を詳しく説明するコンポーネントを作成し、それを点字に変換してローダーとして使用できないか考えました。
それで私はそうしました、結果はここで見ることができます:
shadcn CLI を使用して、次のコマンドでこのコンポーネントをインストールできます。
Copy npx shadcn@latest add @dominik-ui/braille-loader これにより、プロジェクトに 2 つの新しいファイル、ui コンポーネント braille-loader.tsx とユーティリティ ファイル text-to-braille.ts が追加されます。
コンポーネントをインストールした後、読み込み状態を表示する場所にコンポーネントをインポートします。
import { BrailleLoader } を "@/components/ui/braille-loader" からコピーします。

次に、他のローダーと同様にレンダリングします。
Copy < BrailleLoader text = "notra" /> テキスト プロップは自動的に点字に変換されるため、独自の製品名、会社名、またはローディング ラベルを使用できます。
Copy < BrailleLoader text = "acme" /> アニメーション スタイルを変更することもできます。
コピー < BrailleLoader テキスト = "notra" バリアント = "wave" />
< BrailleLoader テキスト = "notra" バリアント = "タイプライター" />
< BrailleLoader テキスト = "notra" バリアント = "シマー" />
< BrailleLoader text = "notra"variant = "pulse" /> AI チャットの読み込み状態で使用するのが好きです。
コピー {isLoading ? (
< BrailleLoader テキスト = "notra" バリアント = "wave" />
そして、 className を受け入れるので、他の shadcn コンポーネントと同じようにスタイルを設定できます。
コピー < BrailleLoader text = "notra" className = "text-lg text-primary" /> アクセシビリティ
このローダーは Unicode 点字文字を使用しているため、スクリーン リーダー ユーザーにとっても適切に動作することを確認したいと考えました。
macOS VoiceOver でテストしたところ、点字文字は元のテキストとして読み上げられませんでした。代わりに、VoiceOver が Unicode 文字名をアナウンスするため、ローダーにノイズが発生し、あまり役に立ちません。
アニメーションは装飾的なものであるため、点字文字自体は aria-hidden="true" を使用して支援技術から非表示にする必要があります。このコンポーネントは、周囲の <span> 要素に「Loading」などのアクセス可能な単純なラベルを公開できるため、スクリーン リーダーのユーザーはアニメーション化された文字をすべて聞くことなく実際の状態を取得できます。
これは、Notra の AI チャット用の小さな読み込みアニメーションとして始まりましたが、通常のスピナーよりもユニークに感じられる楽しい小さなコンポーネントに変わりました。実際のテキストを点字に変換することで、アニメーションはあまり気を散らすことなく、ブランドの個性を少し伝えることができます。もしかしたら、これは楽しい小さなイースターエッグになるかも知れません

ユーザーが見つけられるようにしてください;D
独自のアプリで試してみたい場合は、shadcn CLI を使用してインストールし、独自のテキストを渡すことができます。
pnpmyarn npm bun shadcn Copy $ pnpm dlx shadcn@latest add @dominik-ui/braille-loader
私の仕事をサポートしてくださったスポンサーの皆様に感謝します!

## Original Extract

I built a small Braille loader for our AI chat at Notra, because a regular spinner felt a little too boring.

Hello HN, I'm Dominik the founder of Notra ( https://www.usenotra.com ) we recently built an AI Chat into our app and for that built a few custom components. One of them being our loader, I really liked the idea of integrating our company name into it so I made a component out of it and small blog post ^^

Building an Unique Braille Loader for My AI Startup by Dominik Koch All posts Published May 27, 2026
Building an Unique Braille Loader for My AI Startup
TLDR: If you only want the component and use it in your app, here's the command to add it:
pnpm yarn npm bun shadcn Copy $ pnpm dlx shadcn@latest add @dominik-ui/braille-loader You can also find a quick demo on this page
At my company Notra I've recently been building a new AI chat (that’s actually pretty good). To make it more interesting, I wanted a cool loading animation. Join me on my exploration for finding the perfect loader!
Braille is used in the real world for blind people to be able to read text. Medication, for example, usually has Braille imprinted on the packaging so visually impaired people can actually understand what the Medication they are holding is for.
These four dots represent the lowercase letter p in Braille.
That being said it also makes for a great loader animation. Since we can achieve some sort of snail going in a circle to indicate something is loading. This is very commonly used inside CLIs already.
But a plain spinner would be pretty boring, right?
Since braille already represents real letters I thought why don't I make a component that spells out the name of my company, notra, and turn that into braille to use as a loader.
So I did, you can see the result here:
You can use the shadcn CLI to install this component through the following command:
Copy npx shadcn@latest add @dominik-ui/braille-loader This should add two new files to your project, the ui component braille-loader.tsx and the utility file text-to-braille.ts .
After installing the component, import it into the place where you show loading states:
Copy import { BrailleLoader } from "@/components/ui/braille-loader" ; Then render it like any other loader:
Copy < BrailleLoader text = "notra" /> The text prop is converted to Braille automatically, so you can use your own product name, company name, or loading label:
Copy < BrailleLoader text = "acme" /> You can also change the animation style:
Copy < BrailleLoader text = "notra" variant = "wave" />
< BrailleLoader text = "notra" variant = "typewriter" />
< BrailleLoader text = "notra" variant = "shimmer" />
< BrailleLoader text = "notra" variant = "pulse" /> I like using it in AI chat loading states:
Copy {isLoading ? (
< BrailleLoader text = "notra" variant = "wave" />
) : null } And since it accepts className , you can style it like any other shadcn component:
Copy < BrailleLoader text = "notra" className = "text-lg text-primary" /> Accessibility
Because this loader uses Unicode Braille characters, I wanted to make sure it behaves well for screen reader users too.
When I tested it with macOS VoiceOver, the Braille characters were not read back as the original text. Instead, VoiceOver announced the Unicode character names, which makes the loader noisy and not very useful.
Since the animation is decorative, the Braille characters themselves should be hidden from assistive technology with aria-hidden="true" . The component can still expose a simple accessible label like “Loading” on the surrounding <span> element, so screen reader users get the actual state without hearing every animated character.
This started as a small loading animation for our AI chat at Notra, but it turned into a fun little component that feels more unique than a regular spinner. By converting real text into Braille, the animation can carry a bit of brand personality without becoming too distracting. And who knows, maybe this can be a fun little easter egg for your users to spot ;D
If you want to try it in your own app, you can install it with the shadcn CLI and pass in your own text:
pnpm yarn npm bun shadcn Copy $ pnpm dlx shadcn@latest add @dominik-ui/braille-loader
Thank you to all my sponsors for supporting my work!
