import argparse
import re
import subprocess
from pathlib import Path

VERSION_FILE = Path("go/monogram/version.txt")

def bump_version(current_version, bump_type):
    """Handles version bumping based on the specified type."""
    match = re.match(r"(\d+)\.(\d+)\.(\d+)(?:\.(\d+))?", current_version)
    if not match:
        print("Error: Invalid version format. Expected format: X.Y.Z or X.Y.Z.B")
        return None

    major, minor, patch, build = match.groups()
    major, minor, patch = map(int, (major, minor, patch))
    build = int(build) if build else 0  # Default to 0 if not present

    original_version = current_version  # Store current version for display

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
        print("Error: Invalid bump type. Choose 'major', 'minor', 'patch', or 'build'.")
        return None

    # Format the new version, omitting the build number if it's zero.
    new_version = (
        f"{major}.{minor}.{patch}"
        if build == 0
        else f"{major}.{minor}.{patch}.{build}"
    )

    print(f"Current version: {original_version}")
    print(f"Updated version: {new_version}")

    return new_version

def save_version(new_version):
    """Writes the new version to version.txt and commits the change (without pushing)."""
    VERSION_FILE.parent.mkdir(parents=True, exist_ok=True)
    VERSION_FILE.write_text(new_version)
    print(f"Saved new version to {VERSION_FILE}")

    try:
        subprocess.run(["git", "add", str(VERSION_FILE)], check=True)
        subprocess.run(["git", "commit", "-m", f"Bump version to {new_version}"], check=True)
        print("Version file committed (but not pushed).")
    except subprocess.CalledProcessError as e:
        print(f"Error during Git commit: {e}")

def tag_and_push():
    """Pushes committed changes, tags the version, and pushes the tag."""
    if not VERSION_FILE.exists():
        print("Error: version.txt not found! Cannot publish.")
        return

    new_version = VERSION_FILE.read_text().strip()
    tag_name = f"v{new_version}"

    try:
        subprocess.run(["git", "push", "origin", "main"], check=True)
        subprocess.run(["git", "tag", tag_name], check=True)
        subprocess.run(["git", "push", "origin", tag_name], check=True)
        print(f"Successfully tagged and pushed: {tag_name}")
    except subprocess.CalledProcessError as e:
        print(f"Error during Git operations: {e}")

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

        # Validation in __init__
        # If --publish is used with --bump, then --save must be provided.
        if self.args.publish and self.args.bump and not self.args.save:
            parser.error("--publish with --bump requires --save to record the new version.")

        # Also, if --save is used without --bump, that's not valid.
        if self.args.save and not self.args.bump:
            parser.error("--save requires --bump to produce a new version.")

    def main(self):
        new_version = None

        if self.args.bump:
            current_version = (
                VERSION_FILE.read_text().strip()
                if VERSION_FILE.exists()
                else "0.0.0"
            )
            new_version = bump_version(current_version, self.args.bump)

        if self.args.save and new_version:
            if self.args.yes or input("Do you want to save and commit the new version? (y/N): ").lower() in ("y", "yes"):
                save_version(new_version)
            else:
                print("Skipping save step.")

        if self.args.publish:
            if self.args.yes or input("Do you want to publish the release (push changes and tag)? (y/N): ").lower() in ("y", "yes"):
                tag_and_push()
            else:
                print("Skipping publish step.")

if __name__ == "__main__":
    Main().main()
