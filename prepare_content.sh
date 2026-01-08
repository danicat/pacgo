#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Exit on error
set -e

# Get the absolute path of the script's directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJECT_ROOT="$SCRIPT_DIR"
SITE_DIR="$PROJECT_ROOT/site"

# Clear existing content
rm -rf "$SITE_DIR/content"
mkdir -p "$SITE_DIR/content/steps"

# Process the main README.md
# Fix links to point to the steps directory
sed -e 's#(./screenshot.jpg#(/screenshot.jpg#g' -e 's#step\([0-9][0-9]\)/README.md#steps/step-\1#g' "$PROJECT_ROOT/README.md" > "$SITE_DIR/content/_index.md"

# Process each step directory
for STEP_DIR in "$PROJECT_ROOT"/step[0-9][0-9]; do
  if [ -d "$STEP_DIR" ]; then
    STEP_NUM_STR=$(basename "$STEP_DIR" | sed 's/step//')
    # Convert to integer to remove leading zero
    STEP_NUM_INT=$((10#$STEP_NUM_STR))

    # Brute-force the title from the README
    TITLE=$(head -n 1 "$STEP_DIR/README.md" | sed 's/# //' | sed 's/"/\\"/g')

    # Create a new markdown file for the step
    cat > "$SITE_DIR/content/steps/step-$STEP_NUM_STR.md" << EOL
---
title: "$TITLE"
weight: $STEP_NUM_INT
---
EOL

        # Append the step's README.md, fixing inter-step links
    sed -e 's#../step\([0-9][0-9]\)/README.md#../step-\1/#g' "$STEP_DIR/README.md" >> "$SITE_DIR/content/steps/step-$STEP_NUM_STR.md"

    # Append the step's main.go
    printf "\n## Code for Step %s\n" "$STEP_NUM_STR" >> "$SITE_DIR/content/steps/step-$STEP_NUM_STR.md"
    echo '```go' >> "$SITE_DIR/content/steps/step-$STEP_NUM_STR.md"
    gofmt "$STEP_DIR/main.go" >> "$SITE_DIR/content/steps/step-$STEP_NUM_STR.md"
    echo '```' >> "$SITE_DIR/content/steps/step-$STEP_NUM_STR.md"
  fi

done

# Copy the screenshot to the static directory
mkdir -p "$SITE_DIR/static"
cp "$PROJECT_ROOT/screenshot.jpg" "$SITE_DIR/static/"

echo "Content preparation complete."