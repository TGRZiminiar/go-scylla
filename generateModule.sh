#!/bin/bash
# chmod +x createFile.sh
# ./generateModule.sh product


if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <filename>"
  exit 1
fi

FILENAME=$1

# Create directory structure
mkdir -p "modules/${FILENAME}"
cd "modules/${FILENAME}"

# Create files
echo "package ${FILENAME}" > ${FILENAME}Entity.go
echo "package ${FILENAME}" > ${FILENAME}Model.go

mkdir "${FILENAME}Handler"
echo "package ${FILENAME}handler" > ${FILENAME}Handler/${FILENAME}HttpHandler.go

mkdir "${FILENAME}Repository"
echo "package ${FILENAME}repository" > ${FILENAME}Repository/${FILENAME}Repository.go


mkdir "${FILENAME}Usecase"
echo "package ${FILENAME}usecase" > ${FILENAME}Usecase/${FILENAME}Usecase.go

mkdir "${FILENAME}Pb"
touch "${FILENAME}Pb/${FILENAME}Pb.proto"

echo "Folder and file structure for ${FILENAME} created successfully."
