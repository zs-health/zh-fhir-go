# Deployment Summary for Bangladesh Core FHIR Implementation Guide

This document summarizes the work completed to set up automated publishing for the Bangladesh Core FHIR Implementation Guide (IG) and its accompanying documentation to GitHub Pages.

## What Was Accomplished

1.  **Automated CI/CD Pipeline Setup**: A new GitHub Actions workflow (`.github/workflows/publish-ig.yml`) has been created. This workflow automates the process of building the FHIR IG using SUSHI and the IG Publisher, and then deploying it along with the VitePress documentation to GitHub Pages.
2.  **FHIR R5 Upgrade**: The FHIR version references within the `BD-Core-FHIR-IG` submodule (specifically in `sushi-config.yaml` and `input/bd.fhir.core.xml`) have been updated from FHIR R4 to FHIR R5.
3.  **Documentation for Non-Technical Users**: 
    *   The main `README.md` file has been updated to include clear, non-technical explanations of what FHIR and an IG are, how to view the published content, and how the automated publishing process works.
    *   A new file, `NO-CODER-GUIDE.md`, has been created. This guide provides step-by-step instructions for non-technical users on how to navigate and understand the published FHIR IG, including explanations of profiles, terminology, and how to report issues.
    *   The `BD-Core-FHIR-IG/input/pagecontent/index.xml` file has been updated with a more comprehensive and user-friendly introduction to the IG.
    *   The `docs/index.md` (VitePress documentation homepage) has been updated to include a direct link to the published FHIR IG.
4.  **Workflow Consolidation**: The previous `docs.yml` workflow has been removed, and its functionality for building and deploying VitePress documentation has been integrated into the new `publish-ig.yml` workflow, streamlining the CI/CD process.

## URLs to Access the Published Content

Once successfully deployed, the Bangladesh Core FHIR IG and the VitePress documentation will be accessible at the following URL:

*   **Combined GitHub Pages URL**: [https://zs-health.github.io/zh-fhir-go/](https://zs-health.github.io/zh-fhir-go/)

This single URL will host both the generated FHIR IG content and the VitePress-based project documentation.

## How the Automated Publishing Works

*   **Trigger**: The `publish-ig.yml` workflow is automatically triggered whenever changes are pushed to the `main` branch or a pull request is opened targeting the `main` branch.
*   **Build Process**:
    1.  The repository and its submodules are checked out.
    2.  Node.js dependencies for the VitePress documentation are installed, and the documentation is built.
    3.  Java is set up, which is required for the FHIR IG Publisher.
    4.  The FHIR IG is built using `qligier/fhir-ig-action`, which internally uses SUSHI (to compile FSH files) and the official FHIR IG Publisher.
*   **Deployment**: The built FHIR IG (from `BD-Core-FHIR-IG/output`) and the VitePress documentation (from `docs/.vitepress/dist`) are uploaded as separate artifacts. These artifacts are then deployed to GitHub Pages using the `actions/deploy-pages@v4` action.

This ensures that any updates to the FHIR profiles, FSH files, or project documentation are automatically reflected on the live GitHub Pages site after a successful push to `main`.

## Troubleshooting Tips for Common Issues

*   **Workflow Failure**: If the GitHub Actions workflow fails, check the "Actions" tab in the GitHub repository. Click on the failed workflow run to view the logs. Error messages in the logs will provide clues about what went wrong (e.g., syntax errors in FSH, build issues, deployment problems).
*   **Content Not Updating**: If you push changes but don't see them reflected on GitHub Pages, ensure the workflow ran successfully. Sometimes, caching issues might occur; try clearing your browser cache or waiting a few minutes.
*   **IG Build Errors**: Issues during the FHIR IG build are often related to errors in the FSH files or the `sushi-config.yaml`. The workflow logs will typically show detailed errors from SUSHI or the IG Publisher.
*   **Permissions Issues**: Ensure the `GITHUB_TOKEN` has the necessary `contents: write`, `pages: write`, and `id-token: write` permissions in the workflow file and repository settings.

## Next Steps for Maintaining the IG

*   **Regular Updates**: Continue to update the FSH files in the `BD-Core-FHIR-IG` submodule as the Bangladesh Core FHIR profiles evolve.
*   **Documentation Maintenance**: Keep the `README.md`, `NO-CODER-GUIDE.md`, and VitePress documentation (`docs/` directory) up-to-date with any changes to the IG or project.
*   **Review Workflow Runs**: Periodically check the GitHub Actions runs to ensure the automated publishing continues to function correctly.
*   **Community Feedback**: Actively monitor and respond to issues reported on the GitHub Issues page to continuously improve the IG.
