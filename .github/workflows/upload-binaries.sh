#!/bin/bash

if [[ -z "$GITHUB_TOKEN" ]]; then
  echo "Set GITHUB_TOKEN env variable"
  exit 1
fi

RELEASE_ID=$(jq --raw-output '.release.id' "$GITHUB_EVENT_PATH")
if [ -z "$RELEASE_ID" ]; then
  echo "Release ID is not set, will not upload binaries"
  exit 0
fi

IS_DRAFT=$(jq --raw-output '.release.draft' "$GITHUB_EVENT_PATH")
if [ "$IS_DRAFT" = true ]; then
  echo "This is a draft release, will not upload binaries"
  exit 0
fi

AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"
FILES="release/*"

for file in $FILES; do
  echo "Uploading file ${file}"

  if [[ ! -f "$file" || ! -s "$file" ]]; then
    echo "File ${file} does not exist or is empty"
    continue
  fi

  UPLOAD_URL="https://uploads.github.com/repos/${GITHUB_REPOSITORY}/releases/${RELEASE_ID}/assets?name=${file}"
  tmp=$(mktemp)

  response=$(curl \
    -sSL \
    -XPOST \
    -H "${AUTH_HEADER}" \
    --upload-file "${file}" \
    --header "Content-Type:application/octet-stream" \
    --write-out "%{http_code}" \
    --output $tmp \
    "${UPLOAD_URL}")

  if [ "$?" -ne 0 ]; then
    echo "ERROR: 'curl' did not return success"
    cat $tmp
    rm $tmp
    exit 1
  fi

  if [ "$response" -ge 400 ]; then
    echo "ERROR: Upload was not successful. HTTP status is $response"
    cat $tmp
    rm $tmp
    exit 1
  fi

  cat $tmp | jq .
  rm $tmp

done
