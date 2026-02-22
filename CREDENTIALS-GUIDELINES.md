# Credentials & Secrets (No‑Coder Guide)

This document explains where and how to store the data needed to connect the
repository to external tools and services. It’s written for non‑technical users
so you can manage the settings without touching the code.

## Why This Matters

Whenever the code needs to interact with another platform – for example, to push
a Docker image, publish documentation, or deploy to a cloud provider – it must
present a secret value (an API key, personal access token, certificate, etc.) to
prove it’s allowed to do so. These **secrets should never be stored in plain text**
or committed to the repository.

GitHub provides a secure way to store these values called **Secrets**. They are
encrypted and only exposed to workflows running in your repository or organization.

## Types of Secrets

| Level | Scope | Example Use | Where to Configure |
|-------|-------|-------------|--------------------|
| Repository | Specific repo | Docker registry credentials for this project | **Settings ➜ Secrets ➜ Actions** in the repo UI |
| Organization | All repos in org | Shared cloud credentials | **Organization settings ➜ Secrets** (requires admin) |
| Environment | Subgroup within repo (e.g. `production`) | Different keys for dev/production | **Settings ➜ Environments** in repo |

*Tip:* Organization and environment secrets can only be added by users with the
appropriate admin rights.

## Common Integrations

Below are examples of services you may link; most are free/open-source.

### Docker / GitHub Container Registry

1. Create a personal access token (PAT) or use your username/password.
2. In the repository, go to **Settings ➜ Secrets ➜ Actions** and click **New
   repository secret**.
3. Name it `CR_PAT` (or any name you like) and paste the token.
4. In your workflow, reference it as `${{ secrets.CR_PAT }}`.

The existing `deploy.yml` already uses `GITHUB_TOKEN` for GHCR, so additional
credentials are only needed if you push elsewhere (e.g. Docker Hub).

### Cloud Providers (GCP, AWS, Azure)

Each provider has its own methods; generally:

1. Create a service account with limited permissions (do **not** use personal
   accounts).
2. Generate a JSON key or similar credential file.
3. Add the file contents or token to GitHub secrets, e.g. `GCP_KEY`.
4. Modify your workflow to authenticate using that secret (GitHub Actions
   marketplace hosts official "setup-gcloud", "aws-actions/configure-aws-credentials",
   etc.).

### Other Tools (SonarCloud, Slack, Sentry, etc.)

Follow the vendor’s documentation to generate a token, then store it as a secret
in GitHub and refer to it in your workflow steps.

## Adding a Secret (Step‑by‑Step)

1. Open the repository on GitHub in your browser.
2. Click **Settings** ➜ **Secrets and variables** ➜ **Actions**.
3. Click **New repository secret**.
4. Enter a **Name** (e.g. `MY_API_KEY`) and **Value** (the secret string).
5. (Optional) Add an expiration date or environment restrictions.
6. Click **Add secret**.

The secret value is now encrypted and cannot be viewed again. If you lose it,
you must generate a new one.

## Managing GitHub Security Notices

GitHub will occasionally highlight security concerns such as:

* **Dependabot alerts** for vulnerable dependencies – click the alert to see the
  recommended version or patch.
* **Secret scanning** – if GitHub detects a potential secret in a pushed commit
  (e.g. a PAT), it will send an email to repository admins and invalidate the
  secret if it belongs to GitHub.
* **Code scanning** warnings, if enabled.

When you see these notices:

1. Open the alert and read the recommendation.
2. If it involves a leaked secret, revoke/rotate the secret immediately.
3. If it’s a dependency issue, update the package and run tests.

## Best Practices

* Never commit actual secrets to the repo (even in `.env`). Use GitHub secrets
  instead.
* Use least-privilege credentials – create accounts with only the permissions you
  need.
* Rotate credentials regularly and when someone leaves the team.
* Audit your secrets periodically using **Settings ➜ Secrets**.
* Document which secret names are used in workflows so non‑coders know what to
  create.

## When You’re Done

Once the required secrets are added, any push to `main` will trigger workflows
that automatically use them. For example, the build job might read `DOCKER_TOKEN`
and the deploy job might read `GCP_KEY`; you don’t need to interact with these
once they’re set.

The secret values remain hidden, so you can safely collaborate with others
without exposing credentials.
