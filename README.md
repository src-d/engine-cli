<p align="center"><a href="https://www.sourced.tech"><img src="https://i.imgur.com/UKjRLsI.png" alt="source{d}" height="120px"></a></p>

[Website](https://www.sourced.tech) • [Docs](https://docs.sourced.tech) • [Blog](https://blog.sourced.tech) • [Slack](http://bit.ly/src-d-community) • [Twitter](https://twitter.com/sourcedtech)

**source{d} Engine enables powerful language-agnostic analysis of your source code and git history.**

source{d} Engine exposes powerful Universal AST's to analyze your code and a SQL engine to analyze your git history:

- **Code Retrieval**: retrieve and store git repositories as a dataset.
- **Language Agnostic Code Analysis**: automatically identify languages, parse source code, and extract the pieces that matter in a completely language-agnostic way.
- **Git Analysis** powerful SQL based analysis on top of your git repositories.
- **Querying With Familiar APIs** analyze your code through powerful friendly APIs, such as SQL, gRPC, REST, and various client libraries.

## Contents

- [Quickstart](#quickstart)
- [Examples](#guides-examples)
- [Architecture](#architecture)
- [Babelfish UAST](#babelfish-uast)
- [Clients & Connectors](#clients-connectors)
- [Community](#community)
- [Contributing](#contributing)

## Quickstart

Follow the steps below to get started with source{d| Engine.

#### 1. Install Docker

Follow these instructions:

- [Docker for Mac](https://store.docker.com/editions/community/docker-ce-desktop-mac)
- [Docker for Windows](https://store.docker.com/editions/community/docker-ce-desktop-windows)
- [Docker for Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/#install-docker-ce-1)
```bash
sudo apt-get update
sudo apt-get install docker-ce
```
- [Docker for Arch Linux](https://wiki.archlinux.org/index.php/Docker#Installation)
```bash
sudo pacman -S docker
```

#### 2. Install source{d} Engine

Download the [latest release](releases) for MacOS (Darwin), Linux or Windows.

MacOS / Linux:

```
# Make it executable
chmod +ux srcd
# Move it into your local bin folder to be executable from anywhere
sudo mv srcd /usr/local/bin/
```

Windows:

```
...
```

_or_

Install via NPM

```
npm install -g sourced-engine
```

#### 3. Start source{d} Engine with your local repositories

Now it's time to initialize the source{d} engine and provide it some repositories to analyze:

```
# Without a path it operates on the local folder,
# it works with nested folder.
srcd init

# You can also provide a path
srcd init /home/user/replace/path/

# or a GitHub organization/username 
srcd init https://github.com/org-name

# or just a single repository
srcd init https://github.com/username/repo-name
```

#### 4. Explore the source{d} Engine

To launch the web client, run the following command and start executing queries:

```bash
srcd web
```

In your browser, now go to http://localhost:8080

If you prefer to stay with the command line, you can execute:

```bash
srcd sql
```

This will open a SQL client that allows you to execute queries against your repositories.

#### 5. Start executing queries

**Top 10 repositories by commit count in HEAD**:
```sql
SELECT repository_id,commit_count 
FROM (
    SELECT r.repository_id, COUNT(*) AS commit_count
    FROM ref_commits r
    WHERE r.ref_name = 'HEAD'
    GROUP BY r.repository_id
) AS q
ORDER BY commit_count
DESC
LIMIT 10
```

**Query all files from HEAD**:
```sql
SELECT cf.file_path, f.blob_content 
FROM ref_commits r 
NATURAL JOIN commit_files cf 
NATURAL JOIN files f 
WHERE r.ref_name = 'HEAD' 
AND r.index = 0
```

**Retrieve the UAST for all files from HEAD**:
```sql
SELECT * FROM (
    SELECT cf.file_path,
           UAST(f.blob_content, LANGUAGE(f.file_path,  f.blob_content)) as uast
    FROM ref_commits r 
    NATURAL JOIN commit_files cf 
    NATURAL JOIN files f 
    WHERE r.ref_name = 'HEAD' 
    AND r.index = 0
) t WHERE ARRAY_LENGTH(uast) > 0
```

**Query for all LICENSE & README files across history**:
```sql
SELECT repository_id, blob_content 
FROM files 
WHERE file_path = 'LICENSE' 
OR file_path = 'README.md'
```

<details><summary><b>Show me more queries:</b></summary>
<p>

**Extract all functions as UAST nodes for Java files from HEAD**:

```sql
SELECT
    files.repository_id,
    files.file_path,
    UAST(files.blob_content, LANGUAGE(files.file_path, files.blob_content), '//FunctionGroup') as functions
FROM files
NATURAL JOIN commit_files
NATURAL JOIN commits
NATURAL JOIN refs
WHERE
    refs.ref_name= 'HEAD'
    AND LANGUAGE(files.file_path,files.blob_content) = 'Java'
LIMIT 10;
```

**Find all files where 'trim' method is called**:

```sql
SELECT * FROM (
  SELECT
      files.repository_id,
      files.file_path,
      UAST(files.blob_content, LANGUAGE(files.file_path, files.blob_content), '//Identifier[@roleCall and @Name="trim"]') as functionCall
  FROM files
  NATURAL JOIN commit_files
  NATURAL JOIN commits
  NATURAL JOIN refs
  WHERE
      refs.ref_name = 'HEAD'
) t WHERE ARRAY_LENGTH(functionCall) > 0
```

</p>
</details>

#### 6. Next steps

You can now run the source{d} Engine, choose what you would like to do next:

- [**Analyze your git repositories**](#)
- [**Understand how your code has evolved**](#)
- [**Write your own static analysis rules**](#)
- [**Build a data pipeline for MLonCode**](#)

## Guides & Examples

Collection of guide & examples using the source{d} Engine:

- [SonarSource Java Static Analysis Rules using Babelfish](https://github.com/bblfsh/sonar-checks)


## Architecture

source{d} Engine functions as a CLI tool that provides easy access to components of the source{d} stack for Code As Data. It consists of a daemon managing all of the services (Babelfish, Enry, Gitbase etc.) which are packaged as docker containers.

<p align="center"><img src="https://i.imgur.com/PwRMw0K.png" height="150" /></p>

## Babelfish UAST

One of the most important components of the source{d} engine is the UAST. 

UAST stands for [Universal Abstract Syntax Tree](https://docs.sourced.tech/babelfish/uast/uast-specification), it is a normalized form of a programming language's AST, annotated with language agnostic roles and transformed with language agnostic concepts (e.g. Functions, Imports etc.). It enables advanced static analysis of code and easy feature extraction for statistics or Machine Learning on Code.

To parse a file for a UAST, it is as easy as:

```bash
srcd parse uast /path/to/file
```

or call it from SQL or one of the below clients.

> [Try the online playground!](http://dashboard.bblf.sh)

### Clients & Connectors

For connecting to the language parsing server (Babelfish) and analyzing the UAST, there are several language clients currently supported and maintained:

- [Babelfish Go Client](https://github.com/bblfsh/client-go)
- [Babelfish Python Client](https://github.com/bblfsh/client-python)
- [Babelfish Scala Client](https://github.com/bblfsh/client-scala)

The Gitbase Spark connector is under development, which aims to allow for an easy integration with Spark & PySpark:

- [Gitbase Spark Connector](https://github.com/src-d/gitbase-spark-connector)


## Community

source{d} has an amazing community of developers & contributors who are interested in Code As Data and/or Machine Learning on Code. Please join us! 👋

- [Slack](http://bit.ly/src-d-community)
- [Twitter](https://twitter.com/sourcedtech)
- [Email](mailto:hello@sourced.tech)

## Contributing

Contributions are **welcome and very much appreciated** 🙌
Please refer [to our contribution guide](CONTRIBUTING.md) for more details.