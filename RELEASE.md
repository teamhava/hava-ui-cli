# Release

## Creating a release

This release process itself should be as automated as possible, and has only a few steps:

1. **Trigger a new release by pushing a new `tag`**. 

Ideally releasing should be done often with small increments when possible. Unless a
breaking change is blocking the release, or no fixes/features have been merged, a good
target release cadence is between every 1 or 2 weeks.

Our [release process](.github/workflows/release.yml) will create the following artifacts:
- A signed Darwin Fat Binary (for Intel and ARM MAC devices)
- Linux ARM & x86_64 binary
- Windows ARM & x86_64 binary
- Docker Container pushed to hava/cli


## Retracting a release

If a release is found to be problematic, it can be retracted with the following steps:

- Deleting the GitHub Release
- Add a new `retract` entry in the go.mod for the versioned release

**Note**: do not delete release tags from the git repository since there may already be references to the release
in the go proxy, which will cause confusion when trying to reuse the tag later (the H1 hash will not match and there
will be a warning when users try to pull the new release).


## Background

A good release process has the following qualities:

1. There is a way to plan what should be in a release
1. There is a way to see what is actually in a release
1. Allow for different kinds of releases (major breaking vs backwards compatible enhancements vs patch updates)
1. Specify a repeatable way to build and publish software artifacts


### Planning a release

To indicate a set of features to be released together add each issue to an in-repository
Milestone named with major-minor version to be released (e.g. `v0.1`). It is OK for other
features to be in the release that were not originally planned, and these issues and PRs
do not need to be added to the Milestone in question. Only the set of features that, when
completed, would allow the release to be considered complete. A Milestone is only used to:

- Plan what is desired to be in a release
- Track progress to indicate when we may be ready to cut a new release

Not all releases need to be planned. For instance, patch releases for fixes should be
released when they are ready and when releasing would not interfere with another current
release (where some partial or breaking features have already been merged).

Unless necessary, feature releases should be small and frequent, which may obviate the
need for regular release planning under a Milestone.


### What is in a release

Milestones are specifically for planning a release, not necessarily tracking all changes
that a release may bring (and more importantly, not all releases are necessarily planned
either).

This is one of the (many) reasons for a Changelog. A good Changelog lists changes grouped
by the type of change (new, enhancement, deprecation, breaking, bug fix, security fix), in
chronological order (within groups), linking the PR where the change was made in the
Changelog line. Furthermore, there should be a place to see all released versions, the
release date for each release, the semantic version of the release, and the set of changes
for each release.

**This project auto-generates the Changelog contents for each current release and posts the
generated contents to the GitHub Release page**. Leveraging the GitHub Releases feature
allows GitHub to manage the Changelog on each release outside of the git source tree while
still being hosted with the released assets.

The Changelog is generated from the metadata from in-repository issues and PRs, using
labels to guide what kind of change each item is (e.g. breaking, new feature, bug fix,
etc). Only issues/PRs with select labels are included in the Changelog, and only if the
issue/PR was created after the last release. Additional labels are used to exclude items
from the Changelog.

The above suggestions imply that we should:

- Ensure there is a sufficient title for each PR and issue title to be included in the
  Changelog
- The appropriate label is applied to PRs and/or issues to drive specific change type
  sections (deprecated, breaking, security, bug, etc)

**With this approach as we cultivate good organization of PRs and issues we automatically
get an equally good Changelog.**


### Major, minor, and patch releases

The latest version of the tool is the only supported version, which implies that multiple
parallel release branches will not be a regular process (if ever). Multiple releases can
be planned in parallel, however, only one can be actively developed at a time. That is, if
PRs attached to a release Milestone have been merged into the main branch, that release is
now the "next" release. **This implies that the source of truth for release lies with the
git log and Changelog, not with the release Milestones** (which are purely for planning and
tracking).

Semantic versioning should be used to indicate breaking changes, new features, and fixes.
The exception to this is `< 1.0`, where the major version is not bumped for breaking changes,
instead the minor version indicates both new features and breaking changes.