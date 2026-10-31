---
source: "https://www.seroundtable.com/google-ai-training-control-42430.html"
hn_url: "https://news.ycombinator.com/item?id=50313575"
title: "Google Adopts New AI Training Control"
article_title: "Google Also Adopts New AI Training Control"
image: "https://images.seroundtable.com/google-web-pages-eoYd2U4x.jpg"
author: "thm"
captured_at: "2026-10-31T06:41:16Z"
capture_tool: "hn-digest"
hn_id: 50313575
score: 1
comments: 1
posted_at: "2026-10-31T06:01:40Z"
tags:
  - hacker-news
---

# Google Adopts New AI Training Control

- HN: [50313575](https://news.ycombinator.com/item?id=50313575)
- Source: [www.seroundtable.com](https://www.seroundtable.com/google-ai-training-control-42430.html)
- Score: 1
- Comments: 1
- Posted: 2026-10-31T06:01:40Z

## Translation

Title: Google Adopts New AI Training Control
Article title: Google Also Adopts New AI Training Control
Description: Google is adopting new AI training controls, under consideration by the IETF, that should simplify and streamline AI training controls for publishers and site owners.

Article text:
Google Also Adopts New AI Training Control
Google
The Pulse of the search community
Subscribe
Subscribe Options
Browse By
Browse by Date
Find by Category
Discover by Author
Scan Most Recent
See Comments
Is it new?
Weekly Stats
View Tag Cloud
Light / Dark Mode Settings
Light
Dark
Use Device Settings
Advertise
Contact Us
Home / Google News / Google SEO / Google Adopts New AI Training Control
Google Adopts New AI Training Control
0
—
by
Barry Schwartz
Barry Schwartz
/ Executive Editor
Barry Schwartz is the CEO of RustyBrick and a technologist, a New York Web service firm specializing in customized online technology that helps companies decrease costs and increase sales. Barry is also the founder of the Search Engine Roundtable and the News Editor of Search Engine Land . He is well known & respected for his expertise in the search marketing industry. He only provides consulting services to expert SEOs and also performs search marketing expert witness services. Barry graduated from the City University of New York and lives with his family in the NYC region. You can follow Barry on Twitter at @rustybrick or on LinkedIn and read his full bio over here .
39,502 Articles as
barry.schwartz
Filed Under
Google Search Engine Optimization
Google also announced , like Bing , OpenAI and Anthropic , that it plans to voluntarily adopt a proposal under consideration at IETF , the global internet standard-setting body, that would create a simple, streamlined AI training control for web publishers and creators.
Google said, "This will help websites express their AI training preferences across multiple AI providers in a single step and is based on work from the AIPREF working group at the Internet Engineering Task Force (IETF)."
This new, work-in-progress control will give websites and publishers a new way to communicate their AI training preferences across multiple AI providers. This lets them invoke their preferences in a single step in robots.txt, instead of listing each one individually.
Google plans to implement this in 2027.
Google told me "moving forward with this initiative alongside peer AI providers as part of a shared commitment to foster an open, creator-first web ecosystem, and following feedback from publishers, civil society and regulators calling for an easier way for websites to express their AI training preferences."
Instead of managing separate instructions for every AI developer, website owners will be able to use a shared 'policy token' within robots.txt to set broad preferences across participating AI providers.
Keep in mind, Google also has the Search Console opt out method for Search generative AI.
Google also updated its common crawlers page to mention future AI training control support. Google added a note that says:
Note: Google intends to support the AIPREF shared machine-readable training preference signal (with "user-agent: __usage_ai-training") in early 2027.
John Mueller, Google Search Advocate, wrote:
When a website uses this policy token to opt out of training in its robots.txt file (and Google-Extended is not explicitly specified), Google plans to interpret that as if the publisher had disallowed Google-Extended. Other AI providers can handle it similarly, making it easy for website owners to opt out across all providers, or to pick and choose.
Existing company-specific controls will continue to function alongside this new joint signal. So website owners can easily express a broad choice using the industry signal, while also tailoring their instructions to individual AI developers like Google through controls such as Google-Extended. Websites that specifically allow Google-Extended are eligible to appear in responses and receive clicks via grounding in Gemini Apps.
The new mechanism brings website owners the benefit of the AI training definition developed at the IETF while multistakeholder discussions continue on the rest of the AIPREF vocabulary. As more industry players adopt this mechanism, it gives website owners early and expanding access to the preference signal.
This draft proposal reflects feedback from publishers, civil society groups, regulators and other stakeholders that have called for industry-wide controls on AI training. Support for this proposal across major AI providers will enable a future standard that is cross-industry, robust, and durable. This proposal is consistent with the EU Code of Practice for General Purpose AI (GPAI) Model Providers under the EU AI Act that recognizes robots.txt as a primary method for web publishers to express machine-readable preferences. We'll continue to work with these stakeholders to refine and standardize this through the IETF working group.
We aim to support this new control along with other AI providers in 2027, and will provide updates on our progress.
This may be helpful from Vanessa Fox in better understanding this change.
I also updated my tool to check for the new user-agent. You can read the announcements here: https://t.co/FdQELlzled &amp; https://t.co/jvnON5wTvj &amp; https://t.co/lDUUXvBQHx &amp; https://t.co/H1m7mPEpvU
Also, it seems like the IETF didn't just have these talks with Google, Microsoft, OpenAI and Anthropic, but also both large and small publishers to chime in:
@rustybrick @gaganghotra_ fyi, Barry, i've been spending the last year deeply participating in this group. you'll notice my name at the bottom in the acknowledgments. if you need help understanding, dm
(also, fyi, it's not actually done at the ietf level -- and some significant areas of dispute remain)
this was our latest (short) interim. https://t.co/Sd8qKgMBJz
here's one (of many) videos from our 3 day meeting in london: https://t.co/oRyBooGaOz
here is the mailing list archive: https://t.co/JagxSD2dSz
ive been one of the more vocal participants. you'll notice at the end of the latest interim recording, and also in my latest emails, that i'm still quite unhappy about one thing in the current draft (it basically exempts the preferences from applying to anything a user does, which hurts artists and creators getting hammered by AI spam copycats). but generally the process has been productive, and it's not over yet.
One note - they didn’t actually invite me or anyone else. They don’t do invites. IETF is a public and transparent group anyone can participate in. Its mostly gated by the fact that it takes a lot of time and effort to meaningfully contribute. I just found out about it and chose to invest the 100s of hours and thousands of $ in travel costs to participate this past year.
The Pulse of the search community
Subscribe
Subscribe Options
Google Search Volatility
More Videos
Subscribe to Videos
Most Recent Articles
Google Adopts New AI Training Control
0
-
by
Barry Schwartz
Bing SEO
Microsoft Bing New AI New Controls With AI Training From AI Search Grounding
0
-
by
Barry Schwartz
ChatGPT Ads
ChatGPT Gets Wider Ads, Multi-Image Ads Expands, Expansion & More
0
-
by
Barry Schwartz
Search Forum Recap
Daily Search Forum Recap: October 30, 2026
0
-
by
Barry Schwartz
Search Video Recaps
Search News Buzz Video Recap: Google Volatility X2, Spam Updates, AI Overviews Everywhere, Blending Ads & Organic, Google & Microsoft Earnings
0
-
by
Barry Schwartz
Google Search Engine Optimization
Google AI Overviews & AI Mode 0.9% Click-Through Rate, Opt Out Study Says
2
-
by
Barry Schwartz
Previous Story:
Microsoft Bing New AI New Controls With AI Training From AI Search Grounding
Submit a Thread
The content at
the Search Engine Roundtable are the sole opinion of the authors and
in no way reflect views of RustyBrick ®, Inc
Copyright © 1994-2026
RustyBrick ® , Inc.
Web Development All Rights Reserved.
This work by Search Engine Roundtable is licensed under a Creative Commons Attribution 3.0 United States License.
Creative Commons License and YouTube videos under YouTube's
ToS .

## Original Extract

Google is adopting new AI training controls, under consideration by the IETF, that should simplify and streamline AI training controls for publishers and site owners.

Google Also Adopts New AI Training Control
Google
The Pulse of the search community
Subscribe
Subscribe Options
Browse By
Browse by Date
Find by Category
Discover by Author
Scan Most Recent
See Comments
Is it new?
Weekly Stats
View Tag Cloud
Light / Dark Mode Settings
Light
Dark
Use Device Settings
Advertise
Contact Us
Home / Google News / Google SEO / Google Adopts New AI Training Control
Google Adopts New AI Training Control
0
—
by
Barry Schwartz
Barry Schwartz
/ Executive Editor
Barry Schwartz is the CEO of RustyBrick and a technologist, a New York Web service firm specializing in customized online technology that helps companies decrease costs and increase sales. Barry is also the founder of the Search Engine Roundtable and the News Editor of Search Engine Land . He is well known & respected for his expertise in the search marketing industry. He only provides consulting services to expert SEOs and also performs search marketing expert witness services. Barry graduated from the City University of New York and lives with his family in the NYC region. You can follow Barry on Twitter at @rustybrick or on LinkedIn and read his full bio over here .
39,502 Articles as
barry.schwartz
Filed Under
Google Search Engine Optimization
Google also announced , like Bing , OpenAI and Anthropic , that it plans to voluntarily adopt a proposal under consideration at IETF , the global internet standard-setting body, that would create a simple, streamlined AI training control for web publishers and creators.
Google said, "This will help websites express their AI training preferences across multiple AI providers in a single step and is based on work from the AIPREF working group at the Internet Engineering Task Force (IETF)."
This new, work-in-progress control will give websites and publishers a new way to communicate their AI training preferences across multiple AI providers. This lets them invoke their preferences in a single step in robots.txt, instead of listing each one individually.
Google plans to implement this in 2027.
Google told me "moving forward with this initiative alongside peer AI providers as part of a shared commitment to foster an open, creator-first web ecosystem, and following feedback from publishers, civil society and regulators calling for an easier way for websites to express their AI training preferences."
Instead of managing separate instructions for every AI developer, website owners will be able to use a shared 'policy token' within robots.txt to set broad preferences across participating AI providers.
Keep in mind, Google also has the Search Console opt out method for Search generative AI.
Google also updated its common crawlers page to mention future AI training control support. Google added a note that says:
Note: Google intends to support the AIPREF shared machine-readable training preference signal (with "user-agent: __usage_ai-training") in early 2027.
John Mueller, Google Search Advocate, wrote:
When a website uses this policy token to opt out of training in its robots.txt file (and Google-Extended is not explicitly specified), Google plans to interpret that as if the publisher had disallowed Google-Extended. Other AI providers can handle it similarly, making it easy for website owners to opt out across all providers, or to pick and choose.
Existing company-specific controls will continue to function alongside this new joint signal. So website owners can easily express a broad choice using the industry signal, while also tailoring their instructions to individual AI developers like Google through controls such as Google-Extended. Websites that specifically allow Google-Extended are eligible to appear in responses and receive clicks via grounding in Gemini Apps.
The new mechanism brings website owners the benefit of the AI training definition developed at the IETF while multistakeholder discussions continue on the rest of the AIPREF vocabulary. As more industry players adopt this mechanism, it gives website owners early and expanding access to the preference signal.
This draft proposal reflects feedback from publishers, civil society groups, regulators and other stakeholders that have called for industry-wide controls on AI training. Support for this proposal across major AI providers will enable a future standard that is cross-industry, robust, and durable. This proposal is consistent with the EU Code of Practice for General Purpose AI (GPAI) Model Providers under the EU AI Act that recognizes robots.txt as a primary method for web publishers to express machine-readable preferences. We'll continue to work with these stakeholders to refine and standardize this through the IETF working group.
We aim to support this new control along with other AI providers in 2027, and will provide updates on our progress.
This may be helpful from Vanessa Fox in better understanding this change.
I also updated my tool to check for the new user-agent. You can read the announcements here: https://t.co/FdQELlzled &amp; https://t.co/jvnON5wTvj &amp; https://t.co/lDUUXvBQHx &amp; https://t.co/H1m7mPEpvU
Also, it seems like the IETF didn't just have these talks with Google, Microsoft, OpenAI and Anthropic, but also both large and small publishers to chime in:
@rustybrick @gaganghotra_ fyi, Barry, i've been spending the last year deeply participating in this group. you'll notice my name at the bottom in the acknowledgments. if you need help understanding, dm
(also, fyi, it's not actually done at the ietf level -- and some significant areas of dispute remain)
this was our latest (short) interim. https://t.co/Sd8qKgMBJz
here's one (of many) videos from our 3 day meeting in london: https://t.co/oRyBooGaOz
here is the mailing list archive: https://t.co/JagxSD2dSz
ive been one of the more vocal participants. you'll notice at the end of the latest interim recording, and also in my latest emails, that i'm still quite unhappy about one thing in the current draft (it basically exempts the preferences from applying to anything a user does, which hurts artists and creators getting hammered by AI spam copycats). but generally the process has been productive, and it's not over yet.
One note - they didn’t actually invite me or anyone else. They don’t do invites. IETF is a public and transparent group anyone can participate in. Its mostly gated by the fact that it takes a lot of time and effort to meaningfully contribute. I just found out about it and chose to invest the 100s of hours and thousands of $ in travel costs to participate this past year.
The Pulse of the search community
Subscribe
Subscribe Options
Google Search Volatility
More Videos
Subscribe to Videos
Most Recent Articles
Google Adopts New AI Training Control
0
-
by
Barry Schwartz
Bing SEO
Microsoft Bing New AI New Controls With AI Training From AI Search Grounding
0
-
by
Barry Schwartz
ChatGPT Ads
ChatGPT Gets Wider Ads, Multi-Image Ads Expands, Expansion & More
0
-
by
Barry Schwartz
Search Forum Recap
Daily Search Forum Recap: October 30, 2026
0
-
by
Barry Schwartz
Search Video Recaps
Search News Buzz Video Recap: Google Volatility X2, Spam Updates, AI Overviews Everywhere, Blending Ads & Organic, Google & Microsoft Earnings
0
-
by
Barry Schwartz
Google Search Engine Optimization
Google AI Overviews & AI Mode 0.9% Click-Through Rate, Opt Out Study Says
2
-
by
Barry Schwartz
Previous Story:
Microsoft Bing New AI New Controls With AI Training From AI Search Grounding
Submit a Thread
The content at
the Search Engine Roundtable are the sole opinion of the authors and
in no way reflect views of RustyBrick ®, Inc
Copyright © 1994-2026
RustyBrick ® , Inc.
Web Development All Rights Reserved.
This work by Search Engine Roundtable is licensed under a Creative Commons Attribution 3.0 United States License.
Creative Commons License and YouTube videos under YouTube's
ToS .
