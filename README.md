# Cloak

A simple, secure CLI tool for encrypting and decrypting secret files like `.env` files using AES-256-GCM encryption.
Protect your secrets when pushing code to GitHub or deploying to VPS servers.

- [Cloak](#cloak)
- [Installation](#installation)
- [Usage](#usage)
  - [Encrypt a file](#encrypt-a-file)
  - [Decrypt a file](#decrypt-a-file)
- [Features](#features)
- [Example Workflow](#example-workflow)

## Installation

if using Mac or Linux, you can install via Homebrew:
```bash
brew tap bnmwag/cloak
brew install cloak
```

or you can also download the binaries manually [here](https://github.com/bnmwag/cloak/releases/tag/v0.1.0).

## Usage

### Encrypt a file

```bash
cloak encrypt --input .env --output .env.enc
```

- If no flags provided, defaults are `.env` → `.env.enc`
- You will be prompted to enter a password.

### Decrypt a file

```bash
cloak decrypt --input .env.enc --output .env
```

- If no flags provided, defaults are `.env.enc` → `.env`.
- Password required to decrypt.

## Features

- Password-based encryption (AES-256-GCM)
- Easy encrypt & decrypt commands
- Build binaries for Mac, Linux, and Windows
- Clean CLI UX with smart prompts
- Secure your `.env` files easily
- Fully tested with unit and integration tests

## Example Workflow

This is a typical usage flow when working with Cloak to manage sensitive files across development and production environments.

### 1. Encrypt your `.env` file locally

Encrypt your sensitive `.env` file before pushing to Git:

```bash
cloak encrypt --input .env --output .env.enc
```

You will be prompted to set a password.  
After encryption, a new `.env.enc` file will be created.

**Important:**  
- Do not commit your original `.env` file to Git.
- Only commit the `.env.enc` encrypted file.

### 2. Push the encrypted file to your repository

```bash
git add .env.enc
git commit -m "Add encrypted environment file"
git push
```

### 3. Deploy your project to the server

Clone the repository on your VPS or production server:

```bash
git clone https://github.com/yourusername/yourproject.git
cd yourproject
```

Ensure the `cloak` binary is available on the server.  
You can install it via Homebrew:

```bash
brew tap bnmwag/cloak
brew install cloak
```

Or manually download it if Homebrew is not available.

### 4. Decrypt the `.env.enc` file on the server

Once the project is cloned, decrypt the environment file:

```bash
cloak decrypt --input .env.enc --output .env
```

You will be prompted to enter the same password you used when encrypting.

After decryption, your `.env` file will be available and the application can be started normally.