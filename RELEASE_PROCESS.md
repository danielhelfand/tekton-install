# tekton-install Release Process Documentation

`tekton-install` is released using [`goreleaser`](https://goreleaser.com/). The binaries released are 
available for Mac, Linux, and Windows and available under this repository's [releases page](https://github.com/danielhelfand/tekton-install/releases).

As of now, releases will be put out each time a change is made to the project that alters the functionality of `tekton-install`. This 
generally means that bug fixes or feature additions will result in a new release, but updating documentation, testing, or CI/CD related 
aspects will not require a new release.

### How to Carry Out a Release

Prerequisites:
* Install [`goreleaser`](https://goreleaser.com/install/)
* Have a GitHub access token with permissions to publish packages to this repository

1. Start by cloning this repository and `cd`ing into the resulting directory:

```
git clone https://github.com/danielhelfand/tekton-install
```

2. Set the GitHub access token as an environment variable:

```
export GITHUB_TOKEN="YOUR_GH_TOKEN"
```

3. Push a new branch in the format `release-v<REPLACE_WITH_VERSION>`:

```
git checkout -b release-v<REPLACE_WITH_VERSION
git push origin release-v<REPLACE_WITH_VERSION
```

Switch back to `master` branch for next steps:

```
git checkout master
```

4. Create a tag for the release. 

The version number used should be updated based on the following:
* Update the middle number for bug fixes (`v0.1.1`). After two bug fix version bumps, bump the minor version (`v0.2.1` -> `v0.0.2`).
* Update the last number for features (`v0.0.2`) and reset the bug fix version to 0 if greater than 0.
* At 10 minor version updates, update the major version (`v0.0.9` -> `v1.0.0`).

Run the commands below to tag the version and push to GitHub:
```
git tag -a v0.0.1 -m "v0.0.1 release"
git push origin v0.0.1
```

5. Run the command `goreleaser` at the root of the `tekton-install` directory.

6. Update the version numbers under the [`Install tekton-install`](https://github.com/danielhelfand/tekton-install#install-tekton-install) section of the repository's main README.

7. Update the release notes to describe the features or bug fixes in the release.

8. Update the release from a prerelease to latest release.

### Updating goreleaser.yml

Updating the contents of `.goreleaser.yml` is how to alter how `goreleaser` carries out the release process. 
See the [`goreleaser` documentation](https://goreleaser.com/intro/) for more information on options for updating 
it.