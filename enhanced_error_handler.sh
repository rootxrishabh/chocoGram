#!/bin/bash

notify_error() {
  local exit_code=$1
  local error_string=$2
  local test_uuid=$3
  local question_id=$4
  local user_id=$5

  # Building the JSON request body
  request_body="error : $error_string, user_id : $user_id, uuid : $test_uuid, ques_id : $question_id, exit_code: $exit_code"
  data=$(cat <<EOF
{
    "stacktrace": "$request_body",
    "requestMethod": "POST",
    "requestUrl": "GITPOD",
    "requestId": "$test_uuid",
    "ExceptionSource": "HTTP",
    "name": "GITPOD",
    "jobName": "$test_uuid"
}
EOF
)

  echo "Invoking notification service, notifying about error.."

  # Removing breaklines from the data
  data="${data//$'\n'/}"

  # Send the error notification using cURL
  RESPONSE=$(curl -sS --location --request POST http://notification.codejudge.io/notification/app/send-exception-email --header 'Content-Type: application/json' --data-raw "$data")
  if [ $? -eq 0 ]; then
    echo "Sent error notification"
  else
    echo "Error occurred while sending error notification"
  fi
}

# Main script
if [ $# -lt 5 ]; then
  echo "Error: Missing required arguments."
  echo "Usage: $0 <error_string> <test_uuid> <question_id>"
  exit 1
fi

exit_code=$1
error_string=$2
test_uuid=$3
question_id=$4
user_id=$5
echo "In error handler having error->: $error_string"

# Call the function to notify about the error
notify_error "$exit_code" "$error_string" "$test_uuid" "$question_id" "$user_id"
