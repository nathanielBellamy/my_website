## Project Goal

- The primary objective is to create two new, simple, modern frontend apps using Angular 21: `marketing` and `admin`.
  - We will be using signal stores.
  - We will be using `input<T>.required()` as opposed to `@input()`
  - We will only use the `@`-style notation for conditional HTML - that is, no `*ngIf`. 
  - We will be using Tailwind for our CSS needs.
  - The new app will coexist with the old Svelte frontend, with the ability to redirect to it.
  - Although the new `marketing` and `admin` apps will be written in Angular, the goal is to keep these apps doing nothing more than standard, boring CRUD - we will thus keep all asynchronouse code as close to `Promise`s as possible.
  - That is, whenever possible we will prefer to `await` a `Promise` as opposed to `subscribe`ing to an `Observable`.
  - Both frontends will talk to a single PostgresSQL database through the Go backend running on our NixOS VM.
  - We will be using the Angular Testing Library together with Jest as our unit test framework.
    - I should NOT see TestBed anywhere in this code.
  - We will run CI/CD through Git Actions on GitHub. 
  - We already have `build.sh` and `build-dist.sh`, scripts which build the project for different targets: either the local machine or the linux (NixOS) VM respectively.
  - We will expand upon these build scripts so that a single, interactive script will complete a full deployment, including build, test, and release through Github Actions. 
  - We will use Cypress to write an E2E test suite that can target local, dev, and production environments. 
  - Each app, `marketing` and `admin` will exist in its own directory of the same name.


## This is the `marketing` app

- The first frontend app: `marketing`
  - the `marketing` frontend is the new, Angular, public-facing bundle for `my_website`
  - it will link to the old Svelte Application
  - by default, it will load a home landing page, showing some basic informationa bout me and all the necessary links to navigate the website and all of my social links
  - there will be a `blog` page so that I have somewhere to post things that is not LinkedIn.
  - `blog` posts will be written in standard Markdown and the frontend will format this markdown nicely into html
  - There will be other pages as well that we will add with time and necessity
  - at all times, it should remain possible to navigate to any page with a single click
  - This app will have a simple, professional style: minimal and clean
  - This is the app people will see when they visit nateschieber.dev in their browser.
  - This app will have *at most* readonly privileges to the db.

  ## Additional coding 