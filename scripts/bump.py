import argparse
import re
import subprocess
import sys
from pathlib import Path

VERSION_FILE = Path("go/monogram/version.txt")


class BumpError(Exception):
    """Custom exception for version errors."""
    pass


def get_current_branch():
    """Returns the name of the current Git branch."""
    try:
        result = subprocess.run(
            ["git", "rev-parse", "--abbrev-ref", "HEAD"],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            universal_newlines=True,
            check=True,
        )
        return result.stdout.strip()
    except subprocess.CalledProcessError as e:
        raise BumpError(f"Unable to determine current Git branch: {e.stderr.strip()}") from e


def bump_version(current_version, bump_type):
    """Handles version bumping based on the specified type."""
    match = re.match(r"(\d+)\.(\d+)\.(\d+)(?:\.(\d+))?", current_version)
    if not match:
        raise BumpError("Invalid version format. Expected format: X.Y.Z or X.Y.Z.B")
    
    major, minor, patch, build = match.groups()
    major, minor, patch = map(int, (major, minor, patch))
    build = int(build) if build else 0  # Default to 0 if not present

    original_version = current_version

    if bump_type == "major":
        major += 1
        minor = 0
        patch = 0
        build = 0
    elif bump_type == "minor":
        minor += 1
        patch = 0
        build = 0
    elif bump_type == "patch":
        patch += 1
        build = 0
    elif bump_type == "build":
        build += 1  # Increment the build number
    else:
        raise BumpError("Invalid bump type. Choose 'major', 'minor', 'patch', or 'build'.")

    new_version = f"{major}.{minor}.{patch}" if build == 0 else f"{major}.{minor}.{patch}.{build}"

    print(f"Current version: {original_version}")
    print(f"Updated version: {new_version}")
    return new_version


def save_version(new_version):
    """Writes the new version to version.txt and commits the change (without pushing)."""
    try:
        VERSION_FILE.parent.mkdir(parents=True, exist_ok=True)
        VERSION_FILE.write_text(new_version)
        print(f"Saved new version to {VERSION_FILE}")
    except Exception as e:
        raise BumpError(f"Error writing version file: {e}") from e

    try:
        subprocess.run(["git", "add", str(VERSION_FILE)], check=True)
        subprocess.run(["git", "commit", "-m", f"Bump version to {new_version}"], check=True)
        print("Version file committed (but not pushed).")
    except subprocess.CalledProcessError as e:
        raise BumpError(f"Error during Git commit: {e}") from e


def tag_and_push():
    """Pushes committed changes, tags the version, and pushes the tag."""
    if not VERSION_FILE.exists():
        raise BumpError("version.txt not found! Cannot publish.")

    new_version = VERSION_FILE.read_text().strip()
    tag_name = f"v{new_version}"

    try:
        subprocess.run(["git", "push", "origin", "main"], check=True)
        subprocess.run(["git", "tag", tag_name], check=True)
        subprocess.run(["git", "push", "origin", tag_name], check=True)
        print(f"Successfully tagged and pushed: {tag_name}")
    except subprocess.CalledProcessError as e:
        raise BumpError(f"Error during Git operations: {e}") from e


class Main:
    def __init__(self):
        """Handles argument parsing and validates flag usage."""
        parser = argparse.ArgumentParser(
            description="Bump, save, and publish the version."
        )
        parser.add_argument(
            "--bump",
            choices=["major", "minor", "patch", "build"],
            help="Type of version bump to apply."
        )
        parser.add_argument(
            "--save",
            action="store_true",
            help="Save the bumped version to version.txt and commit it."
        )
        parser.add_argument(
            "--publish",
            action="store_true",
            help="Push committed changes and tag the bumped version."
        )
        parser.add_argument(
            "--yes",
            action="store_true",
            help="Skip confirmation prompts for save and publish."
        )

        self.args = parser.parse_args()

        # Validation: If --publish is used together with --bump, then --save must be provided.
        if self.args.publish and self.args.bump and not self.args.save:
            parser.error("--publish with --bump requires --save to record the new version.")

        # Also, if --save is used without --bump, that's not valid.
        if self.args.save and not self.args.bump:
            parser.error("--save requires --bump to produce a new version.")

    def main(self):
        try:
            # Environmental check: publishing is only allowed on the 'main' branch.
            if self.args.publish:
                current_branch = get_current_branch()
                if current_branch != "main":
                    raise BumpError(f"Publish step can only run on the 'main' branch, but you're on '{current_branch}'.")

            new_version = None
            if self.args.bump:
                current_version = VERSION_FILE.read_text().strip() if VERSION_FILE.exists() else "0.0.0"
                new_version = bump_version(current_version, self.args.bump)

            if self.args.save and new_version:
                if self.args.yes or input("Do you want to save and commit the new version? (y/N): ").lower() in ("y", "yes"):
                    save_version(new_version)
                else:
                    sys.stdout.write("Skipping save step.\n")

            if self.args.publish:
                if self.args.yes or input("Do you want to publish the release (push changes and tag)? (y/N): ").lower() in ("y", "yes"):
                    tag_and_push()
                else:
                    sys.stdout.write("Skipping publish step.\n")

        except Exception as error:
            sys.stderr.write(f"Fatal error: {error}\n")
            sys.exit(-1)


if __name__ == "__main__":
    Main().main()
