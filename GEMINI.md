# Gemini Collaboration Hub for my_website

This document outlines the plan and process for developing the new Angular frontend and administrative features for the `my_website` project, in collaboration with the Gemini AI assistant.

## Project Goal

- The primary objective is to create two new, simple, modern frontend apps using Angular 21: `marketing` and `admin`.
  - The new app will coexist with the old Svelte `frontend`, with the ability to redirect to it.
  - Both frontends will talk to a single PostgresSQL database through the Go backend running on our NixOS VM.
  - We will run CI/CD through Git Actions on GitHub. 
  - We already have `build.sh` and `build-dist.sh`, scripts which build the project for different targets: either the local machine or the linux (NixOS) VM respectively.
  - We will expand upon these build scripts so that a single, interactive script will complete a full deployment, including build, test, and release through Github Actions. 
  - Each app, `marketing` and `admin` will exist in its own directory of the same name.
  - We will never explicitly set `undefined` as a value.

  ## Angular Conventions
  - We will be using signal stores.
  - We will only use the `@`-style notation for conditional HTML - that is, no `*ngIf`. 
  - We will be using `input<T>.required()` as opposed to `@input()`
  - We will keep all HTML in separate `foo.component.html` files
    - that is, I should never see a `template:` param in an Angular component
    - instead, it should always be `templateUrl: ./foo.component.html`
  - Whenever we `Inject` a service into an Angular component, it will be `private readonly`
  - Although the new `marketing` and `admin` apps will be written in Angular, the goal is to keep these apps doing nothing more than standard, boring CRUD - we will thus keep all asynchronouse code as close to `Promise`s as possible.
    - That is, whenever possible we will prefer to `await` a `Promise` as opposed to `subscribe`ing to an `Observable`.
  - We will be using the Angular Testing Library as our main unit test framework
    - I should never see `TestBed` in this code
  - We will use Cypress to write an E2E test suite that can target local, dev, and production environments. 
  - Never pass services into constructors. Always inject services directly: `private readonly fooService: FooService = inject(FooService);`
  - We will not use `Observable.prototype.toPromise()` as it is deprecated. Rather we will wrap these Observables in the rxjs method `firstValueFrom()` in order to convert one-shot Observables into Promises
  

  ## JavaScript/TypeScript conventions
  - We will always prefer `Enum`s over loose strings.
  - All object and interface field names should use camel case: `{ fooBar: string, baz: number }`
  - On the frontend, all ids will be UUID strings, hence all interfaces with an `id` field should define `id: string`

  ## HTML Convetions
  - Any element that the user might interact with directly (e.g. hover over, click, highlight, ...) will be given a unique data-testid so as to be easily selectible programatically.
  - All HTML will follow accessibility compliance, include full aria properties, and be easily readable. 
  - We will prefer tall and skinny HTML as opposed to long-lined HTML
  - Our opening and closing HTML tags will always be aligned vertically
  ```html
  <div>
    foo
  </div>
  ```
  - When an HTML tag has many attributes, we will list them aligned vertically and indendented
  ```html
  <div 
    class="foo bar"
    data-testid="my-div"
    [baz]="bound"
    (func)="func" >
  ```

  ## CSS Convetions
  - We will be using Tailwind for our CSS needs.
  - For common, shared colors and spacing we will use Sass variables from within Tailwind.
  - FOR THE LOVE OF GOD - to import Tailwind in `styles.css` it is just `import 'tailwindcss';`

  ## Git conventions
  - All commits will be done manually by humans. 
  - Gemini will never commit anything.
  - Gemini will never call `git add` - all adding will be done manually by humans.
  - Gemini has readonly access to git.

  ## backend/go
  - We will log an ungodly amount.
  - At every opportunity, log.
  - Catch and log all errors with full stack straces.
  - We will going to log so goddamn much, as it helps you debug.

---

- `marketing/`
  - the `marketing` frontend app is the new, Angular, public-facing bundle for `my_website`
  - it will link to the old Svelte Application
  - by default, it will load a home landing page, showing some basic informationa bout me and all the necessary links to navigate the website and all of my social links
  - there will be a `blog` page so that I have somewhere to post things that is not LinkedIn.
  - `blog` posts will be written in standard Markdown and the frontend will format this markdown nicely into html
  - There will be other pages as well that we will add with time and necessity
  - at all times, it should remain possible to navigate to any page with a single click
  - This app will have a simple, professional style: minimal and clean
  - This is the app people will see when they visit nateschieber.dev in their browser.
  - This app will have *at most* readonly privileges to the db.

- The second frontend app: `admin`
  - This is the app that only I - the admin - will have access to.
  - We already have an `auth` app, which provides lo-fi (poor but only as good as has been needed, which is to say, prevents randos stumbling upon it from immediately seeing everything) security to our dev 
  - We will expand the `auth` app to integrate with Auth0's OAuth implementation, issueing and then validating a JWT.
  - The newly expanded `auth` app wil sit in front of and guard the `admin` app
  - The `admin` app will have full admin read-write access to the db. 
  - The Go `backend` will provide basic CRUD endpoints so that the `admin` app can quickly and easily update any and all data displayed by the `marketing` app

- The `backend` Go app will also need to be updated as a number of dependencies are out of date

## High-Level Plan

A detailed, step-by-step plan will be maintained using the `TODO` list feature. This will allow us to track progress on specific tasks. The major phases are:

1.  **Setup & Planning:** Finalize requirements, set up project structure.
2.  **Database Setup:** Define schema and set up the PostgreSQL database.
3.  **Backend API Extension:** Extend the Go backend to support CRUD operations for the new content.
4.  **Admin Frontend (Angular):**
    *   Set up the new Angular project.
    *   Implement the admin login page.
    *   Build the content management interface (e.g., a rich text editor for blog posts).
5.  **Public Frontend (Angular):**
    *   Build the public-facing pages to display the content from the database.
6.  **Integration & Deployment:**
    *   Integrate the new frontend with the existing application.
    *   Update NixOS configuration for deployment on Linode.
    *   Set up CI/CD pipeline (e.g., using ArgoCD).
    *   [x] Fix ambiguous column issue in `GetAllBlogPosts` (backend)
    *   [x] Fix persistence of `activatedAt` and `deactivatedAt` timestamps for Blog Posts (Frontend default + Backend DTO mapping)
    *   [x] Fix Admin app loading in CI by aligning file serving logic with Marketing app (manual index.html fallback)

## Architectural Decisions

*   **Frontend Framework:** Angular
*   **Database:** PostgreSQL
*   **Backend:** Extend existing Go backend
*   **Hosting:** Linode with NixOS

## Collaboration Process

1.  **Planning:** We will use this document and the `TODO` list to track our work. However, `TODO_NS` are meant for me, the user, and me alone. You should never act on any information you find in a `TODO_NS`
2.  **Changes:** All code changes will be proposed by me (Gemini). I will not modify existing files without discussing the changes first.
3.  **New Files:** New files for the collaboration process will be organized under the `gemini/` directory.
4.  **Clarification:** I will ask for clarification whenever a requirement is ambiguous. Please provide as much detail as possible in your responses.
5.  **Review:** You will have the opportunity to review all proposed changes before they are applied.

## Testing



1. We will NOT test the old-site/ and auth/dev SPAs as they are being sunset and we do not care about them. They work well and have worked without fail as is for 2 years. Who needs test? Not us on these. Let's not waiste effor there. 

2. That said, we very much WILL unit test on all other projects, extensively.

3. Whenever a code path is modified and/or a new code path is added, an accompanying unit test should be written, documenting and demonstrating the new and/or modified behavior.

4. We will use Cypress for end-to-end (E2E) testing. All new display and CRUD functionality for the `admin/` and `marketing/` apps MUST have accompanying E2E tests.

5. E2E tests are located in the top-level `e2e/` directory and can be run using the `./lifecycle/e2e.sh` script.

 
