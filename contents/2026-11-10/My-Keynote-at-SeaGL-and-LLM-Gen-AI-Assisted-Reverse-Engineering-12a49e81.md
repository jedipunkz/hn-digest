---
source: "https://sfconservancy.org/blog/2026/nov/05/foss-in-the-age-of-llm-gen-ai-assisted-reverse-eng/"
hn_url: "https://news.ycombinator.com/item?id=50445077"
title: "My Keynote at SeaGL and LLM-Gen-AI-Assisted Reverse Engineering"
article_title: "My Keynote at SeaGL and LLM-gen-AI-assisted Reverse Engineering - Conservancy Blog - Software Freedom Conservancy"
image: "https://sfconservancy.org/static/img/conservancy-logo.png"
author: "hn_acker"
captured_at: "2026-11-10T20:16:28Z"
capture_tool: "hn-digest"
hn_id: 50445077
score: 1
comments: 0
posted_at: "2026-11-10T19:28:57Z"
tags:
  - hacker-news
---

# My Keynote at SeaGL and LLM-Gen-AI-Assisted Reverse Engineering

- HN: [50445077](https://news.ycombinator.com/item?id=50445077)
- Source: [sfconservancy.org](https://sfconservancy.org/blog/2026/nov/05/foss-in-the-age-of-llm-gen-ai-assisted-reverse-eng/)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T19:28:57Z

## Translation

Title: My Keynote at SeaGL and LLM-Gen-AI-Assisted Reverse Engineering
Article title: My Keynote at SeaGL and LLM-gen-AI-assisted Reverse Engineering - Conservancy Blog - Software Freedom Conservancy
Description: The Software Freedom Conservancy provides a non-profit home and services to Free, Libre and Open Source Software (FLOSS) projects.

Article text:
Get the latest update on our Vizio court case -->
Toggle navigation menu
News
Press Releases
We defend & uphold the rights enshrined in copyleft licenses.
We defend the legal rights of software users in Court when needed.
We help FOSS contributors navigate this ongoing conundrum.
Our annual community-oriented conference.
Leave GitHub today ! Learn why & how.
Internships for those who face under­representation, systemic bias, or discrimination in tech.
We provide non-profit infrastructure and services to our members creating Free/Libre and Open Source Software.
Our tool for evaluating the source code candidates companies must provide for GPLed software.
We designed and built the first ever wireless Internet router designed with software freedom and right to repair in mind.
Learn
The Corresponding Source
A bi-weekly oggcast about legal, policy, and many other issues in the Free, Libre, and Open Source Software (FLOSS) world.
A list of terms you might be unfamiliar with but occur frequently in our work.
Your most frequently asked questions about the Vizio lawsuit, answered in one place.
We defend & uphold the rights enshrined in copyleft licenses.
We defend the legal rights of software users in Court when needed.
We help FOSS contributors navigate this ongoing conundrum.
Our annual community-oriented conference.
Leave GitHub today ! Learn why & how.
Internships for those who face under­representation, systemic bias, or discrimination in tech.
We provide non-profit infrastructure and services to our members creating Free/Libre and Open Source Software.
Our tool for evaluating the source code candidates companies must provide for GPLed software.
We designed and built the first ever wireless Internet router designed with software freedom and right to repair in mind.
Learn
The Corresponding Source
A bi-weekly oggcast about legal, policy, and many other issues in the Free, Libre, and Open Source Software (FLOSS) world.
A list of terms you might be unfamiliar with but occur frequently in our work.
Your most frequently asked questions about the Vizio lawsuit, answered in one place.
My Keynote at SeaGL and LLM-gen-AI-assisted Reverse Engineering
by Denver Gingerich
on November 5, 2026
Over the past few months, there have been an increasing number of reports of people successfully gaining access to their own devices that were designed to prevent them from installing their own updates and modifications. Nearly every one of these reports notes that an LLM-gen-AI was used. While there are obviously a huge number of issues with LLM-gen-AIs (as we note in our LLM-gen-AI recommendations ), it is hard to ignore the benefits to people's digital autonomy that reverse engineering like this brings.
As an example, one device that was previously difficult to install updates on was recently freed by OpenWrt ( part of SFC ). Specifically, the Zyxel XS1930 (a 10- or 12-port network switch, with dual-core 1.0 GHz MIPS CPU), which ships with a fully proprietary OS called ZyNOS, has been upgraded to OpenWrt using the ZyNOS web UI, which only recently became feasible, thanks to the efforts of developers using LLM-gen-AI. Previous efforts to install OpenWrt or other Linux-based systems on the XS1930 required hardware modifications, but now people can update to OpenWrt in a couple of minutes, ensuring they will be able to maintain the security of their device long after Zyxel stops supporting it, not to mention all the new (OpenWrt) features they can now enable if they so choose.
There are many other recent examples of such successful reverse engineering efforts. I covered several of these in the keynote at SeaGL a couple weeks ago (including a guitar amp , the Meta Quest bootloader , a variety of music gear , and the Nintendo Switch 2 - see Slide 16 of the slides for details). Since then there have been even more, including this article on the sub-$100 FM-1, which was covered more thoroughly by an electronic music news site just last week.
Some journalists, recognizing how many of these projects are developed, have been raising the alarm about where this will go if we're not careful. Jason Koebler of 404 Media notes that Inevitably, poorly coded, barely-functional, insecure software is going to be installed on critical hardware, which will lead to devices and networks being hacked . He's right, and this is why we need FOSS communities more than ever.
In the example of the Zyxel XS1930, freed by OpenWrt, LLM-gen-AI was used only to the extent necessary to understand the binary software image format the stock firmware expects. Beyond that, the device runs all the upstream projects distributed by OpenWrt, which have been carefully reviewed and developed by humans over the past several decades. This contrasts with many of the other reverse engineering projects mentioned above, where the software that replaced the proprietary code that previously ran on the device was itself generated by LLM-gen-AI with little or no human review. Especially with the LLM-gen-AI available today, this can be very dangerous, as Koebler correctly noted.
We at SFC, with our long history of finding ways for users to take control of their devices (using a variety of mechanisms, like lawsuits when necessary , our ongoing work toward software freedom on 3-D Printers also building our own community-focused hardware ), have reached out to these reverse engineering projects to see if we at SFC can help, and to make connections with our member projects, where it makes sense and does not increase the risk of project burnout.
A related endeavour that I've personally been pushing since my first year at SFC over 15 years ago is to encourage companies to intentionally make their devices open to software modifications. While there have been hints about big companies doing this recently , which people have good reason to be skeptical about, there are smaller examples like several makers of electronic music gear mentioned earlier , that explicitly encourage people to run their own firmwares on the equipment they sell, when they recognized this would be a huge boost to sales and reputation with the community. If you are part of a company or know a company who wants to promote software modifications on the devices the company sells, please reach out .
Beyond the efforts described above, SFC will continue to look for new ways to engage with people who are working to bring the best of FOSS, including to the new world of increasingly modifiable devices that is now, whether Big Tech likes it or not, truly upon us.
Please email any comments on this entry to
info@sfconservancy.org .
Other Conservancy Blog entries…
Bradley M. Kühn and Denver Gingerich
Bradley M. Kühn and Karen M. Sandler
Software Freedom Conservancy is a 501(c)(3) non-profit charity.
Privacy Policy last updated 22 December 2020.
This page and its contents are licensed under a Creative Commons Attribution-Share Alike 4.0 International License .

## Original Extract

The Software Freedom Conservancy provides a non-profit home and services to Free, Libre and Open Source Software (FLOSS) projects.

Get the latest update on our Vizio court case -->
Toggle navigation menu
News
Press Releases
We defend & uphold the rights enshrined in copyleft licenses.
We defend the legal rights of software users in Court when needed.
We help FOSS contributors navigate this ongoing conundrum.
Our annual community-oriented conference.
Leave GitHub today ! Learn why & how.
Internships for those who face under­representation, systemic bias, or discrimination in tech.
We provide non-profit infrastructure and services to our members creating Free/Libre and Open Source Software.
Our tool for evaluating the source code candidates companies must provide for GPLed software.
We designed and built the first ever wireless Internet router designed with software freedom and right to repair in mind.
Learn
The Corresponding Source
A bi-weekly oggcast about legal, policy, and many other issues in the Free, Libre, and Open Source Software (FLOSS) world.
A list of terms you might be unfamiliar with but occur frequently in our work.
Your most frequently asked questions about the Vizio lawsuit, answered in one place.
We defend & uphold the rights enshrined in copyleft licenses.
We defend the legal rights of software users in Court when needed.
We help FOSS contributors navigate this ongoing conundrum.
Our annual community-oriented conference.
Leave GitHub today ! Learn why & how.
Internships for those who face under­representation, systemic bias, or discrimination in tech.
We provide non-profit infrastructure and services to our members creating Free/Libre and Open Source Software.
Our tool for evaluating the source code candidates companies must provide for GPLed software.
We designed and built the first ever wireless Internet router designed with software freedom and right to repair in mind.
Learn
The Corresponding Source
A bi-weekly oggcast about legal, policy, and many other issues in the Free, Libre, and Open Source Software (FLOSS) world.
A list of terms you might be unfamiliar with but occur frequently in our work.
Your most frequently asked questions about the Vizio lawsuit, answered in one place.
My Keynote at SeaGL and LLM-gen-AI-assisted Reverse Engineering
by Denver Gingerich
on November 5, 2026
Over the past few months, there have been an increasing number of reports of people successfully gaining access to their own devices that were designed to prevent them from installing their own updates and modifications. Nearly every one of these reports notes that an LLM-gen-AI was used. While there are obviously a huge number of issues with LLM-gen-AIs (as we note in our LLM-gen-AI recommendations ), it is hard to ignore the benefits to people's digital autonomy that reverse engineering like this brings.
As an example, one device that was previously difficult to install updates on was recently freed by OpenWrt ( part of SFC ). Specifically, the Zyxel XS1930 (a 10- or 12-port network switch, with dual-core 1.0 GHz MIPS CPU), which ships with a fully proprietary OS called ZyNOS, has been upgraded to OpenWrt using the ZyNOS web UI, which only recently became feasible, thanks to the efforts of developers using LLM-gen-AI. Previous efforts to install OpenWrt or other Linux-based systems on the XS1930 required hardware modifications, but now people can update to OpenWrt in a couple of minutes, ensuring they will be able to maintain the security of their device long after Zyxel stops supporting it, not to mention all the new (OpenWrt) features they can now enable if they so choose.
There are many other recent examples of such successful reverse engineering efforts. I covered several of these in the keynote at SeaGL a couple weeks ago (including a guitar amp , the Meta Quest bootloader , a variety of music gear , and the Nintendo Switch 2 - see Slide 16 of the slides for details). Since then there have been even more, including this article on the sub-$100 FM-1, which was covered more thoroughly by an electronic music news site just last week.
Some journalists, recognizing how many of these projects are developed, have been raising the alarm about where this will go if we're not careful. Jason Koebler of 404 Media notes that Inevitably, poorly coded, barely-functional, insecure software is going to be installed on critical hardware, which will lead to devices and networks being hacked . He's right, and this is why we need FOSS communities more than ever.
In the example of the Zyxel XS1930, freed by OpenWrt, LLM-gen-AI was used only to the extent necessary to understand the binary software image format the stock firmware expects. Beyond that, the device runs all the upstream projects distributed by OpenWrt, which have been carefully reviewed and developed by humans over the past several decades. This contrasts with many of the other reverse engineering projects mentioned above, where the software that replaced the proprietary code that previously ran on the device was itself generated by LLM-gen-AI with little or no human review. Especially with the LLM-gen-AI available today, this can be very dangerous, as Koebler correctly noted.
We at SFC, with our long history of finding ways for users to take control of their devices (using a variety of mechanisms, like lawsuits when necessary , our ongoing work toward software freedom on 3-D Printers also building our own community-focused hardware ), have reached out to these reverse engineering projects to see if we at SFC can help, and to make connections with our member projects, where it makes sense and does not increase the risk of project burnout.
A related endeavour that I've personally been pushing since my first year at SFC over 15 years ago is to encourage companies to intentionally make their devices open to software modifications. While there have been hints about big companies doing this recently , which people have good reason to be skeptical about, there are smaller examples like several makers of electronic music gear mentioned earlier , that explicitly encourage people to run their own firmwares on the equipment they sell, when they recognized this would be a huge boost to sales and reputation with the community. If you are part of a company or know a company who wants to promote software modifications on the devices the company sells, please reach out .
Beyond the efforts described above, SFC will continue to look for new ways to engage with people who are working to bring the best of FOSS, including to the new world of increasingly modifiable devices that is now, whether Big Tech likes it or not, truly upon us.
Please email any comments on this entry to
info@sfconservancy.org .
Other Conservancy Blog entries…
Bradley M. Kühn and Denver Gingerich
Bradley M. Kühn and Karen M. Sandler
Software Freedom Conservancy is a 501(c)(3) non-profit charity.
Privacy Policy last updated 22 December 2020.
This page and its contents are licensed under a Creative Commons Attribution-Share Alike 4.0 International License .
