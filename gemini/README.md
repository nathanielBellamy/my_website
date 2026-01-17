# Gemini Collaboration Hub for my_website

This document outlines the plan and process for developing the new Angular frontend and administrative features for the `my_website` project, in collaboration with the Gemini AI assistant.

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
  - We will be using the Angular Testing Library as our main unit test framework.
  - We will run CI/CD through Git Actions on GitHub. 
  - We already have `build.sh` and `build-dist.sh`, scripts which build the project for different targets: either the local machine or the linux (NixOS) VM respectively.
  - We will expand upon these build scripts so that a single, interactive script will complete a full deployment, including build, test, and release through Github Actions. 
  - We will use Cypress to write an E2E test suite that can target local, dev, and production environments. 
  - Each app, `marketing` and `admin` will exist in its own directory of the same name.


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

## Architectural Decisions

*   **Frontend Framework:** Angular
*   **Database:** PostgreSQL
*   **Backend:** Extend existing Go backend
*   **Hosting:** Linode with NixOS

## Collaboration Process

1.  **Planning:** We will use this document and the `TODO` list to track our work.
2.  **Changes:** All code changes will be proposed by me (Gemini). I will not modify existing files without discussing the changes first.
3.  **New Files:** New files for the collaboration process will be organized under the `gemini/` directory.
4.  **Clarification:** I will ask for clarification whenever a requirement is ambiguous. Please provide as much detail as possible in your responses.
5.  **Review:** You will have the opportunity to review all proposed changes before they are applied.
