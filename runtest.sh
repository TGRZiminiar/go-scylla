#!/bin/bash
# chmod +x runtest.sh
for i in {1..10}
do
    # Use $i to change the number in the name and email
    name="mix${i}"
    email="mix${i}@gmail.com"

    # Run the curl command with the current parameters
    curl --location 'localhost:5000/users_v1/register' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "name": "'"${name}"'",
            "email": "'"${email}"'"
        }'

    # Add a sleep between requests if needed
    # sleep 1
done
