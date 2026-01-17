# Gemini Collaboration Hub for my_website

This document outlines the plan and process for developing the new Angular frontend and administrative features for the `my_website` project, in collaboration with the Gemini AI assistant.

## Project Goal

The primary objective is to create a new, modern frontend using Angular. This frontend will be simpler than the original Svelte application and will feature an administrative interface. This admin page will allow for dynamic content updates (like blog posts) without requiring a full redeployment. The new content will be stored in a PostgreSQL database. The new application will coexist with the old Svelte frontend, with the ability to redirect to it.

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
