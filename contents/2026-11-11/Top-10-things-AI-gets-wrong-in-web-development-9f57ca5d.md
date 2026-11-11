---
source: "https://developer.chrome.com/blog/top-10-things-ai-gets-wrong-in-web-dev"
hn_url: "https://news.ycombinator.com/item?id=50455147"
title: "Top 10 things AI gets wrong in web development"
article_title: "Top 10 things AI gets wrong in web development | Blog | Chrome for Developers"
image: "https://developer.chrome.com/static/blog/top-10-things-ai-gets-wrong-in-web-dev/image/thumbnail.png"
author: "philipwalton"
captured_at: "2026-11-11T16:56:07Z"
capture_tool: "hn-digest"
hn_id: 50455147
score: 5
comments: 0
posted_at: "2026-11-11T16:19:08Z"
tags:
  - hacker-news
---

# Top 10 things AI gets wrong in web development

- HN: [50455147](https://news.ycombinator.com/item?id=50455147)
- Source: [developer.chrome.com](https://developer.chrome.com/blog/top-10-things-ai-gets-wrong-in-web-dev)
- Score: 5
- Comments: 0
- Posted: 2026-11-11T16:19:08Z

## Translation

Title: Top 10 things AI gets wrong in web development
Article title: Top 10 things AI gets wrong in web development | Blog | Chrome for Developers
Description: Even the latest frontier AI models fail around half of our modern web checks. Explore 10 common accessibility, UX, performance, and security gaps.

Article text:
Top 10 things AI gets wrong in web development | Blog | Chrome for Developers
Skip to main content
Docs
Build with Chrome
Learn how Chrome works, participate in origin trials, and build with Chrome everywhere.
Create the best experience for your users with the web's best tools.
Get things done quicker and neater, with our ready-made libraries.
Design a beautiful and performant web with Chrome.
PageSpeed Insights audit open_in_new
Baseline <span class="material-icons" style="font-size: 12px;" aria-hidden="true" translate="no">open_in_new</span>
web.dev <span class="material-icons" style="font-size: 12px;" aria-hidden="true" translate="no">open_in_new</span>
PageSpeed Insights audit <span class="material-icons" style="font-size: 12px;" aria-hidden="true" translate="no">open_in_new</span>
Top 10 things AI gets wrong in web development
Stay organized with collections
Save and categorize content based on your preferences.
Philip Walton
X
GitHub
LinkedIn
Mastodon
Bluesky
Homepage
This week, we released version 1.0 of Modern Web Guidance , a set of skills that embed web platform expertise, best practices, and browser compatibility data directly into your coding agents. Our goal with Modern Web Guidance is to help developers using AI coding agents take full advantage of the power and capabilities of the modern web platform, and ultimately improve web experiences for users.
Since we first introduced our early preview of Modern Web Guidance back in May, coding agents have gotten better and better, and many amazing new frontier models have been released. One of the questions we hear most often is: don't these new frontier models already follow web development best practices?
Unfortunately, the answer is "no"—and we know the answer is "no" because we test it.
Each guide in Modern Web Guidance defines a set of expectations that must be true if the best practices in the guide have been properly followed. Our eval suite then tests those guides by giving the same prompt to an agent running with Modern Web Guidance installed (the guided agent) and another agent running without Modern Web Guidance installed (the unguided agent). The output of both runs is checked against the expectations defined by the guide.
When you compare the guided and the unguided pass rates, you can see that AI models—even the latest frontier models—do not reliably follow modern web development best practices on their own.
On average, unguided coding agents pass only about 50% of the checks in our eval suite, compared to around 90% when Modern Web Guidance is used.
You can check out our eval dashboard to see the full list of specific checks that we run for each guide, and how each model scores on them.
But for a high-level overview, this post covers 10 of the most problematic areas that we see AI models consistently getting wrong, across key disciplines that web developers are most likely to encounter in their day-to-day work, like accessibility, user experience, performance, and security.
1. Accessibility: Keyboard focus and inert management
For people navigating with a keyboard or screen reader, predictable focus order is what makes interactive UIs usable. Without it, focus can wander into hidden background content or get stuck in dozens of extra tab stops.
Focus and inert management is one of the biggest accessibility gaps in our evals, and one of the areas that improves the most with Modern Web Guidance. The most common mistakes include:
Failing to move focus into newly opened views and forgetting to restore focus to the triggering element when navigating back.
Relying on manual JavaScript tab traps and aria-modal="true" on custom modals, instead of using a modal <dialog> or marking background content inert .
Exposing every item in a toolbar or directional grid as a separate tab stop instead of using a single tab stop with roving tabindex .
Creating custom :focus-visible outlines that lack a 3:1 contrast ratio against adjacent backgrounds.
2. Accessibility: Reduced-motion support
Unexpected page movement and large-scale animations aren't just distracting—for users with vestibular disorders or motion sensitivity, they can trigger dizziness, nausea, or headaches.
AI models usually remember to check for prefers-reduced-motion when adding standard CSS transitions or keyframe animations. However, when using newer features like smooth scrolling, view transitions, or scroll-driven animations, they often forget about reduced motion, or they make mistakes when implementing it:
Setting scroll-behavior:smooth globally and unconditionally on the root element.
Using scroll-driven CSS animations with no reduced motion check.
Removing transition durations on carousel items to enable reduced motion during scroll, but then forgetting to also remove transform properties, causing items to jump in size while scrolling.
Applying a blanket 0.01ms duration reset on * , *::before , and *::after when reduced motion is active, wiping out all essential non-motion transitions as well (needed to communicate UI state changes).
Enabling cross-document view transitions unconditionally and overriding ::view-transition-* animations when reduced motion is active, rather than just gating the @view-transition rule behind a reduced motion check.
3. Accessibility: Form validation, format hints, and autofill states
Forms are a cornerstone of most critical user journeys on the web. A minor snag in a checkout form can have a major impact on a business's bottom line.
When building forms, AI models get many of the basics right. However, they often fail on small but critical details, such as validation errors, proper placement of placeholders and format hints, and correctly integrating with password managers and autofill features:
Re-implementing CSS :user-invalid with custom JavaScript blur validation, which often flags untouched required fields as soon as users tab past them.
Letting aria-invalid fall out of sync with visible errors, either setting aria-invalid="true" prematurely on blur , or styling :user-invalid in CSS without syncing aria-invalid in JavaScript.
Relying on disappearing placeholder text to show required formats (such as MM/YY on a card expiry input) instead of rendering a persistent hint linked with aria-describedby .
Placing format hints and password rules below the <input> (instead of above it) where they can be obscured by mobile keyboards and autofill popovers right when the user needs to read them.
Highlighting :autofill inputs by changing border color alone instead of pairing color with a non-color cue that doesn't shift layout (such as an inset box-shadow ).
4. User experience: Swipe and touch gestures
Agents are good at implementing touch gestures using imperative JavaScript and writing their own custom physics logic. These implementations work; however, the JavaScript-powered touch interactions rarely feel as smooth as users have come to expect from their native app counterparts.
Many of these interactions can be implemented with CSS and native scroll APIs, letting the browser run gesture physics off the main thread. Yet in our evals, AI models often overlook these features:
Implementing swipe-to-dismiss or swipe-to-reveal gestures with pointermove or touchmove listeners instead of using CSS Scroll Snap.
Omitting overscroll-behavior-x:none on tab-swipe interfaces, which can accidentally trigger a browser back or forward navigation during an over-swipe.
Driving swipe progress effects on the main thread in JavaScript callbacks instead of using scroll-driven animations.
Not properly supporting keyboard focus on scroll-based swipe-action panels.
5. User experience: Dialogs, popovers, and top-layer transitions
In our evals, models often fall back to imperative JavaScript workarounds when animating, dismissing, or containing scroll around modal dialogs and popovers.
Instead of relying on modern HTML and CSS features built for top-layer overlays, agents frequently hand-roll entry and exit state machines, history hooks, and body scroll locks that are harder to maintain and prone to subtle UI bugs.
When you get these interactions right, it can transform a fragile-feeling experience into a polished experience that users trust. Here are the most common mistakes we see:
Building custom position:fixed overlays with manual click-outside handlers for light dismiss, instead of just using popover .
Creating entry and exit animations by toggling CSS transition classes and coordinating with requestAnimationFrame() and transitionend listeners (with setTimeout() fallbacks), instead of doing it in pure CSS with @starting-style and transition-behavior .
Locking page scroll behind an open modal by setting overflow:hidden on <html> instead of using a <dialog> element to contain overscroll.
Adding popstate listeners to detect Android back gestures and close open modals, rather than using a modal <dialog> element's native back-button close functionality.
6. User experience: Dark mode and theme-adaptive controls
Users enable dark mode to reduce eye strain and glare, so a bright white flash on page load, unthemed form controls, or a toggle that gets stuck out of sync with their OS preference immediately undermines the experience.
A dark mode toggle may seem straightforward, yet unguided models in our evals consistently miss key platform details needed to get theme switching and native form controls right:
Missing <meta name="color-scheme" content="light dark"> in the <head> , causing dark-mode users to see a white canvas flash while external stylesheets load.
Styling dark themes or dark-surfaced components with custom colors without setting color-scheme , leaving native scrollbars and form controls stuck in light mode.
Overlooking the CSS light-dark() function when defining theme-adaptive color tokens that automatically follow the active color-scheme .
Providing only a binary light or dark toggle that sets an explicit override in localStorage with no option to return to the OS default.
Forgetting to add a storage event listener, leaving other open tabs out of sync when the user changes themes.
7. Performance: Runtime responsiveness and smoothness
Users spend most of their time on a web page after it finishes loading, so sluggish interactions and janky transitions can hurt the user experience just as much as a slow initial load.
In our evals, models generally follow the classic load-time best practices that have been promoted in tools like Lighthouse for years, and they usually know to offload obvious heavy loops to a Web Worker. Where they consistently struggle is main-thread responsiveness, which is essential for optimizing INP : isolating layout work, scheduling tasks by priority, and taking advantage of new HTML and CSS features that do rendering work off the main thread.
For example, we often see agents:
Rolling custom setTimeout() queues for background or low-priority work instead of prioritizing user-visible tasks with scheduler.postTask() or taking advantage of task continuation with scheduler.yield() when the browser supports it.
Building complex layouts (like dense data grids) without applying content-visibility:auto and contain-intrinsic-size , which means updates after user interactions trigger a full-page reflow.
Hiding inactive SPA views with display:none (or tearing them down) instead of using content-visibility:hidden (paired with inert ). The former destroys cached rendering state so switching back triggers full layout and paint.
Pausing offscreen <canvas> animations and background loops with Intersection Observer callbacks rather than pairing content-visibility:auto with the contentvisibilityautostatechange event (where available), missing the browser's pre-render margin before elements scroll into view.
Building scroll affordance UI hints using scroll event listeners that force continuous layout calculations, rather than using CSS scroll-state queries when available (with Intersection Observer fallbacks).
8. Performance: Layout shifts a

[truncated]

## Original Extract

Even the latest frontier AI models fail around half of our modern web checks. Explore 10 common accessibility, UX, performance, and security gaps.

Top 10 things AI gets wrong in web development | Blog | Chrome for Developers
Skip to main content
Docs
Build with Chrome
Learn how Chrome works, participate in origin trials, and build with Chrome everywhere.
Create the best experience for your users with the web's best tools.
Get things done quicker and neater, with our ready-made libraries.
Design a beautiful and performant web with Chrome.
PageSpeed Insights audit open_in_new
Baseline <span class="material-icons" style="font-size: 12px;" aria-hidden="true" translate="no">open_in_new</span>
web.dev <span class="material-icons" style="font-size: 12px;" aria-hidden="true" translate="no">open_in_new</span>
PageSpeed Insights audit <span class="material-icons" style="font-size: 12px;" aria-hidden="true" translate="no">open_in_new</span>
Top 10 things AI gets wrong in web development
Stay organized with collections
Save and categorize content based on your preferences.
Philip Walton
X
GitHub
LinkedIn
Mastodon
Bluesky
Homepage
This week, we released version 1.0 of Modern Web Guidance , a set of skills that embed web platform expertise, best practices, and browser compatibility data directly into your coding agents. Our goal with Modern Web Guidance is to help developers using AI coding agents take full advantage of the power and capabilities of the modern web platform, and ultimately improve web experiences for users.
Since we first introduced our early preview of Modern Web Guidance back in May, coding agents have gotten better and better, and many amazing new frontier models have been released. One of the questions we hear most often is: don't these new frontier models already follow web development best practices?
Unfortunately, the answer is "no"—and we know the answer is "no" because we test it.
Each guide in Modern Web Guidance defines a set of expectations that must be true if the best practices in the guide have been properly followed. Our eval suite then tests those guides by giving the same prompt to an agent running with Modern Web Guidance installed (the guided agent) and another agent running without Modern Web Guidance installed (the unguided agent). The output of both runs is checked against the expectations defined by the guide.
When you compare the guided and the unguided pass rates, you can see that AI models—even the latest frontier models—do not reliably follow modern web development best practices on their own.
On average, unguided coding agents pass only about 50% of the checks in our eval suite, compared to around 90% when Modern Web Guidance is used.
You can check out our eval dashboard to see the full list of specific checks that we run for each guide, and how each model scores on them.
But for a high-level overview, this post covers 10 of the most problematic areas that we see AI models consistently getting wrong, across key disciplines that web developers are most likely to encounter in their day-to-day work, like accessibility, user experience, performance, and security.
1. Accessibility: Keyboard focus and inert management
For people navigating with a keyboard or screen reader, predictable focus order is what makes interactive UIs usable. Without it, focus can wander into hidden background content or get stuck in dozens of extra tab stops.
Focus and inert management is one of the biggest accessibility gaps in our evals, and one of the areas that improves the most with Modern Web Guidance. The most common mistakes include:
Failing to move focus into newly opened views and forgetting to restore focus to the triggering element when navigating back.
Relying on manual JavaScript tab traps and aria-modal="true" on custom modals, instead of using a modal <dialog> or marking background content inert .
Exposing every item in a toolbar or directional grid as a separate tab stop instead of using a single tab stop with roving tabindex .
Creating custom :focus-visible outlines that lack a 3:1 contrast ratio against adjacent backgrounds.
2. Accessibility: Reduced-motion support
Unexpected page movement and large-scale animations aren't just distracting—for users with vestibular disorders or motion sensitivity, they can trigger dizziness, nausea, or headaches.
AI models usually remember to check for prefers-reduced-motion when adding standard CSS transitions or keyframe animations. However, when using newer features like smooth scrolling, view transitions, or scroll-driven animations, they often forget about reduced motion, or they make mistakes when implementing it:
Setting scroll-behavior:smooth globally and unconditionally on the root element.
Using scroll-driven CSS animations with no reduced motion check.
Removing transition durations on carousel items to enable reduced motion during scroll, but then forgetting to also remove transform properties, causing items to jump in size while scrolling.
Applying a blanket 0.01ms duration reset on * , *::before , and *::after when reduced motion is active, wiping out all essential non-motion transitions as well (needed to communicate UI state changes).
Enabling cross-document view transitions unconditionally and overriding ::view-transition-* animations when reduced motion is active, rather than just gating the @view-transition rule behind a reduced motion check.
3. Accessibility: Form validation, format hints, and autofill states
Forms are a cornerstone of most critical user journeys on the web. A minor snag in a checkout form can have a major impact on a business's bottom line.
When building forms, AI models get many of the basics right. However, they often fail on small but critical details, such as validation errors, proper placement of placeholders and format hints, and correctly integrating with password managers and autofill features:
Re-implementing CSS :user-invalid with custom JavaScript blur validation, which often flags untouched required fields as soon as users tab past them.
Letting aria-invalid fall out of sync with visible errors, either setting aria-invalid="true" prematurely on blur , or styling :user-invalid in CSS without syncing aria-invalid in JavaScript.
Relying on disappearing placeholder text to show required formats (such as MM/YY on a card expiry input) instead of rendering a persistent hint linked with aria-describedby .
Placing format hints and password rules below the <input> (instead of above it) where they can be obscured by mobile keyboards and autofill popovers right when the user needs to read them.
Highlighting :autofill inputs by changing border color alone instead of pairing color with a non-color cue that doesn't shift layout (such as an inset box-shadow ).
4. User experience: Swipe and touch gestures
Agents are good at implementing touch gestures using imperative JavaScript and writing their own custom physics logic. These implementations work; however, the JavaScript-powered touch interactions rarely feel as smooth as users have come to expect from their native app counterparts.
Many of these interactions can be implemented with CSS and native scroll APIs, letting the browser run gesture physics off the main thread. Yet in our evals, AI models often overlook these features:
Implementing swipe-to-dismiss or swipe-to-reveal gestures with pointermove or touchmove listeners instead of using CSS Scroll Snap.
Omitting overscroll-behavior-x:none on tab-swipe interfaces, which can accidentally trigger a browser back or forward navigation during an over-swipe.
Driving swipe progress effects on the main thread in JavaScript callbacks instead of using scroll-driven animations.
Not properly supporting keyboard focus on scroll-based swipe-action panels.
5. User experience: Dialogs, popovers, and top-layer transitions
In our evals, models often fall back to imperative JavaScript workarounds when animating, dismissing, or containing scroll around modal dialogs and popovers.
Instead of relying on modern HTML and CSS features built for top-layer overlays, agents frequently hand-roll entry and exit state machines, history hooks, and body scroll locks that are harder to maintain and prone to subtle UI bugs.
When you get these interactions right, it can transform a fragile-feeling experience into a polished experience that users trust. Here are the most common mistakes we see:
Building custom position:fixed overlays with manual click-outside handlers for light dismiss, instead of just using popover .
Creating entry and exit animations by toggling CSS transition classes and coordinating with requestAnimationFrame() and transitionend listeners (with setTimeout() fallbacks), instead of doing it in pure CSS with @starting-style and transition-behavior .
Locking page scroll behind an open modal by setting overflow:hidden on <html> instead of using a <dialog> element to contain overscroll.
Adding popstate listeners to detect Android back gestures and close open modals, rather than using a modal <dialog> element's native back-button close functionality.
6. User experience: Dark mode and theme-adaptive controls
Users enable dark mode to reduce eye strain and glare, so a bright white flash on page load, unthemed form controls, or a toggle that gets stuck out of sync with their OS preference immediately undermines the experience.
A dark mode toggle may seem straightforward, yet unguided models in our evals consistently miss key platform details needed to get theme switching and native form controls right:
Missing <meta name="color-scheme" content="light dark"> in the <head> , causing dark-mode users to see a white canvas flash while external stylesheets load.
Styling dark themes or dark-surfaced components with custom colors without setting color-scheme , leaving native scrollbars and form controls stuck in light mode.
Overlooking the CSS light-dark() function when defining theme-adaptive color tokens that automatically follow the active color-scheme .
Providing only a binary light or dark toggle that sets an explicit override in localStorage with no option to return to the OS default.
Forgetting to add a storage event listener, leaving other open tabs out of sync when the user changes themes.
7. Performance: Runtime responsiveness and smoothness
Users spend most of their time on a web page after it finishes loading, so sluggish interactions and janky transitions can hurt the user experience just as much as a slow initial load.
In our evals, models generally follow the classic load-time best practices that have been promoted in tools like Lighthouse for years, and they usually know to offload obvious heavy loops to a Web Worker. Where they consistently struggle is main-thread responsiveness, which is essential for optimizing INP : isolating layout work, scheduling tasks by priority, and taking advantage of new HTML and CSS features that do rendering work off the main thread.
For example, we often see agents:
Rolling custom setTimeout() queues for background or low-priority work instead of prioritizing user-visible tasks with scheduler.postTask() or taking advantage of task continuation with scheduler.yield() when the browser supports it.
Building complex layouts (like dense data grids) without applying content-visibility:auto and contain-intrinsic-size , which means updates after user interactions trigger a full-page reflow.
Hiding inactive SPA views with display:none (or tearing them down) instead of using content-visibility:hidden (paired with inert ). The former destroys cached rendering state so switching back triggers full layout and paint.
Pausing offscreen <canvas> animations and background loops with Intersection Observer callbacks rather than pairing content-visibility:auto with the contentvisibilityautostatechange event (where available), missing the browser's pre-render margin before elements scroll into view.
Building scroll affordance UI hints using scroll event listeners that force continuous layout calculations, rather than using CSS scroll-state queries when available (with Intersection Observer fallbacks).
8. Performance: Layout shifts a

[truncated]
