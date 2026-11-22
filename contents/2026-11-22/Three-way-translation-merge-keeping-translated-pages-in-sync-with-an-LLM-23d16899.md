---
source: "https://maxmart.github.io/posts/three-way-translation-merge/"
hn_url: "https://news.ycombinator.com/item?id=50597161"
title: "Three-way translation merge: keeping translated pages in sync with an LLM"
article_title: "Three-way translation merge: keeping translated pages in sync with an LLM · Max's Marginalia"
image: ""
author: "maxmart"
captured_at: "2026-11-22T07:48:10Z"
capture_tool: "hn-digest"
hn_id: 50597161
score: 1
comments: 0
posted_at: "2026-11-22T07:32:56Z"
tags:
  - hacker-news
---

# Three-way translation merge: keeping translated pages in sync with an LLM

- HN: [50597161](https://news.ycombinator.com/item?id=50597161)
- Source: [maxmart.github.io](https://maxmart.github.io/posts/three-way-translation-merge/)
- Score: 1
- Comments: 0
- Posted: 2026-11-22T07:32:56Z

## Translation

Title: Three-way translation merge: keeping translated pages in sync with an LLM
Article title: Three-way translation merge: keeping translated pages in sync with an LLM · Max's Marginalia
Description: A three-way translation merge gives an LLM the changes to the source page since the last sync and lets it edit the existing translation.

Article text:
M ax's M arginalia Posts GitHub Three-way translation merge: keeping translated pages in sync with an LLM
W e recently rebuilt our company website. It exists in English, Swedish and Norwegian and is a static site built with Astro, with pages stored as plain MDX files in a Git repository. As with any multilingual site, you need some method to keep the translated pages in sync between languages. Due to my underlying hubris, defiant curiosity and lack of respect for existing solutions, I wanted to try an idea I’ve had for doing this.
The usual way is to split each page into strings or segments and translate them one by one. Then every language has the same structure, and you correct a translation in the translation tool. I wanted each language to be a whole page that you can edit like any other page. When we edit the English page, we want the Swedish and Norwegian pages to be updated. But that’s very tedious to do by hand. And we can’t re-translate the whole page, because then we might overwrite manual corrections.
What if you gave an LLM the changes made in the source page, and let it edit the existing translated page, the way a coding agent edits a file?
Then it would have the context of the existing text and could perform a semantically equivalent change that wouldn’t overwrite corrected terms. The LLM would get three things:
the changes made in the source page, as a diff
English current old Swedish current stale in sync English current Diff LLM using an edit tool Swedish updated in sync
If you use Git, you may recognize this process. When Git merges two branches, it doesn’t just compare the two versions of a file. It also looks at the last version the two branches have in common, which is called the merge base . That is a three-way merge.
Git needs the base because the two versions alone don’t say who changed what. If a line is in one version and missing in the other, one side may have added it, or the other side may have removed it. With the base, Git can tell which, and it keeps the changes from both sides.
What the LLM does here is similar to a three-way merge. But where an ordinary merge treats both sides alike, here the change only goes one way. Git has an operation for that too, called a cherry-pick. A cherry-pick is a type of three-way merge that takes the change that one commit made and applies it somewhere else:
result = apply(diff(base, theirs), ours)
My idea is the same principle, with an LLM doing the applying:
result = LLM(diff(base, theirs), ours)
Here the base is the old English page, theirs is the new English page, and ours is the current Swedish page. The LLM edits the Swedish page in place instead of producing a new one.
I call this a three-way translation merge : a three-way merge where the two sides are in different languages, and an LLM does the merging.
You can think of each language as a branch of the English page that never merges back. It says the same things, in another language. A sync to one language cherry-picks the new English changes onto that language’s branch, and the base records how far the branch has come. Git can’t do this itself, because it finds the place for a change by matching lines, and on a translated branch none of them match. The LLM finds the place by meaning.
It may be surprising that a diff of the source page is enough to edit a page in another language. But a diff contains the old lines as well as the new ones, and a few unchanged lines around them for context. The LLM uses those to find the place in the target page. For when it needs more than that, it also gets tools to search and read the current source page.
I implemented this idea in my home-built CMS for our company website. Links are at the end of the post, but this post is about the idea and its implications.
To give you an example, this is what it looks like to edit a page in our website admin:
Some background to the video: The word “club” can have two close translations in Swedish: klubb (club) and förening (association). The first machine translation of our page used klubb , because the English page says “club”. But in this context förening is the more appropriate word in Swedish, so we manually corrected the Swedish page to say that instead.
In the video, an editor later changes the heading in the English page, and then the Swedish page is synced:
The diff shows that only “Built-in” changed, so the LLM changes only that word, and föreningar stays. Translating the whole heading again would have given Landsspecifikt register över klubbar .
To make the diff for a language, you need its base, which is the source page as it was when that language was last synced. There is one base for each target language. Since a target page can change both when someone corrects or localizes it and when it is synced, you can’t know from file timestamps or git history alone when the last sync was. Syncing is an intentional action, so you need to record the base explicitly .
There are many ways to record the base, like marking the sync commits or storing a reference like a commit or blob hash, timestamp or explicit version. All have different tradeoffs. I’ve opted to store each base as a plain file, and update it after each sync. That makes it trivial to detect when a language is out of date: its base differs from the current source page. The diff for the LLM is equally trivial: git diff --no-index base.mdx english.mdx .
You can record the base however you want, though a plain file has one special advantage that we will see. And it’s nice to keep it together with the content.
The base files can be stored anywhere. These are two examples of how it could look:
pages/en/pricing.mdx
pages/sv/pricing.mdx
.l10n/sv/pages/en/pricing.mdx
Example: a file next to each page pages/en/pricing.mdx
pages/sv/pricing.mdx
pages/sv/pricing.mdx.base
In both, the file in bold is the base: the English page as it was when Swedish was last synced.
Thanks to git’s content-addressable object storage, when these base files are equal to the current page or an earlier version of it, as they often are, they are already stored in git and don’t take any additional repository space. If the base files clutter your checkout, git’s sparse checkout can keep them out of the working directory, and a tool can read and write them directly as git objects. To collapse them in pull requests on GitHub, mark them linguist-generated in .gitattributes .
Corrections to the source page
You might want to change the English page without propagating that change to the other languages. Correcting a spelling mistake in English is one case. To allow this, we can let the author select if changes should propagate (default) or not. If they shouldn’t, we take the changes and patch each target language’s base. After patching, it will be like the target languages have already “seen” that change and it won’t show up in the diff for the LLM, so it isn’t translated.
For you Git nerds, this is like git merge -s ours , which records the other branch’s changes as merged without taking any of them.
Note that if the target language is behind on some other change, the patched base will be a version of the page that was never committed to git. A commit hash can’t point to that, which is one reason to store the base as a file.
The translated page is its own source of truth
Since the LLM has the entire translated page as context, and only edits the page, your corrections to the translated page are kept . Most translation tools save every translated string in a translation memory , and that is also where you would correct something. A three-way translation merge has no translation memory. You correct the translated page itself. The advantage of a translation memory is that the correction is then reused on other pages too, which this approach doesn’t do yet (see problem 3 ).
In translation tools that work with strings or segments, every language has the same structure. But since the LLM decides where to edit based on meaning, rather than line numbers, segment ids or translation keys, language versions no longer need to have the same structure . You might want a large banner advertising Vipps payments on the Norwegian page only, or the feature list in a different order on your Spanish page. Changes to the English page would still propagate to those pages.
Here is an example. The samples below are built from our homepage, with the names and prices made up, and I’ve shortened them to fit the figure. The lines that change, and the LLM’s notes, are from a real run with Claude Sonnet 5.5.
In English, the licence price is in one place: the pricing section says that the licence costs SEK 95 per team. The Swedish page is localized, and it says 95 in three places:
In its pricing section, like the English page.
In the payments row, which the Swedish page has moved to the top. It has a sentence there that the English page doesn’t have, saying that the payment solution is included in the licence at 95 kr per team.
In a sentence about another product , which happens to cost 95 kr per month.
Then the price is raised from 95 to 110 in English.
The LLM changed the first two prices and left the third alone, as expected.
This is all still very early, and many questions remain. I’m using it for my small company’s website as an experiment. It works well for us, but that’s a single site.
Using an LLM for translating puts a lot of trust in the model. It basically turns the translation step into a prompt engineering or harness engineering problem. That requires its own methods for reliability. As a start, I’ve built a test suite. It has 50 cases, and in 17 of them the Swedish page differs from the English page, like in the example above. Each case has a list of checks that the Swedish page must pass afterwards. I ran each case three times with Claude Sonnet 5.5, with different inputs for the LLM:
Without the diff , the LLM can’t tell a change in English from a difference that the Swedish page has on purpose. It made mistakes both ways: it left out real changes, and it removed things that only the Swedish page has. In the price example , one of the three runs removed the Swedish sentence about the licence, because the English page has no such sentence.
With the diff, all the runs that failed were in one case . The Swedish page mentions Klarna in one more place than the English page does, and then English removes every mention of it. The diff shows where English removed it, but not that Klarna is now gone from the whole page. With the old English page or the search tool, the LLM can check that. The runs with the tools cost about 15 percent less than the runs with the old page, since the LLM used the tools in only 14 of the 150 runs.
One run more or less in the table is chance. When I ran the Klarna case 20 times, the LLM got it right 17 times with the old English page, 18 times with the search tool, and 5 times with only the diff.
So in this test, the diff is what makes it work, and the LLM needs a way to look at the source page when a change is about the whole page.
All the results are published , with the diff, the LLM’s notes and the prompt for every run. Note that I’ve only tested it with one model, on short pages, in English and Swedish.
1. A mistake can look intentional
If the model gets it wrong, this design can make it hard to see. If the LLM accidentally removes a sentence during a sync, the base is still updated, and from then it’s exactly as if it was removed on purpose. The same goes for a change that the LLM skips. After the sync, that change won’t show up in a diff again. I haven’t seen either happen on our website that I know of. In my implementation every sync is a git commit, so someone could review what it changed, but nothing in the system makes you do it.
One could imagine some form of mechanical checks. For example, we could make the edit tool also require an id for which hunk(s) in the source diff the edit is for. And we could have a stop hook, similar to the

[truncated]

## Original Extract

A three-way translation merge gives an LLM the changes to the source page since the last sync and lets it edit the existing translation.

M ax's M arginalia Posts GitHub Three-way translation merge: keeping translated pages in sync with an LLM
W e recently rebuilt our company website. It exists in English, Swedish and Norwegian and is a static site built with Astro, with pages stored as plain MDX files in a Git repository. As with any multilingual site, you need some method to keep the translated pages in sync between languages. Due to my underlying hubris, defiant curiosity and lack of respect for existing solutions, I wanted to try an idea I’ve had for doing this.
The usual way is to split each page into strings or segments and translate them one by one. Then every language has the same structure, and you correct a translation in the translation tool. I wanted each language to be a whole page that you can edit like any other page. When we edit the English page, we want the Swedish and Norwegian pages to be updated. But that’s very tedious to do by hand. And we can’t re-translate the whole page, because then we might overwrite manual corrections.
What if you gave an LLM the changes made in the source page, and let it edit the existing translated page, the way a coding agent edits a file?
Then it would have the context of the existing text and could perform a semantically equivalent change that wouldn’t overwrite corrected terms. The LLM would get three things:
the changes made in the source page, as a diff
English current old Swedish current stale in sync English current Diff LLM using an edit tool Swedish updated in sync
If you use Git, you may recognize this process. When Git merges two branches, it doesn’t just compare the two versions of a file. It also looks at the last version the two branches have in common, which is called the merge base . That is a three-way merge.
Git needs the base because the two versions alone don’t say who changed what. If a line is in one version and missing in the other, one side may have added it, or the other side may have removed it. With the base, Git can tell which, and it keeps the changes from both sides.
What the LLM does here is similar to a three-way merge. But where an ordinary merge treats both sides alike, here the change only goes one way. Git has an operation for that too, called a cherry-pick. A cherry-pick is a type of three-way merge that takes the change that one commit made and applies it somewhere else:
result = apply(diff(base, theirs), ours)
My idea is the same principle, with an LLM doing the applying:
result = LLM(diff(base, theirs), ours)
Here the base is the old English page, theirs is the new English page, and ours is the current Swedish page. The LLM edits the Swedish page in place instead of producing a new one.
I call this a three-way translation merge : a three-way merge where the two sides are in different languages, and an LLM does the merging.
You can think of each language as a branch of the English page that never merges back. It says the same things, in another language. A sync to one language cherry-picks the new English changes onto that language’s branch, and the base records how far the branch has come. Git can’t do this itself, because it finds the place for a change by matching lines, and on a translated branch none of them match. The LLM finds the place by meaning.
It may be surprising that a diff of the source page is enough to edit a page in another language. But a diff contains the old lines as well as the new ones, and a few unchanged lines around them for context. The LLM uses those to find the place in the target page. For when it needs more than that, it also gets tools to search and read the current source page.
I implemented this idea in my home-built CMS for our company website. Links are at the end of the post, but this post is about the idea and its implications.
To give you an example, this is what it looks like to edit a page in our website admin:
Some background to the video: The word “club” can have two close translations in Swedish: klubb (club) and förening (association). The first machine translation of our page used klubb , because the English page says “club”. But in this context förening is the more appropriate word in Swedish, so we manually corrected the Swedish page to say that instead.
In the video, an editor later changes the heading in the English page, and then the Swedish page is synced:
The diff shows that only “Built-in” changed, so the LLM changes only that word, and föreningar stays. Translating the whole heading again would have given Landsspecifikt register över klubbar .
To make the diff for a language, you need its base, which is the source page as it was when that language was last synced. There is one base for each target language. Since a target page can change both when someone corrects or localizes it and when it is synced, you can’t know from file timestamps or git history alone when the last sync was. Syncing is an intentional action, so you need to record the base explicitly .
There are many ways to record the base, like marking the sync commits or storing a reference like a commit or blob hash, timestamp or explicit version. All have different tradeoffs. I’ve opted to store each base as a plain file, and update it after each sync. That makes it trivial to detect when a language is out of date: its base differs from the current source page. The diff for the LLM is equally trivial: git diff --no-index base.mdx english.mdx .
You can record the base however you want, though a plain file has one special advantage that we will see. And it’s nice to keep it together with the content.
The base files can be stored anywhere. These are two examples of how it could look:
pages/en/pricing.mdx
pages/sv/pricing.mdx
.l10n/sv/pages/en/pricing.mdx
Example: a file next to each page pages/en/pricing.mdx
pages/sv/pricing.mdx
pages/sv/pricing.mdx.base
In both, the file in bold is the base: the English page as it was when Swedish was last synced.
Thanks to git’s content-addressable object storage, when these base files are equal to the current page or an earlier version of it, as they often are, they are already stored in git and don’t take any additional repository space. If the base files clutter your checkout, git’s sparse checkout can keep them out of the working directory, and a tool can read and write them directly as git objects. To collapse them in pull requests on GitHub, mark them linguist-generated in .gitattributes .
Corrections to the source page
You might want to change the English page without propagating that change to the other languages. Correcting a spelling mistake in English is one case. To allow this, we can let the author select if changes should propagate (default) or not. If they shouldn’t, we take the changes and patch each target language’s base. After patching, it will be like the target languages have already “seen” that change and it won’t show up in the diff for the LLM, so it isn’t translated.
For you Git nerds, this is like git merge -s ours , which records the other branch’s changes as merged without taking any of them.
Note that if the target language is behind on some other change, the patched base will be a version of the page that was never committed to git. A commit hash can’t point to that, which is one reason to store the base as a file.
The translated page is its own source of truth
Since the LLM has the entire translated page as context, and only edits the page, your corrections to the translated page are kept . Most translation tools save every translated string in a translation memory , and that is also where you would correct something. A three-way translation merge has no translation memory. You correct the translated page itself. The advantage of a translation memory is that the correction is then reused on other pages too, which this approach doesn’t do yet (see problem 3 ).
In translation tools that work with strings or segments, every language has the same structure. But since the LLM decides where to edit based on meaning, rather than line numbers, segment ids or translation keys, language versions no longer need to have the same structure . You might want a large banner advertising Vipps payments on the Norwegian page only, or the feature list in a different order on your Spanish page. Changes to the English page would still propagate to those pages.
Here is an example. The samples below are built from our homepage, with the names and prices made up, and I’ve shortened them to fit the figure. The lines that change, and the LLM’s notes, are from a real run with Claude Sonnet 5.5.
In English, the licence price is in one place: the pricing section says that the licence costs SEK 95 per team. The Swedish page is localized, and it says 95 in three places:
In its pricing section, like the English page.
In the payments row, which the Swedish page has moved to the top. It has a sentence there that the English page doesn’t have, saying that the payment solution is included in the licence at 95 kr per team.
In a sentence about another product , which happens to cost 95 kr per month.
Then the price is raised from 95 to 110 in English.
The LLM changed the first two prices and left the third alone, as expected.
This is all still very early, and many questions remain. I’m using it for my small company’s website as an experiment. It works well for us, but that’s a single site.
Using an LLM for translating puts a lot of trust in the model. It basically turns the translation step into a prompt engineering or harness engineering problem. That requires its own methods for reliability. As a start, I’ve built a test suite. It has 50 cases, and in 17 of them the Swedish page differs from the English page, like in the example above. Each case has a list of checks that the Swedish page must pass afterwards. I ran each case three times with Claude Sonnet 5.5, with different inputs for the LLM:
Without the diff , the LLM can’t tell a change in English from a difference that the Swedish page has on purpose. It made mistakes both ways: it left out real changes, and it removed things that only the Swedish page has. In the price example , one of the three runs removed the Swedish sentence about the licence, because the English page has no such sentence.
With the diff, all the runs that failed were in one case . The Swedish page mentions Klarna in one more place than the English page does, and then English removes every mention of it. The diff shows where English removed it, but not that Klarna is now gone from the whole page. With the old English page or the search tool, the LLM can check that. The runs with the tools cost about 15 percent less than the runs with the old page, since the LLM used the tools in only 14 of the 150 runs.
One run more or less in the table is chance. When I ran the Klarna case 20 times, the LLM got it right 17 times with the old English page, 18 times with the search tool, and 5 times with only the diff.
So in this test, the diff is what makes it work, and the LLM needs a way to look at the source page when a change is about the whole page.
All the results are published , with the diff, the LLM’s notes and the prompt for every run. Note that I’ve only tested it with one model, on short pages, in English and Swedish.
1. A mistake can look intentional
If the model gets it wrong, this design can make it hard to see. If the LLM accidentally removes a sentence during a sync, the base is still updated, and from then it’s exactly as if it was removed on purpose. The same goes for a change that the LLM skips. After the sync, that change won’t show up in a diff again. I haven’t seen either happen on our website that I know of. In my implementation every sync is a git commit, so someone could review what it changed, but nothing in the system makes you do it.
One could imagine some form of mechanical checks. For example, we could make the edit tool also require an id for which hunk(s) in the source diff the edit is for. And we could have a stop hook, similar to the

[truncated]
