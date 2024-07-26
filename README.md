# SBOM (Software Bill-of-Materials)

## Description

A "Software Bill-of-Materials" is a record of various software components
required to produce some software artifact. "sbom" is a simple, single-
executable utility which makes it easy to generate and manage a software
bill-of-materials file which can be included in your release.

The software bill-of-materials file (*.sbom.json) can be generated at the
start of software development and maintained over time or it can be created at
the time of publication. How you want to manage your software bill-of-materials
is up to you.

The generated sbom file contains a basic schema for packages with enough
information for most common packages and build tools. The information provided
in each property is up to you.

## Example Usage

The following examples are only a subset of the functionality provided. Use `sbom <COMMAND> help` to learn more about what can be done.

### Create an sbom file in the root directory of your project.

`sbom init my-package`

### Add information to your package.

`sbom info add --name "My Package" --version "1.0.0.0" --author "Ronnie Bar-Kochba" --license "MIT" --repository "https://www.github.com/rbarkoch/sbom"`

### Output your package's information.

`sbom info ls`

**Output**
```json
{
    "package": "my-package",
    "name": "My Package",
    "author": "Ronnie Bar-Kochba",
    "license": "MIT",
    "repository": "https://www.github.com/rbarkoch/sbom"
}
```

### Add a package to your bill-of-materials.

`sbom package add "some-dependency" --name "Some Dependency" --author "Joe Smith" --license "MIT" --version "2.1.3.0" --type "nuget" --description "A library for doing a thing.""`

### Add a package from an existing sbom file. (NOT IMPLEMENTED YET)

`sbom package import "/some/path/package.sbom.json"`

### Output a package's information.

`sbom package ls "some-dependency"`

**Output**
```json
{
    "package": "some-dependency",
    "name": "Some Dependency",
    "description": "A library for doing a thing",
    "type": "nuget",
    "version": "2.1.3.0",
    "author": "Joe Smith",
    "license": "MIT",
}
```

### Output all packages.

`sbom package ls`

**Output**
```json
{
    "go": {
        "description": "An open source programming language supported by Google.",
        "type": "build tool",
        "version": "1.22.0"
    },

    "some-dependency": {
        "name": "Some Dependency",
        "description": "A library for doing a thing",
        "type": "nuget",
        "version": "2.1.3.0",
        "author": "Joe Smith",
        "license": "MIT",
        "packages": {
            "some-sub-package": {
                "description": "A library for doing another thing",
                "type": "nuget",
                "version": "1.2.4.0-dev",
                "author": "Jane Doe",
                "license": "MIT"
            }
        }
    }
}
```

### Update properties for a package.

`sbom package add some-dependency --name "Some Dependency"`

### Add a sub-package. (NOT IMPLEMENTED YET)

`sbom package add "some-dependency/some-sub-package" --name "Some Sub Package"`


## Example SBOM file.

The following example shows a resulting sbom file.

```json
{
    
    "package": "my-package",
    "description": "Some package that does something.",
    "author": "Ronnie Bar-Kochba",
    "license": "MIT",
    "repository": "https://www.github.com/rbarkoch/sbom",


    "packages": {
        "go": {
            "description": "An open source programming language supported by Google.",
            "type": "build tool",
            "version": "1.22.0"
        },

        "some-dependency": {
            "name": "Some Dependency",
            "description": "A library for doing a thing",
            "type": "nuget",
            "version": "2.1.3.0",
            "author": "Joe Smith",
            "license": "MIT",
            "packages": {
                "some-sub-package": {
                    "name": "Some Sub Package",
                    "description": "A library for doing another thing",
                    "type": "nuget",
                    "version": "1.2.4.0-dev",
                    "author": "Jane Doe",
                    "license": "MIT"
                }
            }
        }
    }
}
```