#!/bin/bash

set -e

AGENT_CONFIG_ARCHIVE_FILE="config.gz"
AGENT_CONFIG_MOUNT="/media/config-image"
CLUSTER_IMAGE_SET="/etc/assisted/manifests/cluster-image-set.yaml"
AGENT_CONFIG_MOUNT="/media/config_image"

copy_archive_contents() {
    tmpdir=$(mktemp --tmpdir -d "config-image--XXXXXXXXXX")
    cp ${AGENT_CONFIG_MOUNT}/${AGENT_CONFIG_ARCHIVE_FILE} "${tmpdir}"
    gunzip -f "${tmpdir}"/${AGENT_CONFIG_ARCHIVE_FILE}
    unzipped_file=$(echo "${tmpdir}/${AGENT_CONFIG_ARCHIVE_FILE}" | cut -d'.' -f1)

    # Get the releaseImage in the archive and verify it matches the current cluster-image-set
    release_image=$(grep releaseImage ${CLUSTER_IMAGE_SET} | sed -n -e 's/^.*releaseImage: //p')
    arch_image_set=$(cpio -icv --to-stdout ${CLUSTER_IMAGE_SET} < "${unzipped_file}")
    if [[ $(< "${CLUSTER_IMAGE_SET}") != "${arch_image_set}" ]]; then
       echo "The cluster-image-set in archive does not match current release ${release_image}"
       cleanup_files
       return 1
    fi
    echo "Archive on ${devname} contains release ${release_image}"

    # Get array from string
    IFS=',' read -r -a array <<< "${CONFIG_IMAGE_FILES}"

    # Copy expected files from archive, overwriting the existing file
    for file in "${array[@]}"
    do
       cpio -icvdu "${file}" < "${unzipped_file}"
       echo "Copied file ${file}"
    done

    echo "Successfully copied contents of ${AGENT_CONFIG_ARCHIVE_FILE} on ${devname}"

    cleanup_files
    return 0
}

cleanup_files() {

    if [[ -f "${unzipped_file}" ]]; then
       rm "${unzipped_file}"
    fi
    if [[ -d "${tmpdir}" ]]; then
       rmdir "${tmpdir}"
    fi
}

# This script will be invoked by a udev rule when it detects a device with the correct label
devname="$1"
systemd-mount --no-block --automount=yes --collect "$devname" "${AGENT_CONFIG_MOUNT}"
echo "Mounted ${devname} on ${AGENT_CONFIG_MOUNT}"

while true
do
    if [[ -f ${AGENT_CONFIG_MOUNT}/${AGENT_CONFIG_ARCHIVE_FILE} ]]; then
       # Copy contents from archive
       if copy_archive_contents; then
          break
       fi
    else
       echo "Could not find ${AGENT_CONFIG_ARCHIVE_FILE} in ${AGENT_CONFIG_MOUNT}"
    fi

    echo "Retrying to copy contents from ${AGENT_CONFIG_MOUNT}/${AGENT_CONFIG_ARCHIVE_FILE}"
    sleep 5
done

